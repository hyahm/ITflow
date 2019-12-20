package main

import (
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"itflow/bug"
	"itflow/bug/bugconfig"
	"itflow/db"
)

func main() {
	// 初始化配置文件
	goconfig.InitConf("bug.conf")

	//初始化mysql
	db.InitMysql()

	// 初始化redis
	db.InitRedis()
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
