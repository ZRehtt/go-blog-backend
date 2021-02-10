package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//Tag 文章标签管理
type Tag struct {
	Model
	Name      string    `json:"name" gorm:"type:varchar(100);not null;comment:标签名称"`
	CreatedBy string    `json:"createdBy" gorm:"type:varchar(100);not null;comment:标签创建者"`
	UpdatedBy string    `json:"updatedBy" gorm:"type:varchar(100);not null;comment:标签修改者"`
	Articles  []Article `json:"articles" gorm:"many2many:article_tag;foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:标签下所属文章列表"`
	State     int       `json:"state" gorm:"type:tinyint(3);default:1;Dcomment:标签状态，状态 0为禁用、1为启用"`
}

//TagPage 分页查询标签的结构
type TagPage struct {
	PageNumber int
	PageSize   int
	Tag
}

//CreateTag 新增标签
func CreateTag(tag *Tag) error {
	err := db.Table("tag").Select("Name", "CreatedBy", "UpdatedBy", "State").Create(&tag).Error
	if err != nil {
		logrus.WithError(err).Error("Can't create tag in db")
		return err
	}
	return nil
}

//GetTagsByPage 获取指定页码和数量的标签
func GetTagsByPage(pageNumber, pageSize int) ([]Tag, error) {
	var tags []Tag

	err := db.Select(
		"id", "created_at", "updated_at", "deleted_at", "name", "count(article_id) as articles",
	).Table("tag").
		Preload("Articles").
		Joins("left join article_tag on tag.id = article_tag.tag_id").
		Group("id").
		Offset(pageNumber).Limit(pageSize).
		//标签根据文章数量排序，数量相同时再按名称排序
		Order("articles desc, name desc").Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Can't get tags by page in db")
		return nil, err
	}
	return tags, nil
}

//GetTagTotal 根据约束计算标签总数
func GetTagTotal(page TagPage) (int64, error) {
	var count int64
	query := db.Select("tag.id").Table("tag")
	if page.Name != "" {
		query = query.Where("name like ?", "%"+page.Name+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		logrus.WithError(err).Error("Error get tag total!")
		return 0, err
	}
	return count, nil
}

//ExistTagByName 检查是否有同名标签
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Can't exist tag by name in db")
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//GetTagByID 根据id获取标签
func GetTagByID(id int) (*Tag, error) {
	tag := Tag{}
	err := db.Table("tag").Where("id = ?", id).First(&tag).Error
	if err != nil {
		logrus.WithError(err).Error("Can't get tag by id in db")
		return nil, err
	}
	return &tag, nil
}

//ExistTagByID 根据ID确定标签是否存在
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Table("tag").Select("id").Where("id = ? AND deleted_at IS NULL", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Can't get tag by id in db")
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//DeleteTag 删除标签
func DeleteTag(id int) error {
	err := db.Where("id = ? AND deleted_at IS NULL", id).Delete(&Tag{}).Error
	if err != nil {
		logrus.WithError(err).Error("Can't delete tag in db")
		return err
	}
	return nil
}

//UpdateTag 更新标签，忽略created_at字段
func UpdateTag(id int, tags Tag) error {
	err := db.Table("tag").Where("id = ? AND deleted_at IS NULL", id).Omit("created_at").Updates(&tags).Error
	if err != nil {
		logrus.WithError(err).Error("Can't update tag in db")
		return err
	}
	return nil
}
