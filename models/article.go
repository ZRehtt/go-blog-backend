package models

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//Article 文章管理
type Article struct {
	Model
	Username    string `json:"username" gorm:"type:varchar(100);not null;comment:文章所属作者名"`
	Tags        []Tag  `json:"tags" gorm:"many2many:article_tag;foreignKey:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;comment:文章所属标签"`
	Title       string `json:"title,omitempty" gorm:"type:varchar(100);comment:文章标题"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255);comment:文章简述"`
	Content     string `json:"content,omitempty" gorm:"type:longtext;comment:文章内容"`
	State       int    `json:"state,omitempty" gorm:"type:tinyint(3);default:1;comment:文章状态，状态 0为禁用、1为启用"`
}

/*
foreignKey:user_id, 在主表也就是article表，关联字段为user_id，
references:id, 被关联表和被关联字段，也就是user表的id字段
OnUpdate:CASCADE，是指定级联更新
OnDelete:CASCADE，是指定级联删除
对于一对多关系，先往被关联表中插入数据，再插入主表数据，他会物理关联到被关联表中。
*/

//ArticlePage 分页查询结构
type ArticlePage struct {
	PageNumber int
	PageSize   int
	Article
}

//CreateArticle 创建文章
func CreateArticle(article *Article) error {
	//对于many2many关联，创建或者更新时会先Upsert关联，可以使用select和Omit跳过自动保存
	err := db.Table("article").Create(article).Error
	if err != nil {
		logrus.WithError(err).Error("Con't create article in db")
		return err
	}
	return nil
}

//GetArticleByID 根据ID查询文章
func GetArticleByID(id int) (*Article, error) {
	article := Article{}
	err := db.Table("article").Preload("Tags").Where("id = ? AND deleted_at IS NULL", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Con't find article by ID in db")
		return nil, err
	}
	return &article, nil
}

//GetArticleTotal 根据约束条件获取文章总数
func GetArticleTotal(page ArticlePage) (int64, error) {
	var count int64
	query := db.Select("article.id").Table("article")
	if page.Title != "" {
		query = query.Where("title like ?", "%"+page.Title+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		logrus.WithError(err).Error("Con't get articles total in db")
		return 0, err
	}
	return count, nil
}

//GetArticlesByPage 获取指定页码和数量的文章
func GetArticlesByPage(pageNumber, pageSize int) ([]*Article, error) {
	var articles []*Article
	pageSet := (pageNumber - 1) * pageSize
	//不需要content字段
	err := db.Select(
		"id", "created_at", "updated_at", "deleted_at", "user_id",
		"title", "description", "state",
	).Table("article").Preload("User").Preload("Tags").Offset(pageSet).Limit(pageSize).
		//文章根据创建时间排序
		Order("created_at desc").Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logrus.WithError(err).Error("Con't get articles by page in db")
		return nil, err
	}
	return articles, nil
}

//UpdateArticle 更新单篇文章
func UpdateArticle(article Article) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		//更新多对多关系，替换关联，将article数据库中tags替换成传入的article中的tags
		if err = tx.Model(&article).Association("Tags").Replace(article.Tags); err != nil {
			return err
		}
		//更新文章
		if err = tx.Updates(&article).Error; err != nil {
			logrus.WithError(err).Error("Con't update article in db")
			return err
		}
		return nil
	})
	return err
}

//DeleteArticle 删除单篇文章
func DeleteArticle(id int) error {
	//可以添加Select方法来删除关联关系
	if err := db.Where("id = ? AND deleted_at IS NULL", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
