package config

import (
    "log"
    "github.com/go-ini/ini"
)

var App *ini.File

func Setup() {
    var err error
    App, err = ini.Load("app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'app.ini': %v", err)
    }
    err = App.Section("server").MapTo(Server)
    if err != nil {
        log.Fatalf("App.MapTo Server err: %v", err)
    }
    err = App.Section("database").MapTo(Database)
    if err != nil {
        log.Fatalf("App.MapTo Database err: %v", err)
    }
}