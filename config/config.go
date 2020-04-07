package config

import (
    // "fmt"
    "log"
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

var App AppYaml

func Setup() {
    content, err := ioutil.ReadFile("app.yaml")
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    err = yaml.Unmarshal(content, &App)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
}

