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

//GetTags ...api/v1/tags?name=(test)&state=1 获取文章标签列表
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	state := -1
	var err error
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			logrus.WithError(err).Error("Con't set state")
			return
		}
		maps["state"] = state
	}

	code := utils.SUCCESS
	pageNumber := utils.GetPage(c)
	pageSize := viper.GetInt("app.page_size")
	tags, err := models.GetTags(pageNumber, pageSize, maps)
	if err != nil {
		logrus.WithError(err).Error("Error getting tags")
		return
	}
	data["lists"] = tags
	total, err := models.GetTagTotal(maps)
	if err != nil {
		logrus.WithError(err).Error("Error getting tag total")
		return
	}
	data["total"] = total

	logrus.WithFields(logrus.Fields{
		"maps": maps,
		"data": data,
	}).Info("getted Tags!")

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": data,
	})
}

//AddTag 新建标签
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state, _ := strconv.Atoi(c.DefaultQuery("state", "0"))
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		if ok, _ := models.ExistTagByName(name); !ok {
			code = utils.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = utils.ERROR_EXIST_TAG
		}
	}

	logrus.WithField("tag", name).Info("Added Tag!")

	c.JSON(http.StatusCreated, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": make(map[string]string),
	})
}

//UpdateTag 更新指定标签
func UpdateTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.Query("name")
	updatedBy := c.Query("updated_by")

	valid := validation.Validation{}
	var state = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(updatedBy, "modifiupdated_by").Message("修改人不能为空")
	valid.MaxSize(updatedBy, 100, "updated_by").Message("修改人最长为100字符")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		code = utils.SUCCESS
		if ok, _ := models.ExistTagByID(id); ok {
			data := make(map[string]interface{})
			data["updated_by"] = updatedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.UpdateTag(id, data)
		} else {
			code = utils.ERROR_EXIST_TAG
		}
	}

	logrus.WithField("tag", name).Info("updated Tag!")

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": make(map[string]string),
	})
}

//DeleteTag 删除指定标签
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		code = utils.SUCCESS
		if ok, _ := models.ExistTagByID(id); ok {
			models.DeleteTag(id)
		} else {
			code = utils.ERROR_NOT_EXIST_TAG
		}
	}

	logrus.WithField("tagID", id).Info("deleted Tag!")

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": make(map[string]string),
	})
}
