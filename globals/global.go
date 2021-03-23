package globals

import (
	"github.com/ZRehtt/go-blog-backend/pkg/setting"
	"gorm.io/gorm"
)

//声明全部变量，主要是和配置文件结合
var (
	GDB *gorm.DB

	ServerSetting *setting.ServerConfig
	AppSetting *setting.AppConfig
	DatabaseSetting *setting.DatabaseConfig
	JWTSetting *setting.JWTConfig
)
