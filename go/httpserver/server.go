package httpserver

import (
	"fmt"
	"itflow/handle"
	"itflow/midware"
	"itflow/routegroup"
	"net/http"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/xmux"
)

func RunHttp() {
	router := xmux.NewRouter()
	router.SetHeader("Access-Control-Allow-Origin", "*")
	router.SetHeader("Content-Type", "application/x-www-form-urlencoded,application/json; charset=UTF-8")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type,Access-Token,X-Token,smail")
	router.AddMidware(midware.CheckToken) // 所有的请求头优先检查token是否有效， 登录除外

	router.AddGroup(routegroup.User)
	router.AddGroup(routegroup.Bug)
	router.AddGroup(routegroup.Search)
	router.AddGroup(routegroup.DashBoard) // 后续
	router.AddGroup(routegroup.Version)
	router.AddGroup(routegroup.Admin) // 已完成
	router.AddGroup(routegroup.Project)
	router.AddGroup(routegroup.Env)
	router.AddGroup(routegroup.StatusGroup) // 已完成
	router.AddGroup(routegroup.UserManager)

	router.AddGroup(routegroup.RoleGroup)

	router.AddGroup(routegroup.Department)
	router.AddGroup(routegroup.Position)
	router.AddGroup(routegroup.Status) // 已完成

	router.AddGroup(routegroup.Type)
	router.AddGroup(routegroup.Level)
	router.AddGroup(routegroup.Important)
	router.AddGroup(routegroup.Log)          // 后面处理
	router.AddGroup(routegroup.UserGroup)    // 已完成
	router.AddGroup(routegroup.DefaultValue) // 已完成

	router.AddGroup(routegroup.Email)
	router.AddGroup(routegroup.Share)
	// router.AddGroup(routegroup.Api)

	router.Post("/uploadimg", handle.UploadImgs).DelMidware(midware.CheckToken)
	router.Post("/upload/headimg", handle.UploadHeadImg)
	router.Get("/showimg/{imgname}", handle.ShowImg).SetHeader("Content-Type", "image/png").DelMidware(midware.CheckToken)

	router.Get("/get/expire/{token}", handle.GetExpire).DelMidware(midware.CheckToken)
	// 生产环境中， 这个要么注释， 要么
	if goconfig.ReadBool("debug", false) {
		doc := router.ShowApi("/docs").DelMidware(midware.CheckToken)
		router.AddGroup(doc)
	}

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
