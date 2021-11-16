package model

import (
	"blog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//Check if category exists
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//Add category
func CreateCate(data *Category) int {
	//data.PassWord = ScryptPw(data.PassWord)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//todo Query all the articles below the category

//search category list
func GetCate(pageSize int, pageNumber int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

//Edit category
func EditCate(id int, data *Category) int {
	var cate Category
	maps := make(map[string]interface{})
	maps["name"] = data.Name
	err = db.Model(&cate).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//Delete category
func DeleteCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
