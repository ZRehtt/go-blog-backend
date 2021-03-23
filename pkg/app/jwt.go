package app

import (
	"github.com/dgrijalva/jwt-go"
)

//Claims 自定义认证信息，包含密钥，签发人
type Claims struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	jwt.StandardClaims
}

//GenerateToken 生成Token
func GenerateToken(appKey, appSecret string) (string, error) {
	//nowTime := time.Now()
	return "", nil
}

//ParseToken 解析和检验Token
func ParseToken(token string) (*Claims, error) {
	return nil, nil
}
