package v1

import (
	"net/http"
	"strconv"

	"github.com/ZRehtt/go-blog-backend/models"
	"github.com/ZRehtt/go-blog-backend/service"
	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//GetArticleByID /api/v1/articles/{id}
func GetArticleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	var article *models.Article
	var err error
	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		article, err = models.GetArticleByID(id)
		if err != nil {
			logrus.WithError(err).Error("Failed to get article")
			return
		}
		code = utils.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("id", id).Info("getted Article!")

	service.Response(c, http.StatusOK, code, *article)
}

//GetArticles 根据分页信息查询文章列表
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	page := models.ArticlePage{}

	err := c.ShouldBind(&page)
	if err != nil {
		logrus.WithError(err).Error("Error binding apge")
		return
	}
	code := utils.SUCCESS
	articles, err := models.GetArticlesByPage(page.PageNumber, page.PageSize)
	if err != nil {
		code = utils.ERROR
		logrus.WithError(err).Error("Error getting articles by page")
		return

	}
	data["articles"] = articles

	count, err := models.GetArticleTotal(page)
	if err != nil {
		code = utils.ERROR
		logrus.WithError(err).Error("Error getting count by page")
		return
	}
	data["count"] = count

	logrus.WithField("page", page).Info("getted articles by page!")

	service.Response(c, http.StatusOK, code, data)
}

//AddArticle 新增文章
func AddArticle(c *gin.Context) {
	article := models.Article{}
	if err := c.ShouldBind(&article); err != nil {
		logrus.WithError(err).Error("Error binding article")
		return
	}

	code := utils.INVALID_PARAMS

	if article.ID == 0 {
		// //当ID为零时需要新增用户，插入userID
		// token := c.GetHeader("token")
		// //验证token
		// claims, err := utils.ParseToken(token)
		// if err != nil {
		// 	logrus.WithError(err).Error("Failed to parse token")
		// 	return
		// }
		// //验证userID
		// user, err := models.GetUserByName(claims.Username)
		// if err != nil {
		// 	logrus.WithError(err).Error("Failed to parse user")
		// 	return
		// }
		// article.UserID = user.ID
		err := models.AddArticle(article)
		if err != nil {
			logrus.WithError(err).Error("Failed to get article from user")
			return
		}
		code = utils.SUCCESS
	} else {
		code = utils.ERROR_EXIST_TAG
	}

	logrus.WithField("AddArticle", article).Info("Added article!")

	service.Response(c, http.StatusCreated, code, article)
}

//UpdateArticle .../api/v1/articles/{id}
func UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article := models.Article{}
	//从请求中获取参数
	if err := c.ShouldBind(&article); err != nil {
		logrus.WithError(err).Error("Failed to bind article to update")
		return
	}

	valid := validation.Validation{}
	if arg := article.State; arg != 0 && arg != 1 {
		valid.Range(arg, 0, 1, "state").Message("状态只允许0或1")
		return
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		err := models.UpdateArticle(article)
		if err != nil {
			code = utils.ERROR
			logrus.WithError(err).Error("Failed to update article")
			return
		}
		code = utils.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("article", article).Info("Updated article!")

	service.Response(c, http.StatusOK, code, "文章更新成功")
}

//DeleteArticle ...
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := utils.INVALID_PARAMS
	var err error
	if !valid.HasErrors() {
		err = models.DeleteArticle(id)
		if err != nil {
			logrus.WithError(err).Error("Failed to delete article")
			return
		}
		code = utils.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	service.Response(c, http.StatusOK, code, "文章删除成功！")
}
