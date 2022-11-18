package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORYNAME_USED
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCategory(category *Category) int {
	err := db.Create(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的所有文章
//查询分类列表
func GetCategoryList(pageSize, pageNum int) ([]Category, int64) {
	var categorylist []Category
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categorylist).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return categorylist, total
}

//编辑分类信息
func EditCategory(id int, category *Category) int {
	var category1 Category
	var maps = make(map[string]interface{})
	maps["name"] = category.Name
	err := db.Model(&category1).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCategory(id int) int {
	var category Category
	db.Where("id=?", id).Find(&category)
	err := db.Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
