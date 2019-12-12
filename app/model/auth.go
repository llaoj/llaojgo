package model

type Auth struct {
    ID int `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func (a *Auth) Check(where map[string]interface{}) bool {
    // db.Select("id").Where(where).First(a)
    // if a.ID > 0 {
    //     return true
    // }

    return true
}