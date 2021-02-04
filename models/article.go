package models

//Article 文章管理
type Article struct {
	TagID       int    `json:"tagID,omitempty" gorm:"tag_id;comment:文章标签ID"`
	Tag         *Tag   `json:"tag,omitempty" gorm:"tag" gorm:"tag;comment:文章所属标签"`
	Title       string `json:"title" gorm:"title;comment:文章标题"`
	Description string `json:"description,omitempty" gorm:"description;comment:文章简述"`
	Content     string `json:"content" gorm:"content;comment:文章内容"`
	CreatedBy   string `json:"createdBy,omitempty" gorm:"created_by;comment:文章创建者"`
	UpdatedBy   string `json:"updatedBy,omitempty" gorm:"updated_by;comment:文章修改者"`
	State       int    `json:"state" gorm:"state;default:1;comment:文章状态，状态 0为禁用、1为启用"`
}
