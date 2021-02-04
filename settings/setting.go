package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//NewViper 初始化viper配置
func NewViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/config/")
	viper.AddConfigPath("./conf")

	//监视配置文件变化，重新读取配置数据
	//Viper在运行时拥有读取配置文件的能力。
	//只需要调用viper实例的WatchConfig函数，也可以指定一个回调函数来获得变动的通知
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	//搜索路径，并读取配置数据
	return viper.ReadInConfig()
}
