package bootstrap

import (
	"MySystem/config"
	"MySystem/daos"
)

//Bootstrap 配置参数，启动相关组件
func Bootstrap(cfg string) {
	parseConfig(cfg)
	initMysql()
	//initRedis()
}

func Destory() {
	daos.CloesMySQL()
}
func parseConfig(cfg string) {
	config.ParseConfig(cfg)
}

func initMysql() {
	engine, err := config.CreateMysql(config.Config().Mysql)
	if err != nil {
		panic(err)
	}
	daos.SetMySql(engine)
}

func initRedis() {
	rc, err := config.CreateRedis(config.Config().Redis)
	if err != nil {
		panic(err)
	}
	daos.SetRedis(rc)
}