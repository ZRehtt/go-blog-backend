package app

import (
	"github.com/ZRehtt/go-blog-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

//GetPage ...
func GetPage(c *gin.Context) int {
	page := utils.StrTo(c.Query("page")).MustInt()
	if page == 0 {
		return 1
	}
	return page
}

////pageSize 默认每页显示10条
//var pageSize = globals.AppSetting.PageSize

//GetPageOffset 分页处理
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
