package status

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Status 状态路由组
var Status *xmux.GroupRoute

func init() {
	Status = xmux.NewGroupRoute().AddPageKeys("status")

	Status.Post("/status/list", Read)

	Status.Post("/status/add", Create).Bind(&model.Status{}).
		AddModule(midware.JsonToStruct)

	Status.Get("/status/remove", Delete)

	Status.Post("/status/update", Update).Bind(&model.Status{}).AddModule(midware.JsonToStruct)

}
