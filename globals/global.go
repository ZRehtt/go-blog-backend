package globals

import (
	"github.com/ZRehtt/go-blog-backend/pkg/setting"
)

//声明全部变量，主要是和配置文件结合
var (
	Config string

	ServerSetting   *setting.ServerConfig
	AppSetting      *setting.AppConfig
	DatabaseSetting *setting.DatabaseConfig
	LogSetting      *setting.LogConfig
	JWTSetting      *setting.JWTConfig
)
