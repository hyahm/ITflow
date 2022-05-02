package routegroup

import (
	"itflow/handle/taskcontroller"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var TaskManager *xmux.RouteGroup

func init() {
	TaskManager = xmux.NewRouteGroup()
	// 创建bug页面
	TaskManager.Post("/bug/create", taskcontroller.Create).Bind(&model.Bug{}).
		AddModule(midware.JsonToStruct)

	TaskManager.Get("/bug/edit", taskcontroller.Get)
	TaskManager.Post("/bug/update", taskcontroller.Update).Bind(&model.Bug{}).
		AddModule(midware.JsonToStruct)

	TaskManager.Post("/bug/receive", taskcontroller.Receive).Bind(&model.Bug{}).
		AddModule(midware.JsonToStruct)

	TaskManager.Post("/bug/complete", taskcontroller.Complete).Bind(&model.Bug{}).
		AddModule(midware.JsonToStruct)
	// 1630339320
}
