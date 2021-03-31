package service

import (
	"github.com/ZRehtt/go-blog-backend/internal/models"
	"go.uber.org/zap"
)

//Tag表单输入参数绑定和验证
type CountTagRequest struct {
	Name  string `form:"name" binding:"min=3,max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

// type TagListRequest struct {
// 	Name  string `form:"name" binding:"min=3,max=100"`
// 	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
// }

type CreateTagRequest struct {
	Name      string `json:"name" form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `json:"createdBy" form:"createdBy" binding:"required,min=2,max=100"`
	State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	//gte: 大于等于
	ID        uint32 `json:"id" form:"id" binding:"required,gte=1"`
	Name      string `json:"name" form:"name" binding:"min=3,max=100"`
	State     uint8  `json:"state" form:"state,default=1" binding:"oneof=0 1"`
	UpdatedBy string `json:"updatedBy" form:"updatedBy" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

//
func CountTags(param *CountTagRequest) (int64, error) {
	count, err := models.GetTagTotal(models.TagPage{
		Tag: models.Tag{
			Name:  param.Name,
			State: param.State,
		},
	})
	if err != nil {
		zap.L().Error("could not count tags in service.", zap.Any("err", err))
		return 0, err
	}
	return count, nil
}

//
func ListTagsByPage(page, pageSize int) ([]models.Tag, error) {
	var tags []models.Tag
	tags, err := models.GetTagsByPage(page, pageSize)
	if err != nil {
		zap.L().Error("failed to list tags in service.", zap.Any("err", err))
		return nil, err
	}
	return tags, nil
}

//CreateTag ...
func CreateTag(param *CreateTagRequest) error {
	err := models.CreateTag(&models.Tag{
		Name:  param.Name,
		State: param.State,
		Model: &models.Model{
			CreatedBy: param.CreatedBy,
		},
	})
	if err != nil {
		zap.L().Error("failed to create tag in service.", zap.Any("err", err))
		return err
	}
	return nil
}

func UpdateTag(param *UpdateTagRequest) error {
	err := models.UpdateTag(param.ID, param.Name, param.UpdatedBy, param.State)
	if err != nil {
		zap.L().Error("failed to update tag in service.", zap.Any("err", err))
		return err
	}
	return nil
}

//
func DeleteTag(param *DeleteTagRequest) error {
	err := models.DeleteTag(param.ID)
	if err != nil {
		zap.L().Error("failed to delete tag in service.", zap.Any("err", err))
		return err
	}
	return nil
}
