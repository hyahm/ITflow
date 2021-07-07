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
	golog.Debugf("url: %s -- addr: %s -- method: %s -- exectime: %f\n", r.URL.Path, r.RemoteAddr, r.Method, time.Since(start).Seconds())
}

func RunHttp() {
	router := xmux.NewRouter()

	router.SetHeader("Access-Control-Allow-Origin", goconfig.ReadString("cross", "*"))
	router.SetHeader("Content-Type", "application/x-www-form-urlencoded,application/json; charset=UTF-8")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type,Access-Token,X-Token,smail,authorization")
	router.AddModule(midware.CheckToken) // 所有的请求头优先检查token是否有效， 登录除外
	if goconfig.ReadBool("debug") {
		// 打印所有接口的执行时间
		router.MiddleWare(GetExecTime)
	}
	router.AddGroup(routegroup.User) // 已完成
	router.AddGroup(routegroup.Doc)
	router.AddGroup(routegroup.Key)
	router.AddGroup(routegroup.Bug)
	router.AddGroup(routegroup.Search)
	router.AddGroup(routegroup.DashBoard)   // 后续
	router.AddGroup(routegroup.Version)     // 已完成
	router.AddGroup(routegroup.Admin)       // 已完成
	router.AddGroup(routegroup.Project)     // 已完成
	router.AddGroup(routegroup.Env)         // 已完成
	router.AddGroup(routegroup.StatusGroup) // 已完成
	router.AddGroup(routegroup.UserManager)

	router.AddGroup(routegroup.RoleGroup) // 已完成

	router.AddGroup(routegroup.Department)
	router.AddGroup(routegroup.Position) // 已完成
	router.AddGroup(routegroup.Status)   // 已完成

	router.AddGroup(routegroup.Level)        // 已完成
	router.AddGroup(routegroup.Important)    // 已完成
	router.AddGroup(routegroup.Log)          // 后面处理
	router.AddGroup(routegroup.UserGroup)    // 已完成
	router.AddGroup(routegroup.DefaultValue) // 已完成

	router.AddGroup(routegroup.Email)
	router.AddGroup(routegroup.Share)
	// router.AddGroup(routegroup.Api)

	router.Post("/uploadimg", handle.UploadImgs).DelModule(midware.CheckToken)
	router.Post("/upload/headimg", handle.UploadHeadImg)
	router.Get("/showimg/{imgname}", handle.ShowImg).SetHeader("Content-Type", "image/png").DelModule(midware.CheckToken)

	router.Get("/get/expire/{token}", handle.GetExpire).DelModule(midware.CheckToken)
	// 生产环境中， 这个要么注释， 要么
	if goconfig.ReadBool("debug", false) {
		doc := router.ShowApi("/docs").DelModule(midware.CheckToken)
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
