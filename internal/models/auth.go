package models

import "gorm.io/gorm"

//Auth 权限管理，用于保存签发的JWT凭证
type Auth struct {
	*Model
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

//GetAuth 获取单个认证信息
func GetAuth(a Auth) (*Auth, error) {
	var auth Auth
	db := db.Where("app_key = ? AND app_secret = ? AND is_deleted = ?", a.AppKey, a.AppSecret, 0)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &auth, nil
}
