package routegroup

import (
	"itflow/handle/taskcontroller"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var TaskManager *xmux.RouteGroup

func init() {
	TaskManager = xmux.NewRouteGroup()
	// 创建bug页面
	TaskManager.Post("/bug/create", taskcontroller.Create).BindJson(&model.Bug{})
	TaskManager.Get("/bug/edit", taskcontroller.Get)
	TaskManager.Post("/bug/update", taskcontroller.Update).BindJson(&model.Bug{})
	TaskManager.Post("/bug/receive", taskcontroller.Receive).BindJson(&model.Bug{})

	TaskManager.Post("/bug/complete", taskcontroller.Complete).BindJson(&model.Bug{})
	// 1630339320
}
