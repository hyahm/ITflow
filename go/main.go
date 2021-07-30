package main

import (
	_ "embed"
	"flag"
	"itflow/cache"
	"itflow/db"
	"itflow/httpserver"
	"log"
	"os"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

//go:embed bug.ini.default
var configBytes []byte

//go:embed bug.sql
var bugsql string

func main() {
	var make bool
	var conf string
	flag.BoolVar(&make, "c", false, "生成默认配置文件到当前路径")
	flag.StringVar(&conf, "f", "bug.ini", "默认配置文件路径")

	flag.Parse()
	if make {
		err := os.WriteFile(conf, configBytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(("配置文件已生成"))
		os.Exit(0)
	}
	defer golog.Sync()

	// 初始化配置文件
	goconfig.InitConf(conf, goconfig.INI)
	if goconfig.ReadBool("debug", false) {
		golog.Level = golog.DEBUG
	} else {
		golog.InitLogger(goconfig.ReadString("log.path", ""),
			goconfig.ReadInt64("log.size", 0),
			goconfig.ReadBool("log.everyday", false))
	}
	// //初始化mysql
	db.InitMysql(bugsql)
	// // // 初始化缓存表
	// db.InitCacheTable()
	cache.LoadConfig()

	// // // 初始化日志

	httpserver.RunHttp()
}
