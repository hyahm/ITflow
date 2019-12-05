package main

import (
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
	"itflow/bug"
	"itflow/bug/bugconfig"
	"log"
)

func main() {
	// 初始化配置文件
	goconfig.InitConf("bug.conf")

	//初始化mysql
	initMysql()
	// 初始化缓存（后面会使用redis）
	bugconfig.LoadConfig()
	// 初始化数据表, 避免数据错误， 请用sql导入
	//if goconfig.ReadBool("initdb") {
	//	autodb.InitDb()
	//}
	// 初始化日志
	golog.InitLogger(goconfig.ReadString("logpath"),
		goconfig.ReadInt64("logsize"),
		goconfig.ReadBool("logeveryday"))

	////
	bug.RunHttp()
}

func initMysql() {
	conf := &gomysql.Sqlconfig{
		DbName:   goconfig.ReadString("mysql.db"),
		Host:     goconfig.ReadString("mysql.host"),
		UserName: goconfig.ReadString("mysql.user"),
		Password: goconfig.ReadString("mysql.pwd"),
		Port:     goconfig.ReadInt("mysql.port"),
	}

	err := gomysql.SaveConf("bug", conf)
	if err != nil {
		log.Fatalln(err)
	}
}
