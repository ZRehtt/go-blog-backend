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

//GetTagByID .../api/v1/tags?id= 根据ID获取标签
func GetTagByID(c *gin.Context) {
	tagID, _ := strconv.Atoi(c.Query("id"))
	tag, err := models.GetTagByID(tagID)
	if err != nil {
		logrus.WithError(err).Error("Error getting tag by id")
		return
	}

	logrus.WithField("tag", *tag).Info("Getted tag by id")

	service.Response(c, http.StatusOK, 200, *tag)
}

//GetTags ...api/v1/tags 获取文章标签列表
func GetTags(c *gin.Context) {
	page := models.TagPage{}
	data := make(map[string]interface{})

	err := c.ShouldBind(&page)
	if err != nil {
		logrus.WithError(err).Error("Con't bind page of tag")
		return
	}

	code := utils.INVALID_PARAMS
	tags, err := models.GetTagsByPage(page.PageNumber, page.PageSize)
	if err != nil {
		code = utils.ERROR
		logrus.WithError(err).Error("Con't get tags by page")
		return
	}
	data["tags"] = tags

	count, err := models.GetTagTotal(page)
	if err != nil {
		code = utils.ERROR
		logrus.WithError(err).Error("Con't get count")
		return
	}
	data["count"] = count

	code = utils.SUCCESS

	logrus.WithFields(logrus.Fields{
		"data": data,
	}).Info("getted Tags!")

	service.Response(c, http.StatusOK, code, data)

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": code,
	// 	"msg":  utils.GetMessage(code),
	// 	"data": data,
	// })
}

//CreateTag 新建标签 /api/v1/tags
func CreateTag(c *gin.Context) {
	tag := models.Tag{}
	err := c.ShouldBind(&tag)
	if err != nil {
		logrus.WithError(err).Error("Error binding tag")
		return
	}

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		if ok, _ := models.ExistTagByName(tag.Name); !ok {
			code = utils.SUCCESS
			err = models.CreateTag(&tag)
			if err != nil {
				logrus.WithError(err).Error("Error add tag")
				return
			}
		} else {
			code = utils.ERROR_EXIST_TAG
		}
	}

	logrus.WithField("tag", tag).Info("Added Tag!")

	service.Response(c, http.StatusCreated, code, "标签添加成功！")
}

//UpdateTag 更新指定标签 /api/v1/tags/{id}
func UpdateTag(c *gin.Context) {
	tag := models.Tag{}
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.ShouldBind(&tag)
	if err != nil {
		logrus.WithError(err).Error("Error binding tag of updateTag")
		return
	}

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		code = utils.SUCCESS
		err = models.UpdateTag(id, tag)
		if err != nil {
			code = utils.ERROR
			logrus.WithError(err).Error("Error update tag")
			return
		}
	}

	logrus.WithField("tag", tag).Info("updated Tag!")

	service.Response(c, http.StatusOK, code, "标签更新成功！")
}

//DeleteTag 删除指定标签 /api/v1/tags/{id}
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		err := models.DeleteTag(id)
		if err != nil {
			code = utils.ERROR
			logrus.WithError(err).Error("Error delete tag")
			return
		}
		code = utils.SUCCESS
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	logrus.WithField("tagID", id).Info("deleted Tag!")

	service.Response(c, http.StatusOK, code, "标签删除成功！")
}
