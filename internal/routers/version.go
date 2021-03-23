package routers

import (
	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Version 应用版本信息
func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": globals.AppSetting.Version,
	})
}
