package config

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
	"gopkg.in/redis.v5"
)

var (
	config GlobalConfig
)

//GlobalConfig 全局配置
type GlobalConfig struct {
	HTTPBind          string
	Mysql             MysqlConfig
	Redis             RedisConfig
}

// Config 返回配置文件
func Config() GlobalConfig {
	return config
}

// RedisConfig Redis配置参数
type RedisConfig struct {
	RedisAddr string
	RedisDb   int
}

// CreateRedis 初始化Redis组件
func CreateRedis(config RedisConfig) (*redis.Client, error) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
		DB:   config.RedisDb,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RedisClient, nil
}

// MysqlConfig mysql配置参数
type MysqlConfig struct {
	DbDsn   string
	ShowSQL bool
}

// CreateMysql 初始化数据库组件
func CreateMysql(config MysqlConfig) (*xorm.Engine, error) {
	mysql, err := xorm.NewEngine("mysql", config.DbDsn)
	if err != nil {
		return nil, err
	}
	mysql.ShowSQL(config.ShowSQL)
	return mysql, nil
}

// ParseConfig 解析配置文件
func ParseConfig(cfg string) {
	viper.SetConfigFile(cfg)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}
