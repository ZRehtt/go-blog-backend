package app

import (
	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/pkg/errcode"
	"github.com/gin-gonic/gin"
)

//统一接口响应处理信息

//
type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalRows int64 `json:"totalRows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{ctx}
}

func (r *Response) ToResponse(code int, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(code, data)
}

func (r *Response) ToResponseList(code int, list interface{}, totalRows int64) {
	r.Ctx.JSON(code, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  globals.AppSetting.PageSize,
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
