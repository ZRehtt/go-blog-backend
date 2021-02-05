package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//GetPage 获取分页页码
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * viper.GetInt("app.page_size")
	}
	return result
}
