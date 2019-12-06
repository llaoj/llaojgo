package config

import (
    "log"
    "github.com/go-ini/ini"
)

var App *ini.File

func init() {
    var err error
    App, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'app.ini': %v", err)
    }
}