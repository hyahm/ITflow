package bug

import (
	"fmt"
	"time"

	"itflow/bug/handle"
	"itflow/bug/route"
	"net/http"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/xmux"
)

func RunHttp() {

	router := xmux.NewRouter()
	router.SetHeader("Access-Control-Allow-Origin", "*")
	router.SetHeader("Content-Type", "application/x-www-form-urlencoded,application/json; charset=UTF-8")
	router.SetHeader("Access-Control-Allow-Headers", "Content-Type,Access-Token,X-Token,smail")
	router.AddGroup(route.User)
	router.AddGroup(route.Bug)

	router.AddGroup(route.Version)
	router.AddGroup(route.Admin)
	router.AddGroup(route.Project)
	router.AddGroup(route.Env)

	router.AddGroup(route.Role)
	router.AddGroup(route.Department)
	router.AddGroup(route.Position)
	router.AddGroup(route.Status)

	router.AddGroup(route.Email)
	router.AddGroup(route.Share)
	router.AddGroup(route.Api)

	router.Pattern("/uploadimg").Post(handle.UploadImgs)
	router.Pattern("/upload/headimg").Post(handle.UploadHeadImg)
	router.Pattern("/showimg/{imgname}").Get(handle.ShowImg)
	//
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
			print("has ssl key?")
			panic(err)
		}
	} else {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}

}
