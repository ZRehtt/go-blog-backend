package v1

import (
	"net/http"

	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/internal/service"
	"github.com/ZRehtt/go-blog-backend/pkg/app"
	"github.com/ZRehtt/go-blog-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

//func (a Tag) GetTagByID(c *gin.Context) {}

func (a Tag) ListTags(c *gin.Context) {
	param := service.CountTagRequest{}
	code := app.HttpSuccess
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		code = app.ErrorInvalidParams
		zap.L().Error("failed to bind tag list param.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusBadRequest, code)
		return
	}

	//分页信息
	page := app.GetPage(c)
	pageSize := globals.AppSetting.PageSize
	//统计标签总数
	totalRows, err := service.CountTags(&param)
	if err != nil {
		code = app.HttpError
		zap.L().Error("failed to count tags in routers.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusInternalServerError, code)
		return
	}

	//获取标签列表
	tags, err := service.ListTagsByPage(app.GetPageOffset(page, pageSize), pageSize)
	if err != nil {
		code = app.ErrorNotFound
		zap.L().Error("failed to get tags list in routers.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusNotFound, code)
		return
	}

	app.SuccessResponse(c, code, gin.H{
		"count": totalRows,
		"tags":  tags,
	})
	return
}

func (a Tag) CreateTag(c *gin.Context) {
	param := service.CreateTagRequest{}
	code := app.HttpSuccess
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		code = app.ErrorInvalidParams
		zap.L().Error("failed to bind tag create param.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusBadRequest, code)
		return
	}

	err = service.CreateTag(&param)
	if err != nil {
		code = app.HttpError
		zap.L().Error("failed to create tag in routers.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusInternalServerError, code)
		return
	}

	app.SuccessResponse(c, code, gin.H{
		"tag": param,
	})
	return
}

func (a Tag) UpdateTag(c *gin.Context) {
	param := service.UpdateTagRequest{}
	param.ID = utils.StrTo(c.Param("id")).MustUint32()
	code := app.HttpSuccess
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		code = app.ErrorInvalidParams
		zap.L().Error("failed to bind tag update param.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusBadRequest, code)
		return
	}

	err = service.UpdateTag(&param)
	if err != nil {
		code = app.HttpError
		zap.L().Error("failed to update tag in routers.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusInternalServerError, code)
		return
	}

	app.SuccessResponse(c, code, "tag updated successful!")
	return
}

func (a Tag) DeleteTag(c *gin.Context) {
	param := service.DeleteTagRequest{}
	param.ID = utils.StrTo(c.Param("id")).MustUint32()
	code := app.HttpSuccess
	//绑定并验证参数
	err := c.ShouldBind(&param)
	if err != nil {
		code = app.ErrorInvalidParams
		zap.L().Error("failed to bind tag delete param.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusBadRequest, code)
		return
	}

	err = service.DeleteTag(&param)
	if err != nil {
		code = app.HttpError
		zap.L().Error("failed to delete tag in routers.", zap.Any("err", err))
		app.ErrorResponse(c, http.StatusInternalServerError, code)
		return
	}

	app.SuccessResponse(c, code, "tag deleted successful!")
	return
}
