package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//Article 文章管理
type Article struct {
	Model
	User        User   `json:"user" gorm:"one2one:user;foreignKey:user_id;references:id;comment:文章所属用户"`
	UserID      uint   `json:"userID" gorm:"type:int;not null;comment:文章所属用户ID"`
	Tags        []Tag  `json:"tags" gorm:"many2many:article_tag;foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:文章所属标签"`
	Title       string `json:"title,omitempty" gorm:"type:varchar(100);comment:文章标题"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255);comment:文章简述"`
	Content     string `json:"content,omitempty" gorm:"type:longtext;comment:文章内容"`
	State       int    `json:"state,omitempty" gorm:"type:tinyint(3);default:1;comment:文章状态，状态 0为禁用、1为启用"`
}

//ArticlePage 分页查询结构
type ArticlePage struct {
	PageNumber int
	PageSize   int
	Article
}

//GetArticleByID 根据ID查询文章
func GetArticleByID(id int) (*Article, error) {
	article := &Article{}
	err := db.Table("article").Preload("User").Preload("Tags").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Con't find article by ID!")
		return nil, err
	}
	return article, nil
}

//GetArticleTotal 根据约束条件获取文章总数
func GetArticleTotal(page ArticlePage) (int64, error) {
	var count int64
	query := db.Select("article.id").Table("article")
	if page.Title != "" {
		query = query.Where("title like ?", "%"+page.Title+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//GetArticlesByPage 获取指定页码和数量的文章
func GetArticlesByPage(pageNumber, pageSize int) ([]Article, error) {
	var articles []Article
	pageSet := (pageNumber - 1) * pageSize
	//不需要content字段
	err := db.Select(
		"id", "created_at", "updated_at", "deleted_at", "user_id",
		"title", "description", "state",
	).Table("article").Preload("User").Preload("Tags").Offset(pageSet).Limit(pageSize).
		//文章根据创建时间排序
		Order("created_at desc").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

//GetArticle 根据ID获取文章
// func GetArticle(id int) (*Article, error) {
// 	var article Article
// 	err := db.Where("id = ? AND deleted_at IS NULL", id).First(&article).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}
// 	return &article, nil
// }

//UpdateArticle 更新单篇文章
func UpdateArticle(article Article) error {
	// if err := db.Model(&Article{}).Where("id = ? AND deleted_at IS NULL", id).Updates(data).Error; err != nil {
	// 	return err
	// }
	// return nil
	//在事务中执行操作
	err := db.Transaction(func(tx *gorm.DB) error {
		//更新多对多关系
		if err = tx.Model(&article).Association("Tags").Replace(article.Tags); err != nil {
			return err
		}
		//更新文章
		if err = tx.Updates(&article).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

//AddArticle 添加文章
// func AddArticle(data map[string]interface{}) error {
// 	article := Article{
// 		TagID:     data["tag_id"].(int),
// 		Title:     data["title"].(string),
// 		Content:   data["content"].(string),
// 		CreatedBy: data["created_by"].(string),
// 		State:     data["state"].(int),
// 	}
// 	if err := db.Create(&article).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
func AddArticle(article Article) (err error) {
	err = db.Table("article").Create(&article).Error
	return
}

//DeleteArticle 删除单篇文章
func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
