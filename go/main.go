package main

import (
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
	// 初始化配置文件
	goconfig.InitConf("bug.ini", goconfig.INI)
	if goconfig.ReadBool("debug", false) {
		golog.Level = golog.ALL
	}
	// //初始化mysql
	db.InitMysql()
	// // // 初始化redis
	db.InitCacheTable()

	// // // 初始化缓存（后面会使用redis）
	cache.LoadConfig()

	// // // 初始化日志
	golog.InitLogger(goconfig.ReadString("log.path", ""),
		goconfig.ReadInt64("log.size", 0),
		goconfig.ReadBool("log.everyday", false))
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
