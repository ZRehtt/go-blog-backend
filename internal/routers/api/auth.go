package api

import (
	"net/http"

	"github.com/ZRehtt/go-blog-backend/internal/service"
	"github.com/ZRehtt/go-blog-backend/pkg/app"
	"github.com/ZRehtt/go-blog-backend/pkg/errcode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//GetAuth ...
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	//验证输入参数
	err := c.ShouldBind(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to bind auth request")
		response.ToErrorResponse(errcode.ErrorInvalidParams)
		return
	}

	stc := service.NewService(c.Request.Context())
	//检验auth是否已存在
	err = stc.CheckAuth(&param)
	if err != nil {
		logrus.WithError(err).Error("failed to check auth")
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	//如果auth已存在，就生成对应的token
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		logrus.WithError(err).Error("failed to generate token.")
		response.ToErrorResponse(errcode.ErrorUnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(http.StatusOK, gin.H{
		"token": token,
	})
}
