package config

import (
    "log"
    "github.com/go-ini/ini"
)

var inif *ini.File

var Server = &serverSection{}
var Database = &databaseSection{}
var App = &appSection{}

func Setup() {
    var err error
    inif, err = ini.Load("app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'app.ini': %v", err)
    }

    err = inif.Section("server").MapTo(Server)
    if err != nil {
        log.Fatalf("inif.MapTo Server err: %v", err)
    }

    err = inif.Section("database").MapTo(Database)
    if err != nil {
        log.Fatalf("inif.MapTo Database err: %v", err)
    }
    
    err = inif.Section("app").MapTo(App)
    if err != nil {
        log.Fatalf("inif.MapTo Database err: %v", err)
    }
}