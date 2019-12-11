package model

type Tag struct {
    Model

    Name string `json:"name"`
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

func (t *Tag) ExistByName(name string) bool {
    db.Select("id").Where("name = ?", name).First(t)
    if t.ID > 0 {
        return true
    }
    return false
}

func (t *Tag) ExistByID(id int) bool {
    db.Select("id").Where("id = ?", id).First(t)
    if t.ID > 0 {
        return true
    }

    return false
}

func (t *Tag) Create() bool{
    db.Create(t)

    return true
}


func (t *Tag) Update(data interface {}) bool {
    db.Model(t).Updates(data)

    return true
}

