package model

import (
    "log"
    "fmt"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "laojgo/config"
)

var db *gorm.DB

func Setup() {
    var err error
    db, err = gorm.Open(config.App.Db.Default.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        config.App.Db.Default.User, 
        config.App.Db.Default.Password, 
        config.App.Db.Default.Host, 
        config.App.Db.Default.DbName))
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    db.SingularTable(true)
    db.LogMode(false)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}

type Model struct {
    ID        uint `gorm:"primary_key" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at"`
}
