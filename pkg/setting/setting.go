package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Setting ...
type Setting struct {
	vp *viper.Viper
}

//NewSetting 配置初始化器
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchConfigChange()
	return s, nil
}

//WatchConfigChange 监听配置文件变化，重新读取配置数据
func (s *Setting) WatchConfigChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(e fsnotify.Event) {
			_ = s.ReadAllConfig()
		})
	}()
}
