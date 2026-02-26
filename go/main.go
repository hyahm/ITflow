package main

import (
	_ "embed"
	"flag"
	"itflow/cache"
	"itflow/db"
	"itflow/httpserver"
	"itflow/model"
	"log"
	"os"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

//go:embed bug.ini
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
	goconfig.InitConf(conf)
	if goconfig.ReadBool("debug", false) {

		golog.InitLogger(goconfig.ReadEnv("LOG_PATH", goconfig.ReadString("log.path", "")),
			goconfig.ReadInt64("log.size", 0),
			goconfig.ReadBool("log.everyday", false))
	}

	// //初始化mysql
	// switch goconfig.ReadString("db.driver") {
	// case "postgres":
	// 	db.InitPgDatabase(bugsql)
	// case "mysql":
	db.InitMysqlDatabase(bugsql)
	// default:
	// log.Fatalf("不支持的数据库类型: %s", goconfig.ReadString("db.driver"))
	// }

	// // // 初始化缓存表
	// db.InitCacheTable()
	cache.LoadConfig()
	model.InitPagePermId() // 初始化角色id
	// model.InitCache()
	// // // 初始化日志

	httpserver.RunHttp()
}

// 5  3   18
