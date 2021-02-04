package main

import (
	"net/http"

	"github.com/ZRehtt/go-blog-backend/api"
	"github.com/ZRehtt/go-blog-backend/models"
	"github.com/ZRehtt/go-blog-backend/settings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//这里用logrus简单记录日志，设置日志记录级别为Debug，只调试使用
	logrus.SetLevel(logrus.DebugLevel)

	if err := settings.NewViper(); err != nil {
		logrus.WithError(err).Error("Failed to read config file!")
	}

	if err := models.NewDatabase(); err != nil {
		logrus.WithError(err).Error("Failed to connect database!")
	}
	logrus.Debug("database is ready to use.")

	router := api.NewRouter()

	server := &http.Server{
		Addr:           ":" + viper.GetString("server.port"),
		Handler:        router,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Failed to listen server!")
	}
}
