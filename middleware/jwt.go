package middleware

import (
	"net/http"
	"time"

	"github.com/ZRehtt/go-blog-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//JWT 身份验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := utils.SUCCESS
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			code = utils.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(authToken)
			if err != nil {
				code = utils.ERROR_AUTH_CHECK_TOKEN_FAIL
				logrus.WithError(err).Error("Failed to check token")
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = utils.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != utils.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  utils.GetMessage(code),
				"data": make(map[string]interface{}),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
