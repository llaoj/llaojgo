package model

import (
    "time"
    
    "github.com/jinzhu/gorm"
)

type Tag struct {
    Model

    Name string `json:"name" form:"name" binding:"required,max=10"`
    State int `json:"state" binding:"oneof=0 1`
    CreatedBy string `json:"created_by"`
    UpdatedBy string `json:"updated_by"`
}

func (t *Tag) TableName() string {
    return "blog_tag"
}

func (t *Tag) Get(offset, limit int, where interface {}) (tags []Tag) {
    db.Where(where).Offset(offset).Limit(limit).Find(&tags)
    return
}

func (t *Tag) GetTotal(where interface {}) (count int){
    db.Model(&Tag{}).Where(where).Count(&count)
    return
}

func (t *Tag)  ExistByName(name string) bool {
    db.Select("id").Where("name = ?", name).First(t)
    if t.ID > 0 {
        return true
    }
    return false
}

func (t *Tag) Add() bool{
    db.Create(t)

    return true
}

func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedAt", time.Now().Unix())
    return nil
}

func (t *Tag) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("UpdatedAt", time.Now().Unix())
    return nil
}