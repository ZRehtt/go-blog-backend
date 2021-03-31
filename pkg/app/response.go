package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//统一接口响应处理信息

type Response struct {
	Ctx *gin.Context
	//状态码
	Code int
	//响应信息
	Msg string
	//响应数据
	Data interface{}
}

//NewResponse ...
func NewResponse(ctx *gin.Context, httpCode, code int, msg string, data interface{}) {
	ctx.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//SuccessResponse 返回成功的响应
func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	NewResponse(ctx, http.StatusOK, code, GetCodeMsg(code), data)
}

//ErrorResponse 返回错误响应信息
func ErrorResponse(ctx *gin.Context, httpCode, code int) {
	NewResponse(ctx, httpCode, code, GetCodeMsg(code), "")
}
