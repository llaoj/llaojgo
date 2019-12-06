package config

type DatabaseSection struct {
    Type string
    User string
    Password string
    Host string
    Name string
}

var Database = &DatabaseSection{}