package models

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//Tag 文章标签管理
type Tag struct {
	*Model
	Name     string    `json:"name" gorm:"type:varchar(100);not null;comment:标签名称"`
	Articles []Article `json:"articles" gorm:"many2many:article_tag;foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:标签下所属文章列表"`
	State    uint8     `json:"state" gorm:"type:tinyint(3);default:1;comment:标签状态，状态 0为禁用、1为启用"`
}

//TagPage 分页查询标签的结构
type TagPage struct {
	PageOffSet int
	PageSize   int
	Tag
}

//CreateTag 新增标签
func CreateTag(tag *Tag) error {
	err := db.Table("tag").Select("Name", "CreatedBy", "State").Create(&tag).Error
	if err != nil {
		zap.L().Error("Can't create tag in db", zap.Any("err", err))
		return err
	}
	return nil
}

//GetTagsByPage 获取指定页码和数量的标签
func GetTagsByPage(pageOffSet, pageSize int) ([]Tag, error) {
	var tags []Tag

	err := db.Select(
		"id", "created_at", "updated_at", "name", "count(article_id) as articles",
	).Table("tag").
		Preload("Article").
		Joins("left join article_tag on tag.id = article_tag.tag_id").
		Group("id").
		Offset(pageOffSet).Limit(pageSize).
		//标签根据文章数量排序，数量相同时再按名称排序
		Order("articles desc, name desc").Find(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("Can't get tags by page in db", zap.Any("err", err))
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
		zap.L().Error("Error get tag total!", zap.Any("err", err))
		return 0, err
	}
	return count, nil
}

//ExistTagByName 检查是否有同名标签
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("Can't exist tag by name in db", zap.Any("err", err))
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//GetTagByID 根据id获取标签
func GetTagByID(id uint32) (*Tag, error) {
	tag := Tag{}
	err := db.Table("tag").Where("id = ?", id).First(&tag).Error
	if err != nil {
		zap.L().Error("Can't get tag by id in db", zap.Any("err", err))
		return nil, err
	}
	return &tag, nil
}

//ExistTagByID 根据ID确定标签是否存在
func ExistTagByID(id uint32) (bool, error) {
	var tag Tag
	err := db.Table("tag").Select("id").Where("id = ? AND deleted_at IS NULL", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("Can't get tag by id in db", zap.Any("err", err))
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

//DeleteTag 根据ID删除标签
func DeleteTag(id uint32) error {
	err := db.Where("id = ? AND deleted_at IS NULL", id).Delete(&Tag{}).Error
	if err != nil {
		zap.L().Error("Can't delete tag in db", zap.Any("err", err))
		return err
	}
	return nil
}

//UpdateTag 更新标签，忽略created_at字段
func UpdateTag(id uint32, name, updatedBy string, state uint8) error {
	err := db.Table("tag").Where("id = ? AND deleted_at IS NULL", id).Omit("created_at").Updates(&Tag{
		Name:  name,
		State: state,
		Model: &Model{
			UpdatedBy: updatedBy,
		},
	}).Error
	if err != nil {
		zap.L().Error("Can't update tag in db", zap.Any("err", err))
		return err
	}
	return nil
}
