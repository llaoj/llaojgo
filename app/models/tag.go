package models

type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    UpdateddBy string `json:"updated_by"`
    State int `json:"state"`
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