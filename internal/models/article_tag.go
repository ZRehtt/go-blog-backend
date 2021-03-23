package models

//ArticleTag ...
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tagID"`
	ArticleID uint32 `json:"articleID"`
}
