package status

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Status 状态路由组
var Status *xmux.RouteGroup

func init() {
	Status = xmux.NewRouteGroup().AddPageKeys("status")

	Status.Post("/status/list", Read)

	Status.Post("/status/add", Create).BindJson(&model.Status{})

	Status.Get("/status/remove", Delete)

	Status.Post("/status/update", Update).BindJson(&model.Status{})

}
