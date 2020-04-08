package model

import (
    "log"
    "fmt"
    "time"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "laojgo/app"
)

var db *gorm.DB

func Setup() {
    var err error
    db, err = gorm.Open(app.Cfg.Db.Default.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", 
        app.Cfg.Db.Default.User, 
        app.Cfg.Db.Default.Password, 
        app.Cfg.Db.Default.Host, 
        app.Cfg.Db.Default.DbName))
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
