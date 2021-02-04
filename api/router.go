package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//NewRouter ...
func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "my-blog",
			})
		})
	}

	return router
}
