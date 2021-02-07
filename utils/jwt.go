package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

//jwt加密密钥
var jwtKey = []byte("secret")

//Claims 是一些保存在token中的用户实体对象元数据，需要内嵌jwt的标准
type Claims struct {
	Username string
	jwt.StandardClaims
}

//GenerateToken 根据用户名和密码生成token
func GenerateToken(username string) (string, error) {
	//设置token过期时间，这里定为3小时
	expireTime := time.Now().Add(3 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Rehtt", //指定token发行人
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//获取完整带加密和签名的token
	token, err := tokenClaims.SignedString(jwtKey)
	if err != nil {
		logrus.WithError(err).Error("Error getting token")
		return "", err
	}
	return token, nil
}

//ParseToken 根据token值解析到claims对象信息
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权声明，方法内部主要是具体的解码和校验的过程，最终返回*jwt.Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if tokenClaims != nil {
		//从tokenClaims中获得Claims对象，使用断言将对象转换为需要的*Claims
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
