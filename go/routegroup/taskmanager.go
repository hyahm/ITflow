package routegroup

import (
	"itflow/handle/taskcontroller"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var TaskManager *xmux.GroupRoute

func init() {
	TaskManager = xmux.NewGroupRoute()
	// 创建bug页面
	TaskManager.Post("/bug/create", taskcontroller.Create).Bind(&model.Bug{}).
		AddModule(midware.JsonToStruct)

	TaskManager.Get("/bug/edit", taskcontroller.Get)
}