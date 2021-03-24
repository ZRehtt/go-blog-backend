package setting

import "time"

/*
定义全局变量，通过viper读完配置反序列化到变量中，其他模块通过全局变量获取配置信息
和配置文件一一对应
*/

//Config 全局配置
type Config struct {
	*AppConfig      `mapstructure:"app"`
	*ServerConfig   `mapstructure:"server"`
	*DatabaseConfig `mapstructure:"database"`
	*JWTConfig      `mapstructure:"JWT"`
}

//AppConfig 应用相关配置信息
type AppConfig struct {
	PageSize int    `yaml:"page_size"`
	Mode     string `yaml:"mode"`
	Host     string `yaml:"host"`
	Version  string `yaml:"version"`
}

//ServerConfig 服务相关配置信息
type ServerConfig struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

//DatabaseConfig 数据库配置信息
type DatabaseConfig struct {
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	Type         string `yaml:"type"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         string `yaml:"port"`
	DBName       string `yaml:"dbname"`
	Config       string `yaml:"config"`
}

//JWTConfig jwt令牌相关配置
type JWTConfig struct {
	Secret    string        `yaml:"secret"`
	Issuer    string        `yaml:"issuer"`
	ExpiresAt time.Duration `yaml:"expires_at"`
}

var configs = make(map[string]interface{})

//ReadConfig 用于读取配置结构体的配置信息
func (s *Setting) ReadConfig(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	if _, ok := configs[k]; !ok {
		configs[k] = v
	}
	return nil
}

//ReadAllConfig ...
func (s *Setting) ReadAllConfig() error {
	for k, v := range configs {
		err := s.ReadConfig(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
