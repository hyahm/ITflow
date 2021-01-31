package main

import (
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/httpserver"
	"os"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

func main() {
	defer golog.Sync()
	conf := "bug.ini"
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	// 初始化配置文件
	goconfig.InitConf(conf, goconfig.INI)
	if goconfig.ReadBool("debug", false) {
		golog.Level = golog.DEBUG
	}
	// //初始化mysql
	db.InitMysql()
	// // // 初始化缓存表
	// db.InitCacheTable()
	cache.LoadConfig()

	// // // 初始化日志
	if !goconfig.ReadBool("debug", false) {
		golog.InitLogger(goconfig.ReadString("log.path", ""),
			goconfig.ReadInt64("log.size", 0),
			goconfig.ReadBool("log.everyday", false))
	}
	httpserver.RunHttp()
}
