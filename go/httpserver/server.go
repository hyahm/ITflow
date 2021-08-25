package httpserver

import (
	"fmt"
	"itflow/handle"
	"itflow/midware"
	"itflow/routegroup"
	"net/http"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetExecTime(handle func(http.ResponseWriter, *http.Request), w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	handle(w, r)
	golog.Infof("url: %s -- addr: %s -- method: %s -- exectime: %f", r.URL.Path, r.RemoteAddr, r.Method, time.Since(start).Seconds())
}

func RunHttp() {
	router := xmux.NewRouter()

	router.SetHeader("Access-Control-Allow-Origin", goconfig.ReadString("cross", "*"))
	router.SetHeader("Content-Type", "application/x-www-form-urlencoded,application/json; charset=UTF-8")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type,Access-Token,X-Token,smail,authorization")
	router.AddModule(midware.CheckToken) // 所有的请求头优先检查token是否有效， 登录除外
	// if goconfig.ReadBool("debug") {
	// 	// 打印所有接口的执行时间
	// 	router.MiddleWare(GetExecTime)
	// }

	// 面板菜单
	router.AddGroup(routegroup.DashBoard)
	// 系统设置菜单
	router.AddGroup(routegroup.SystemSetting)
	// 用户管理菜单
	router.AddGroup(routegroup.UserManager)
	// 任务管理菜单

	router.AddGroup(routegroup.TaskManager)
	// 文档管理菜单
	router.AddGroup(routegroup.Document)
	// 设置中心菜单
	router.AddGroup(routegroup.SettingCenter)
	// 公共部分接口
	router.AddGroup(routegroup.Public)

	router.AddGroup(routegroup.User)
	router.AddGroup(routegroup.Bug)
	router.AddGroup(routegroup.Search)

	router.AddGroup(routegroup.Admin) // 已完成

	router.AddGroup(routegroup.Department)

	router.AddGroup(routegroup.Share)
	// router.AddGroup(routegroup.Api)

	router.Post("/uploadimg", handle.UploadImgs).DelModule(midware.CheckToken)

	router.Get("/showimg/{string:imgname}", handle.ShowImg).SetHeader("Content-Type", "image/png").DelModule(midware.CheckToken)

	router.Get("/get/expire/{token}", handle.GetExpire).DelModule(midware.CheckToken)
	// // 生产环境中， 这个要么注释， 要么
	// if goconfig.ReadBool("debug", false) {
	// 	doc := router.ShowApi("/docs").DelModule(midware.CheckToken)
	// 	router.AddGroup(doc)
	// }
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world"))
	}).DelModule(midware.CheckToken)
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
