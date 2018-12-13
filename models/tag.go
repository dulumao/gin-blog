package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	Status     int    `json:"status"`
}

// 设置表名
func (t Tag) TableName() string {
	return "blog_tag"
}

// 创建前
func (t *Tag) BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("CreatedDate", time.Now().Unix())
	return
}

// 更新前
func (t *Tag) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("ModifiedDate", time.Now().Unix())
	return
}

// 获取标签
func GetTags(pageNum int, pageSize int, where interface{}) (tags []Tag) {
	db.Where(where).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 获取标签总数
func GetTagTotal(where interface{}) (count int) {
	db.Model(&Tag{}).Where(where).Count(&count)
	return
}

// 通过标签名字查询标签是否存在
func TagIsExist(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func CreateTag(name, CreatedBy string, status int) {
	db.Create(&Tag{
		Name:      name,
		CreatedBy: CreatedBy,
		Status:    status,
	})
}
