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
)

func main() {

	goconfig.InitConf("bug.conf")
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
	bugconfig.LoadConfig()
	if goconfig.ReadBool("initdb") {
		autodb.InitDb()
	}

	golog.InitLogger(goconfig.ReadString("logpath"),
		goconfig.ReadInt64("logsize"),
		goconfig.ReadBool("logeveryday"))

	////
	golog.Info("this is an info log")
	bugservices()
}

func bugservices() {
	r := make(chan os.Signal, 0)

	go bug.RunHttp(r)
	//go gareload.ListenReload(r)

	fmt.Println("exit code:", <-r)
	select {
	case <-r:
		fmt.Println("http")
	}
}
