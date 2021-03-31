package service

import "github.com/ZRehtt/go-blog-backend/internal/models"

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tagID" binding:"gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	TagID     uint32 `form:"tagID" binding:"required,gte=1"`
	Title     string `form:"title" binding:"required,min=2,max=100"`
	Desc      string `form:"desc" binding:"required,min=2,max=255"`
	Content   string `form:"content" binding:"required,min=2,max=4294967295"`
	CoverURL  string `form:"coverURL" binding:"required,url"`
	CreatedBy string `form:"createdBy" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID        uint32 `form:"id" binding:"required,gte=1"`
	TagID     uint32 `form:"tagID" binding:"required,gte=1"`
	Title     string `form:"title" binding:"min=2,max=100"`
	Desc      string `form:"desc" binding:"min=2,max=255"`
	Content   string `form:"content" binding:"min=2,max=4294967295"`
	CoverURL  string `form:"coverURL" binding:"url"`
	UpdatedBy string `form:"updatedBy" binding:"required,min=2,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	ID       uint32      `json:"id"`
	Title    string      `json:"title"`
	Desc     string      `json:"desc"`
	Content  string      `json:"content"`
	CoverUrl string      `json:"coverURL"`
	State    uint8       `json:"state"`
	Tag      *models.Tag `json:"tag"`
}
