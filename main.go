package main

import (
	"net/http"
	"time"

	"github.com/ZRehtt/go-blog-backend/api"
	"github.com/ZRehtt/go-blog-backend/models"
	"github.com/ZRehtt/go-blog-backend/settings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//日志配置
	settings.NewLogger()
	logrus.WithField("logger", "logger").Info("Logger is ready!")

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
		ReadTimeout:    time.Second * time.Duration(viper.GetInt("server.read_timeout")),  //允许读取的最大时间
		WriteTimeout:   time.Second * time.Duration(viper.GetInt("server.write_timeout")), //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                                                           //请求头的最大字节数
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Failed to listen server!")
	}
}
