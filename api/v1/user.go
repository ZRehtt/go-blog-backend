package v1

import (
	"net/http"

	"github.com/ZRehtt/go-blog-backend/models"
	"github.com/ZRehtt/go-blog-backend/service"
	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Register 注册
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}

	valid.Required(username, "username").Message("用户名不能为空")
	valid.Required(password, "password").Message("密码不能为空")
	valid.MinSize(password, 8, "password").Message("密码最少为8个字符")

	code := utils.INVALID_PARAMS
	if !valid.HasErrors() {
		//判断用户是否已存在
		if ok := models.CheckUserByName(username); !ok {
			//用户名不存在时再将密码加密
			passwordHash, err := models.HashPassword(password)
			if err != nil {
				logrus.WithError(err).Error("Failed to hash password")
				return
			}
			err = models.CreateUser(models.User{Username: &username, PasswordHash: &passwordHash})
			if err != nil {
				logrus.WithError(err).Error("Failed to create user when registering")
				return
			}
			code = utils.SUCCESS
		} else {
			service.Response(c, http.StatusConflict, code, "用户已存在!")
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Infof("err.key: %s, err.message: %s\n", err.Key, err.Message)
		}
	}

	service.Response(c, http.StatusCreated, code, "用户注册成功！")

}

//
// func GetAuth(c *gin.Context) {
// 	username := c.Query("username")
// 	password := c.Query("password")

// 	user := models.User{Username: username, Password: password}
// 	valid := validation.Validation{}
// 	data := make(map[string]interface{})
// 	code := utils.INVALID_PARAMS
// 	ok, _ := valid.Valid(&user)
// 	if ok {
// 		isExist := models.CheckAuth(username, password)
// 		if isExist {
// 			token, err := utils.GenerateToken(username)
// 			if err != nil {
// 				code = utils.ERROR_AUTH_TOKEN
// 			} else {
// 				data["token"] = token
// 				code = utils.SUCCESS
// 			}
// 		} else {
// 			code = utils.ERROR_AUTH
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": code,
// 		"msg":  utils.GetMessage(code),
// 		"data": data,
// 	})
// }
