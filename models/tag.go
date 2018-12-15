package models

type Tag struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `json:"name" gorm:"size:20"`
	CreatedBy  string `json:"created_by" gorm:"size:20"`
	ModifiedBy string `json:"modified_by" gorm:"size:20"`
	Status     int    `json:"status" gorm:"not null"`
	Model
}

// 设置表名
func (t Tag) TableName() string {
	return "blog_tag"
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
func TagIsExistByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// 创建tag
func CreateTag(name, CreatedBy string, status int) *Tag {
	tag := &Tag{
		Name:      name,
		CreatedBy: CreatedBy,
		Status:    status,
	}
	db.Create(tag)
	return tag
}

// 通过标签id获取标签
func GetTagById(id int) (tag *Tag) {
	tag = &Tag{}
	db.First(tag, id)
	return
}

// 更新tag信息
func UpdateTagById(id int, v interface{}) int64 {
	return db.Model(&Tag{}).Where("id = ?", id).Updates(v).RowsAffected
}

// 软删除
func DeleteTagById(id int) int64 {
	return db.Where("id = ?", id).Delete(&Tag{}).RowsAffected
}
