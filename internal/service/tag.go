package service

import (
	"github.com/ZRehtt/go-blog-backend/internal/models"
	"github.com/ZRehtt/go-blog-backend/pkg/app"
	"github.com/sirupsen/logrus"
)

//Tag表单输入参数绑定和验证
type CountTagRequest struct {
	Name  string `form:"name" binding:"min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof 0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof 0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof 0 1"`
}

type UpdateTagRequest struct {
	//gte: 大于等于
	ID        uint32 `form:"id" binding:"required,gte=1"`
	Name      string `form:"name" binding:"min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof 0 1"`
	UpdatedBy string `form:"updated_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//
func (s *Service) CountTags(param *CountTagRequest) (int64, error) {
	var tag models.Tag
	tag.Name = param.Name
	tag.State = param.State
	count, err := models.New(s.db).CountTags(tag)
	if err != nil {
		logrus.WithError(err).Error("could not count tags in service.")
		return 0, err
	}
	return count, nil
}

//
func (s *Service) GetTagsList(param *TagListRequest, pager *app.Pager) ([]*models.Tag, error) {
	return nil, nil
}

//CreateTag ...
func (s *Service) CreateTag(param *CreateTagRequest) (*models.Tag, error) {
	var tag models.Tag
	tag.Name = param.Name
	tag.State = param.State
	tag.CreatedBy = param.CreatedBy
	mTag, err := models.New(s.db).CreateTag(tag)
	if err != nil {
		logrus.WithError(err).Error("failed to create tag in service.")
		return nil, err
	}
	return mTag, nil
}

func (s *Service) UpdateTag(param *UpdateTagRequest) error {
	var tag models.Tag
	tag.ID = param.ID
	tag.Name = param.Name
	tag.State = param.State
	tag.UpdatedBy = param.UpdatedBy
	err := models.New(s.db).UpdateTag(tag)
	if err != nil {
		logrus.WithError(err).Error("failed to update tag in service.")
		return err
	}
	return nil
}

//
func (s *Service) DeleteTag(param *DeleteTagRequest) error {
	var tag models.Tag
	tag.ID = param.ID
	err := models.New(s.db).DeleteTag(tag)
	if err != nil {
		logrus.WithError(err).Error("failed to delete tag in service.")
		return err
	}
	return nil
}
