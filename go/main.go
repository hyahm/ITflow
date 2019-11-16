package main

import (
	"fmt"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
	"itflow/bug"
	"itflow/bug/autodb"
	"itflow/bug/bugconfig"
	"log"
	"os"
	"os/signal"
)

func main() {
	// 初始化配置文件
	goconfig.InitConf("bug.conf")

	//初始化mysql
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
	// 初始化缓存（后面会使用redis）
	bugconfig.LoadConfig()
	// 初始化数据表
	if goconfig.ReadBool("initdb") {
		autodb.InitDb()
	}
	// 初始化日志
	golog.InitLogger(goconfig.ReadString("logpath"),
		goconfig.ReadInt64("logsize"),
		goconfig.ReadBool("logeveryday"))

	////
	bugservices()
}

func bugservices() {
	r := make(chan os.Signal, 0)
	signal.Notify(r, os.Interrupt, os.Kill)
	go bug.RunHttp()
	//go gareload.ListenReload(r)

	fmt.Println("exit code:", <-r)
	select {
	case <-r:
		fmt.Println("http")
	}
}
