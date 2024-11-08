package statusgroup

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

// StatusGroup 状态组路由
var StatusGroup *xmux.RouteGroup

func init() {
	StatusGroup = xmux.NewRouteGroup().AddPageKeys("statusgroup")

	StatusGroup.Post("/statusgroup/add", Create).BindJson(&model.StatusGroup{})

	StatusGroup.Post("/statusgroup/edit", Update).BindJson(&model.StatusGroup{})
	StatusGroup.Post("/statusgroup/list", Read)

	StatusGroup.Get("/statusgroup/remove", Delete)
	StatusGroup.Post("/statusgroup/keyname", GetStatusGroupName)
}
