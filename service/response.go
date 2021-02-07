package service

import (
	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/gin-gonic/gin"
)

//ServiceResponse ...
type ServiceResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Response 统一响应信息
func Response(ctx *gin.Context, httpCode, errCode int, data interface{}) {
	ctx.JSON(httpCode, ServiceResponse{
		Code: errCode,
		Msg:  utils.GetMessage(errCode),
		Data: data,
	})
}
