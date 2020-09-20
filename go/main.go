package main

import (
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

func main() {
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
	db.InitCacheTable()
	cache.LoadConfig()

	// // // 初始化日志
	if !goconfig.ReadBool("debug", false) {
		golog.InitLogger(goconfig.ReadString("log.path", ""),
			goconfig.ReadInt64("log.size", 0),
			goconfig.ReadBool("log.everyday", false))
	}

	// ////
	signalChan := make(chan os.Signal)

	go func() {
		//阻塞程序运行，直到收到终止的信号
		<-signalChan
		log.Println("Cleaning before stop...")
		err := db.SaveCacheTable()
		if err != nil {
			log.Println(err)
		}
		golog.Info("save successed")
		os.Exit(0)
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	httpserver.RunHttp()
}
