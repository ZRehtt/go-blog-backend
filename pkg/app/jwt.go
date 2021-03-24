package app

import (
	"time"

	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

//MyClaims 自定义认证信息，包含密钥，签发人
type MyClaims struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(globals.JWTSetting.Secret)
}

//GenerateToken 生成Token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(globals.JWTSetting.ExpiresAt)
	claims := MyClaims{
		AppKey:    utils.MD5(appKey),
		AppSecret: utils.MD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    globals.JWTSetting.Issuer,
		},
	}

	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//生成完整的签名字符串
	token, err := tokenClaims.SignedString(GetJWTSecret())
	if err != nil {
		logrus.WithError(err).Error("failed to get token!")
		return "", err
	}
	return token, nil
}

//ParseToken 解析和检验Token
func ParseToken(token string) (*MyClaims, error) {
	//用于解析鉴权声明，方法内部主要是具体的解码和校验的过程，最终返回*jwt.Token
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		logrus.WithError(err).Error("Failed to parse token with claim!")
		return nil, err
	}
	if tokenClaims != nil {
		//从tokenClaims中获得Claims对象，使用断言将对象转换为需要的*Claims
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
