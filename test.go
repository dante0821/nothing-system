package main

import (
	"MySystem/models"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func main1() {

	//数据库连接参数

	params := fmt.Sprintf("root:szc910821@tcp(127.0.0.1:3306)/mybase?charset=utf8")

	var err error

	//连接数据库

	engine, err = xorm.NewEngine("mysql", params)

	if err != nil {

		panic(err)

	}

	//添加统一前缀

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")

	engine.SetTableMapper(tbMapper)

	defer engine.Close()

	//创建表

	err = engine.Sync2(new(models.PersonInfo))

	if err != nil {

		panic(err)

	}

}