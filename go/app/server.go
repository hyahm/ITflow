package app

import (
	"fmt"
	"time"

	"itflow/app/handle"
	"itflow/midware"
	"itflow/route"
	"net/http"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/xmux"
)

func RunHttp() {

	router := xmux.NewRouter()
	router.SetHeader("Access-Control-Allow-Origin", "*")
	router.SetHeader("Content-Type", "application/x-www-form-urlencoded,application/json; charset=UTF-8")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type,Access-Token,X-Token,smail")
	router.AddMidware(midware.CheckToken) // 所有的请求头优先检查token是否有效， 登录除外

	router.AddGroup(route.User)
	router.AddGroup(route.Bug)

	router.AddGroup(route.Version)
	router.AddGroup(route.Admin)
	router.AddGroup(route.Project)
	router.AddGroup(route.Env)

	router.AddGroup(route.Role).AddGroup(route.Department).AddGroup(route.Position).AddGroup(route.Status)

	router.AddGroup(route.Type)
	router.AddGroup(route.Level).AddGroup(route.Important).AddGroup(route.Log).AddGroup(route.UserGroup)

	router.AddGroup(route.Email).AddGroup(route.Share).AddGroup(route.Api)

	router.Pattern("/uploadimg").Post(handle.UploadImgs)
	router.Pattern("/upload/headimg").Post(handle.UploadHeadImg)
	router.Pattern("/showimg/{imgname}").Get(handle.ShowImg)
	//
	doc := xmux.ShowApi("/docs", router).DelMidware(midware.CheckToken)
	router.AddGroup(doc)
	listenaddr := goconfig.ReadString("listenaddr", ":10001")

	fmt.Println("listen on " + listenaddr)
	server := http.Server{
		Addr:         listenaddr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	if goconfig.ReadBool("ssl", false) {
		if err := server.ListenAndServeTLS(goconfig.ReadString("certfile", "ssl.cert"), goconfig.ReadString("keyfile", "ssl.key")); err != nil {
			panic(err)
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}

}
