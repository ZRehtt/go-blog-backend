package service

import (
	"errors"

	"github.com/ZRehtt/go-blog-backend/internal/models"
	"github.com/sirupsen/logrus"
)

//AuthRequest 认证信息接口入参校验
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

//CheckAuth 检验auth是否已存在
func (s *Service) CheckAuth(param *AuthRequest) error {
	auth, err := models.New(s.db).GetAuth(models.Auth{AppKey: param.AppKey, AppSecret: param.AppSecret})
	if err != nil {
		logrus.WithError(err).Error("failed to get auth by request.")
		return err
	}
	if auth.ID > 0 {
		//说明auth已存在
		return nil
	}
	return errors.New("auth info does not exist")
}
