package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ZRehtt/go-blog-backend/globals"
	"github.com/ZRehtt/go-blog-backend/internal/models"
	"github.com/ZRehtt/go-blog-backend/internal/routers"
	"github.com/ZRehtt/go-blog-backend/pkg/logger"
	"github.com/ZRehtt/go-blog-backend/pkg/setting"
	"go.uber.org/zap"
)

var (
	config string
	mode   string
)

//
func init() {
	err := setupFlag()
	if err != nil {
		fmt.Printf("failed to init setup flag: %v\n", err)
		return
	}
	if err = setupSetting(); err != nil {
		fmt.Printf("failed to init setting: %v\n", err)
		return
	}
	if err = logger.NewLogger(globals.LogSetting); err != nil {
		fmt.Printf("failed to init logger: %v\n", err)
		return
	}
	if err = models.NewDatabase(globals.DatabaseSetting); err != nil {
		zap.L().Error("failed to init database", zap.Any("err", err))
		return
	}
}

func main() {
	router := routers.NewRouter()

	server := &http.Server{
		Addr:           ":" + globals.ServerSetting.Port,
		Handler:        router,
		ReadTimeout:    globals.ServerSetting.ReadTimeout * time.Second,  //允许读取的最大时间
		WriteTimeout:   globals.ServerSetting.WriteTimeout * time.Second, //允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                                          //请求头的最大字节数
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.L().Error("Failed to listen server!", zap.Any("err", err))
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

	err = set.ReadConfig("log", &globals.LogSetting)
	if err != nil {
		return err
	}

	err = set.ReadConfig("JWT", &globals.JWTSetting)
	if err != nil {
		return err
	}

	return nil
}
