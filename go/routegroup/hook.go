package routegroup

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Hook *xmux.RouteGroup

func init() {
	Hook = xmux.NewRouteGroup()
	Hook.AddModule(midware.JsonToStruct)
	Hook.Post("/gitlab/{name}", handle.Gitlab).DelModule(midware.CheckToken).AddModule(midware.JsonToStruct)
	Hook.Post("/gitee/{name}", handle.Gitee).DelModule(midware.CheckToken).AddModule(midware.JsonToStruct)
	Hook.Post("/github/{name}", handle.Github).DelModule(midware.CheckToken).AddModule(midware.JsonToStruct)
	Hook.Post("/random/hook", handle.RandomHook)
}
