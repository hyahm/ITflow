package main

import (
	"bug"
	"bug/autodb"
	"bug/bugconfig"
	"fmt"
	"gaconfig"
	"galog"
	"os"
)

func main() {

	//testimg.TestDoc()
	gaconfig.InitConf("bug.conf")

	autodb.InitDb()

	bugconfig.LoadConfig()
	galog.InitLogger()

	////
	galog.Info("this is an info log")
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
