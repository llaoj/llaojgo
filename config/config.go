package config

import (
    "fmt"

    "github.com/go-ini/ini"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"

    "laojgo/pkg/log"
)

func Setup() {
    var endpoint = "acm.aliyun.com"
    var namespaceId = "cd568b20-bebf-4dd1-a399-621317b0754a"
    var accessKey = "LTAI4FfdykbqGtVxZ5ES3FE1"
    var secretKey = "AeEk8viwKUJxtydAZB2mvFi8OPN4vZ"
    clientConfig := constant.ClientConfig{
        Endpoint:       endpoint + ":8080",
        NamespaceId:    namespaceId,
        AccessKey:      accessKey,
        SecretKey:      secretKey,
        TimeoutMs:      5 * 1000,
        ListenInterval: 30 * 1000,
    }
    configClient, err := clients.CreateConfigClient(map[string]interface{}{
        "clientConfig": clientConfig,
    })
    if err != nil {
        log.App.Error(err)
        return
    }
    var dataId = "app.ini"
    var group = "laojgo"

    // 获取配置
    content, err := configClient.GetConfig(vo.ConfigParam{
        DataId: dataId,
        Group:  group,
    })
    if err != nil {
        log.App.Error(err)
        return
    }

    load([]byte(content))
    log.App.Info(fmt.Sprintf("GetConfig: %v", content))

    // 监听配置
    configClient.ListenConfig(vo.ConfigParam{
        DataId: dataId,
        Group:  group,
        OnChange: func(namespace, group, dataId, data string) {
            load([]byte(data))
            log.App.Info(fmt.Sprintf("ListenConfig namespace: %v group: %v dataId: %v data: %v", namespace, group, dataId, data))
        },
    })
}

var Server = &serverSection{}
var Database = &databaseSection{}
var App = &appSection{}

func load(content []byte) {
    var err error
    var inif *ini.File
    inif, err = ini.Load(content)
    // inif, err = ini.Load("app.ini")
    if err != nil {
        log.App.Error(err)
    }

    if err = inif.Section("server").MapTo(Server); err != nil {
        log.App.Error(err)
    }
    if err = inif.Section("database").MapTo(Database); err != nil {
        log.App.Error(err)
    }
    if err = inif.Section("app").MapTo(App); err != nil {
        log.App.Error(err)
    }
}


