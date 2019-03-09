package config

import (
	"git.bitboolean.com/pay/misc/driver"
	"git.bitboolean.com/pay/misc/utils"
)

var (
	config GlobalConfig
)

//GlobalConfig 全局配置
type GlobalConfig struct {
	HTTPBind          string
	Mysql             driver.MysqlConfig
	Redis             driver.RedisConfig
}

// Config 返回配置文件
func Config() GlobalConfig {
	return config
}

// ParseConfig 解析配置文件
func ParseConfig(cfg string) {
	util.ParseConfig(cfg, &config)
}
