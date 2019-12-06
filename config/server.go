package config

type ServerSection struct {
    RunMode string
    HttpPort int
}

var Server = &ServerSection{}