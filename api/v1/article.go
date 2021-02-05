package v1

import (
	"net/http"
	"strconv"

	"github.com/ZRehtt/go-blog-backend/models"
	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//GetArticleByID /api/v1/articles/{id}
func GetArticleByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := utils.INVALID_PARAMS
	var data interface{}
	var err error
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data, err = models.GetArticle(id)
			if err != nil {
				logrus.WithError(err).Error("Failed to get article")
				return
			}
			code = utils.SUCCESS
		} else {
			code = utils.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("id", id).Info("getted Article!")

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": data,
	})
}

//GetArticles ...
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}
	state := -1
	var err error
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			logrus.WithError(err).Error("Con't set state")
			return
		}
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	tagID := -1
	if arg := c.Query("tag_id"); arg != "" {
		tagID, err = strconv.Atoi(arg)
		if err != nil {
			logrus.WithError(err).Error("Con't set tagID")
			return
		}
		maps["tag_id"] = tagID
		valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		code = utils.SUCCESS
		data["lists"], _ = models.GetArticles(utils.GetPage(c), viper.GetInt("app.page_size"), maps)
		data["total"], _ = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("tag_id", tagID).Info("getted articles!")

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": data,
	})
}

//AddArticle ...
func AddArticle(c *gin.Context) {
	tagID, _ := strconv.Atoi(c.Query("tag_id"))
	title := c.Query("title")
	description := c.Query("description")
	content := c.Query("content")
	state, _ := strconv.Atoi(c.DefaultQuery("state", "0"))
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Min(tagID, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(description, "description").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		if ok, _ := models.ExistTagByID(tagID); !ok {
			code = utils.SUCCESS
			data := make(map[string]interface{})
			data["tag_id"] = tagID
			data["title"] = title
			data["description"] = description
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
		} else {
			code = utils.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("tag_id", tagID).Info("Added article!")

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": make(map[string]string),
	})
}

//UpdateArticle ...
func UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tagID, _ := strconv.Atoi(c.Query("tag_id"))
	title := c.Query("title")
	description := c.Query("description")
	content := c.Query("content")
	updatedBy := c.Query("updated_by")

	valid := validation.Validation{}
	state := -1
	var err error
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			logrus.WithError(err).Error("Con't set state")
			return
		}
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(description, 355, "description").Message("描述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.MaxSize(updatedBy, 100, "updated_by").Message("修改人最长为100字符")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistArticleByID(id) {
			if isTagID, _ := models.ExistTagByID(tagID); !isTagID {
				data := make(map[string]interface{})
				if tagID > 0 {
					data["tag_id"] = tagID
				}
				if title != "" {
					data["title"] = title
				}
				if description != "" {
					data["description"] = description
				}
				if content != "" {
					data["content"] = content
				}
				data["updated_by"] = updatedBy
				models.UpdateArticle(id, data)
				code = utils.SUCCESS
			} else {
				code = utils.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = utils.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("id", id).Info("Updated article!")

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": make(map[string]string),
	})
}

//DeleteArticle ...
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := utils.INVALID_PARAMS
	var err error
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			err = models.DeleteArticle(id)
			if err != nil {
				logrus.WithError(err).Error("Failed to delete article")
				return
			}
			code = utils.SUCCESS
		} else {
			code = utils.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": make(map[string]string),
	})
}
