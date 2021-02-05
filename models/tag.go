package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//Tag 文章标签管理
type Tag struct {
	Model
	Name      *string `json:"name,omitempty" gorm:"name;comment:标签名称"`
	CreatedBy *string `json:"createdBy,omitempty" gorm:"created_by;comment:标签创建者"`
	UpdatedBy *string `json:"updatedBy,omitempty" gorm:"updated_by;comment:标签修改者"`
	State     *int    `json:"state,omitempty" gorm:"state;comment:标签状态，状态 0为禁用、1为启用"`
}

//var dbTableName = db.Table("blog_tag")

//GetTags 根据分页和约束获取标签列表
func GetTags(pageNumber, pageSize int, maps interface{}) ([]*Tag, error) {
	var tags []*Tag
	if pageSize > 0 && pageNumber > 0 {
		err = db.Where(maps).Find(&tags).Offset(pageSize).Limit(pageNumber).Error
	} else {
		err = db.Table("blog_tag").Where(maps).Find(&tags).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return tags, nil
}

//GetTagTotal 根据约束计算标签总数
func GetTagTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		logrus.WithError(err).Error("Error get tag total!")
		return 0, err
	}
	return count, nil
}

//ExistTagByName 检查是否有同名标签
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_at IS NULL", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//AddTag 添加标签
func AddTag(name string, state int, createdBy string) error {
	if err := db.Create(&Tag{
		Name:      &name,
		State:     &state,
		CreatedBy: &createdBy,
	}).Error; err != nil {
		return err
	}
	return nil
}

//ExistTagByID 根据ID确定标签是否存在
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_at IS NULL", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//DeleteTag 删除标签
func DeleteTag(id int) error {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return err
	}
	return nil
}

//UpdateTag 更新标签
func UpdateTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ? AND deleted_at IS NULL", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

//CleanAllTag 清除所有的标签
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Tag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
