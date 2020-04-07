package config

type AppYaml struct {
    Name string `yaml:"name"`
    Server Server `yaml:"server"`
    Db struct {
        Default Db `yaml:"default"`
    } `yaml:"db"`
}

type Db struct {
    Type string `yaml:"type"`
    Host string `yaml:"host"`
    Port int `yaml:"port"`
    User string `yaml:"user"`
    Password string `yaml:"password"`
    DbName string `yaml:"db_name"`
}

type Server struct {
    RunMode string `yaml:"run_mode"`
    HttpPort int `yaml:"http_port"`
}