package middleware

import (
	"strings"

	"github.com/ZRehtt/go-blog-backend/pkg/app"
	"github.com/ZRehtt/go-blog-backend/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//JWTAuth 身份验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ecode := errcode.Success
		//从请求头Header的Authorization中找出token，但获取的是Bearer token的结构，需要分割
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			ecode = errcode.ErrorInvalidParams
		}
		//按空格分割以获取完整的token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ecode = errcode.ErrorInvalidParams
		}

		//token信息保存在parts[1]中
		_, err := app.ParseToken(parts[1])
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				ecode = errcode.ErrorUnauthorizedTokenTimeout
			default:
				ecode = errcode.ErrorUnauthorizedToken
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}

		//可将认证信息保存到请求上下文
		//后续处理函数可以使用c.Get("")获取当前用户信息
		//c.Set("", tokenClaims.Issuer)
		c.Next()
	}
}
