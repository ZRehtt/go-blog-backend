package v1

import (
	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/internal/service"
	"github.com/ZRehtt/go-blog-backend/pkg/app"
	"github.com/ZRehtt/go-blog-backend/pkg/errcode"
	"github.com/ZRehtt/go-blog-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

//func (a Tag) GetTagByID(c *gin.Context) {}

func (a Tag) ListTags(c *gin.Context)   {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to bind tag list param.")
		response.ToErrorResponse(errcode.ErrorInvalidParams)
		return
	}

	//分页信息
	stc := service.NewService(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: globals.AppSetting.PageSize}
	//统计标签总数
	totalRows, err := stc.CountTags(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		logrus.WithError(err).Error("failed to count tags in routers.")
		return
	}

	//获取标签列表
	tags, err := stc.GetTagsList(&param, &pager)
	if err != nil {
		logrus.WithError(err).Error("failed to get tags list in routers.")
		return
	}

	response.ToResponseList(http.StatusOK, tags, totalRows)
	return
}

func (a Tag) CreateTag(c *gin.Context)  {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to bind tag create param.")
		response.ToErrorResponse(errcode.ErrorInvalidParams)
		return
	}

	stc := service.NewService(c.Request.Context())
	tag, err := stc.CreateTag(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to create tag in routers.")
		return
	}

	response.ToResponse(http.StatusCreated, gin.H{
		"tag": *tag,
	})
	return
}

func (a Tag) UpdateTag(c *gin.Context)  {
	param := service.UpdateTagRequest{}
	param.ID = utils.StrTo(c.Param("id")).MustUint32()
	response := app.NewResponse(c)
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to bind tag update param.")
		response.ToErrorResponse(errcode.ErrorInvalidParams)
		return
	}

	stc := service.NewService(c.Request.Context())
	err = stc.UpdateTag(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to update tag in routers.")
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	response.ToResponse(http.StatusOK, "tag updated successful!")
}

func (a Tag) DeleteTag(c *gin.Context)  {
	param := service.DeleteTagRequest{}
	param.ID = utils.StrTo(c.Param("id")).MustUint32()
	response := app.NewResponse(c)
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to bind tag delete param.")
		response.ToErrorResponse(errcode.ErrorInvalidParams)
		return
	}

	stc := service.NewService(c.Request.Context())
	err = stc.DeleteTag(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to delete tag in routers.")
		response.ToErrorResponse(errcode.ServerError)
		return
	}

	response.ToResponse(http.StatusOK, "tag deleted successful!")
}
