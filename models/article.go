package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//Article 文章管理
type Article struct {
	Model
	TagID       int    `json:"tagID,omitempty" gorm:"tag_id;comment:文章标签ID"`
	Tag         Tag    `json:"tag,omitempty" gorm:"tag;comment:文章所属标签"`
	Title       string `json:"title,omitempty" gorm:"title;comment:文章标题"`
	Description string `json:"description,omitempty" gorm:"description;comment:文章简述"`
	Content     string `json:"content,omitempty" gorm:"content;type:text;comment:文章内容"`
	CreatedBy   string `json:"createdBy,omitempty" gorm:"created_by;comment:文章创建者"`
	UpdatedBy   string `json:"updatedBy,omitempty" gorm:"updated_by;comment:文章修改者"`
	State       int    `json:"state,omitempty" gorm:"state;comment:文章状态，状态 0为禁用、1为启用"`
}

//数据库表名都带有'blog_'
//var dbTB = db.Table("blog_article")

//ExistArticleByID 根据ID判断指定文章是否存在
func ExistArticleByID(id int) bool {
	var article Article
	err := db.Table("blog_article").Select("id").Where("id = ? AND deleted_at IS NULL", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Con't find article by ID!")
		return false
	}
	if article.ID > 0 {
		return true
	}
	return false
}

//GetArticleTotal 根据约束条件获取文章总数
func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Table("blog_article").Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//GetArticles 获取指定页码和数量的文章数
func GetArticles(pageNumber, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	err := db.Table("blog_article").Preload("blog_tag").Where(maps).Offset(pageNumber).Limit(pageSize).Find(articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

//GetArticle 根据ID获取文章
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Table("blog_article").Where("id = ? AND deleted_at IS NULL", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

//UpdateArticle 更新单篇文章
func UpdateArticle(id int, data interface{}) error {
	if err := db.Table("blog_article").Model(&Article{}).Where("id = ? AND deleted_at IS NULL", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

//AddArticle 添加文章
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	if err := db.Table("blog_article").Create(&article).Error; err != nil {
		return err
	}
	return nil
}

//DeleteArticle 删除单篇文章
func DeleteArticle(id int) error {
	if err := db.Table("blog_article").Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}
	return nil
}

//CleanAllArticle 删除所有的文章
func CleanAllArticle() error {
	if err := db.Table("blog_article").Unscoped().Where("deleted_at IS NOT NULL").Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
