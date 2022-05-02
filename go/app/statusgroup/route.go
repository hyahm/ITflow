package statusgroup

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// StatusGroup 状态组路由
var StatusGroup *xmux.RouteGroup

func init() {
	StatusGroup = xmux.NewRouteGroup().AddPageKeys("statusgroup")

	StatusGroup.Post("/statusgroup/add", Create).BindJson(&model.StatusGroup{})

	StatusGroup.Post("/statusgroup/edit", Update).Bind(&model.StatusGroup{}).
		AddModule(midware.JsonToStruct)
	StatusGroup.Post("/statusgroup/list", Read)

	StatusGroup.Get("/statusgroup/remove", Delete)
	StatusGroup.Post("/statusgroup/keyname", GetStatusGroupName)
}
