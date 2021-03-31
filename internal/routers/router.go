package routers

import (
	"github.com/ZRehtt/go-blog-backend/internal/routers/api"
	v1 "github.com/ZRehtt/go-blog-backend/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

//
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	r.GET("/version", Version)
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(middleware.JWTAuth())
	{
		apiv1.POST("/tags", tag.CreateTag)
		apiv1.GET("/tags", tag.ListTags)
		apiv1.PUT("/tags/:id", tag.UpdateTag)
		//apiv1.PATCH("/tags/:id/state")
		apiv1.DELETE("/tags/:id", tag.DeleteTag)

		apiv1.POST("/articles", article.CreateArticle)
		apiv1.GET("/articles", article.ListArticles)
		apiv1.GET("/articles/:id", article.GetArticleByID)
		apiv1.PUT("/articles/:id", article.UpdateArticle)
		//apiv1.PATCH("/articles/:id/state")
		apiv1.DELETE("/articles/:id", article.DeleteArticle)
	}

	return r
}
