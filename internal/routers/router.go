package routers

import (
	"github.com/gin-gonic/gin"
)

//
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/version", Version)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags")
		apiv1.GET("/tags")
		apiv1.PUT("/tags/:id")
		apiv1.PATCH("/tags/:id/state")
		apiv1.DELETE("/tags/:id")

		apiv1.POST("/articles")
		apiv1.GET("/articles")
		apiv1.GET("/articles/:id")
		apiv1.PUT("/articles/:id")
		apiv1.PATCH("/articles/:id/state")
		apiv1.DELETE("/articles/:id")
	}

	return r
}
