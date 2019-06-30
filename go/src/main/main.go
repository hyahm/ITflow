package main

import (
	"bug"
	"bug/autodb"
	"bug/bugconfig"
	"fmt"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"os"
)

func main() {

	//testimg.TestDoc()
	goconfig.InitConf("bug.conf")

	autodb.InitDb()

	bugconfig.LoadConfig()
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
