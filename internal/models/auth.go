package models

//Auth 权限管理，用于保存签发的JWT凭证
type Auth struct {
	*Model
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}
