package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignKey:Cid"`
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int;not null" json:"cid"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Img         string   `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func CreateArticle(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询分类下的所有文章
func GetCateArticleList(id, pageSize, pageNum int) ([]Article, int, int64) {
	var cateArticleList []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArticleList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATEGORYNAME_NOTEXIT, 0
	}
	return cateArticleList, errmsg.SUCCESS, total
}

//todo 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id=?", id).First(&article).Error
	if err != nil {
		return article, errmsg.ERROR_ARTICLE_NOTEXIT
	}
	return article, errmsg.SUCCESS
}

//todo 查询文章列表
func GetArticleList(pageSize, pageNum int) ([]Article, int, int64) {
	var articlelist []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articlelist).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articlelist, errmsg.SUCCESS, total
}

//编辑文章
func EditArticle(id int, article *Article) int {
	var article1 Article
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["description"] = article.Description
	maps["content"] = article.Content
	maps["img"] = article.Img

	err := db.Model(&article1).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	db.Where("id=?", id).Find(&article)
	err := db.Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
