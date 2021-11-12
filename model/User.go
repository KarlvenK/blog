package model

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	PassWord string `gorm:"type:varchar(500);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//Check if user exists
func CheckUser(name string) int {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

//Add user
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//search user list
func GetUsers(pageSize int, pageNumber int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

//Edit user
