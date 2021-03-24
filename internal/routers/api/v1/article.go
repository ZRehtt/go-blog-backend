package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) GetArticleByID(c *gin.Context) {}
func (a Article) ListArticles(c *gin.Context)   {}
func (a Article) CreateArticle(c *gin.Context)  {}
func (a Article) UpdateArticle(c *gin.Context)  {}
func (a Article) DeleteArticle(c *gin.Context)  {}
