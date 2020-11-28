package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//重新设置表名
func (Article) TableName() string {
	return TablePrefix + "article"
}

//创建触发器
func (article *Article) BeforeCreate(tx *gorm.DB) (err error) {
	article.CreatedOn = time.Now().Unix()
	return
}

//根据Id判断是不是存在
func ExistsArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

//获取所有文章
func GetArticles(page int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(page).Limit(pageSize).Find(&articles)
	return
}

//获取文章的个数
func GetArticlesTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//根据文章Id获取文章
func GetArticleById(id int) (art Article) {
	db.Preload("Tag").Where("id = ?", id).First(&art)
	return
}

//添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tagId"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["createdBy"].(string),
		State:     data["state"].(int),
	})
	return true
}

//修改文章
func EditArticle(id int, data map[string]interface{}) bool {
	log.Println(data)
	db.Model(&Article{}).Select("TagID", "Title", "Desc", "Content", "ModifiedBy", "State", "ModifiedOn").Where("id = ?", id).Updates(Article{
		TagID:      data["tagId"].(int),
		Title:      data["title"].(string),
		Desc:       data["desc"].(string),
		Content:    data["content"].(string),
		ModifiedBy: data["modifiedBy"].(string),
		State:      data["state"].(int),
	})
	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}
