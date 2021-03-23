package main

import (
	"flag"
	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/internal/models"
	"github.com/ZRehtt/go-blog-backend/internal/routers"
	"github.com/ZRehtt/go-blog-backend/pkg/logger"
	"github.com/ZRehtt/go-blog-backend/pkg/setting"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

var (
	config string
	mode string
)

//
func init() {
	//启动自定义日志
	logger.NewLogger()
	logrus.WithField("logger", "logger").Info("Logger is ready!")

	err := setupFlag()
	if err != nil {
		logrus.Printf("init setupFlag with error: %v\n", err)
		return
	}

	err = setupSetting()
	if err != nil {
		logrus.Printf("init setupSetting with error: %v\n", err)
		return
	}

	err = models.NewDatabase(globals.DatabaseSetting)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to connect database!")
		return
	}
}

func main() {
	router := routers.NewRouter()

	server := &http.Server{
		Addr: ":" + globals.ServerSetting.Port,
		Handler: router,
		ReadTimeout: globals.ServerSetting.ReadTimeout * time.Second, //允许读取的最大时间
		WriteTimeout: globals.ServerSetting.WriteTimeout * time.Second, //允许写入的最大时间
		MaxHeaderBytes: 1 << 20, //请求头的最大字节数
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithError(err).Error("Failed to listen server!")
	}
}


func setupFlag() error {
	flag.StringVar(&config, "config", "conf/", "指定要使用的配置文件路径")
	flag.StringVar(&mode, "mode", "", "应用启动模式")
	flag.Parse()
	return nil
}

func setupSetting() error {
	set, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	err = set.ReadConfig("Server", &globals.ServerSetting)
	if err != nil {
		return err
	}

	err = set.ReadConfig("App", &globals.AppSetting)
	if err != nil {
		return err
	}

	err = set.ReadConfig("Database", &globals.DatabaseSetting)
	if err != nil {
		return err
	}

	err = set.ReadConfig("JWT", &globals.JWTSetting)

	return nil
}

