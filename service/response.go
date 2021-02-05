package service

import (
	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/gin-gonic/gin"
)

//Gin ...
type Gin struct {
	Ctx *gin.Context
}

//ServiceResponse ...
type ServiceResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Response 统一响应信息
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.Ctx.JSON(httpCode, ServiceResponse{
		Code: errCode,
		Msg:  utils.GetMessage(errCode),
		Data: data,
	})
}
