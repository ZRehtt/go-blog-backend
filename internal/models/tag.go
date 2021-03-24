package models

import "github.com/sirupsen/logrus"

//Tag ...
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8   `json:"state"`
}

//CountTags 计数
func (d *database) CountTags(t Tag) (int64, error) {
	var count int64
	db := d.db
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&Tag{}).Where("is_deleted = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//
func (d *database) ListTags(t Tag, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	db := d.db
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_deleted = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

//CreateTag 新建一个标签
func (d *database) CreateTag(t Tag) (*Tag, error) {
	var tag Tag
	if err := d.db.Create(&t).Error; err != nil {
		return nil, err
	}
	tag.ID = t.ID
	tag.Name = t.Name
	tag.State = t.State
	tag.CreatedBy = t.CreatedBy
	tag.IsDeleted = t.IsDeleted
	return &tag, nil
}

//UpdateTag 更新单个标签
func (d *database) UpdateTag(t Tag) error {
	err := d.db.Model(&Tag{}).Where("id = ? AND is_deleted = ?", t.ID, 0).Updates(&t).Error
	if err != nil {
		logrus.WithError(err).Error("failed to update tag in database.")
		return err
	}
	return nil
}

//DeleteTag 根据ID删除单个标签
func (d *database) DeleteTag(t Tag) error {
	err := d.db.Where("id = ? AND is_deleted = ?", t.ID, 0).Delete(&t).Error
	if err != nil {
		logrus.WithError(err).Error("failed to delete tag in database.")
		return err
	}
	return nil
}
