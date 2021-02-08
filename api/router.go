package api

import (
	v1 "github.com/ZRehtt/go-blog-backend/api/v1"
	"github.com/gin-gonic/gin"
)

//NewRouter ...
func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(gin.DebugMode)

	router.POST("/register", v1.Register)

	//router.GET("/auth", v1.GetAuth)

	apiv1 := router.Group("api/v1")
	{
		//------Tags------
		apiv1.GET("/tags/:id", v1.GetTagByID)
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.CreateTag)
		apiv1.PUT("/tags/:id", v1.UpdateTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//------Articles------
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticleByID)
		apiv1.POST("/articles", v1.CreateArticle)
		apiv1.PUT("/articles/:id", v1.UpdateArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}

	return router
}
