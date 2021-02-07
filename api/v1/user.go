package v1

import (
	"net/http"

	"github.com/ZRehtt/go-blog-backend/models"
	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Register 注册
func Register(c *gin.Context) {
	requestUser := models.User{}
	err := c.ShouldBind(&requestUser)
	if err != nil {
		logrus.WithError(err).Error("Error binding user of register")
		return
	}

	//获取用户信息
	username := requestUser.Username
	password := requestUser.Password

	if len(username) == 0 || len(password) == 0 {
		logrus.Info("username and pssword are required")
		return
	}
	//判断用户是否已存在
	if ok := models.CheckUserByName(username); ok {
		logrus.Info("User already exists.")
		return
	}
	if ok := models.CreateUser(username, password); !ok {
		logrus.WithError(err).Error("Error create user")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg": "用户注册成功",
	})
}

//
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	user := models.User{Username: username, Password: password}
	valid := validation.Validation{}
	data := make(map[string]interface{})
	code := utils.INVALID_PARAMS
	ok, _ := valid.Valid(&user)
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := utils.GenerateToken(username)
			if err != nil {
				code = utils.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = utils.SUCCESS
			}
		} else {
			code = utils.ERROR_AUTH
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  utils.GetMessage(code),
		"data": data,
	})
}

//
func CreateUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if ok := models.CheckUserByName(username); !ok {
		models.CreateUser(username, password)
	} else {
		logrus.WithField("user", username).Info("Username has been used")
	}
	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "用户创建成功",
	})
}
