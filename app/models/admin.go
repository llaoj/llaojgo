package models

import (
	"github.com/jinzhu/gorm"
	. "laojgo/database"
)

type Admins struct {
	gorm.Model  // 这里的配置可以让ORM 自动维护 时间戳字段，很爽有木有
	Name string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
}

// Insert 新增admin用户
func (admin *Admins) Insert() (userID uint, err error) {

	result := DB.Create(&admin) //  这里的DB变量是 database 包里定义的，Create 函数是 gorm包的创建数据API
	userID = admin.ID
	if result.Error != nil {
		err = result.Error
	}
	return  // 返回新建数据的id 和 错误信息，在控制器里接收
}

// Destroy 删除admin用户
func (admin *Admins) Destroy() (err error) {

	result := DB.Delete(&admin)
	if result.Error != nil {
		err = result.Error
	}
	return
}


// Update 修改admin用户
func (admin *Admins) Update(id int64) (user Admins, err error) {
	result := DB.Model(&admin).Where("id = ?", id).Updates(&admin)
	if result.Error != nil {
		err = result.Error
	}
	return
}


// FindOne 查询admin用户
func (admin *Admins) FindAll() (admins []Admins, err error) {

	result := DB.Find(&admins) // 这里的 &admins 跟返回参数要一致

	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
