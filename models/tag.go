package models

import (
	"time"

	"gorm.io/gorm"
)

//定义Tag数据模型
type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"create_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//设置表名
func (Tag) TableName() string {
	return TablePrefix + "tag"
}

//创建之前添加时间戳
func (tag *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	tag.CreatedOn = time.Now().Unix()
	return
}

//获取所有标签
func GetTags(page int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(page).Limit(pageSize).Find(&tags)
	return
}

//统计标签的个数
func GetTagsTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

//判断标签存不存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//根据Id获取标签
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

//创建
func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
	return true
}

//修改标签
func EditTag(id int, name string, modifiedBy string) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(&Tag{Name: name, ModifiedBy: modifiedBy})
	return true
}

//删除标签
func DeleteTag(id int) bool {
	// db.Model(&Tag{}).Where("id = ?", id).Update("State", 0)
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
