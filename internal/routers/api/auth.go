package api

import (
	"github.com/ZRehtt/go-blog-backend/internal/service"
	"github.com/ZRehtt/go-blog-backend/pkg/app"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//GetAuth ...
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	code := app.HttpSuccess
	//验证输入参数
	err := c.ShouldBind(&param)
	if err != nil {
		code = app.ErrorInvalidParams
		zap.L().Error("failed to bind auth request", zap.Any("err", err))
		app.ErrorResponse(c, code, app.GetCodeMsg(code), "")
		return
	}

	//检验auth是否已存在
	err = service.CheckAuth(&param)
	if err != nil {
		code = app.UnauthorizedAuthNotExist
		zap.L().Error("failed to check auth", zap.Any("err", err))
		app.ErrorResponse(c, code, app.GetCodeMsg(code), "")
		return
	}

	//如果auth已存在，就生成对应的token
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		code = app.ErrorUnauthorizedTokenGenerate
		zap.L().Error("failed to generate token.", zap.Any("err", err))
		app.ErrorResponse(c, code, app.GetCodeMsg(code), "")
		return
	}

	app.SuccessResponse(c, app.GetCodeMsg(code), gin.H{
		"token": token,
	})
}
