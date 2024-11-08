package routegroup

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Hook *xmux.RouteGroup

func init() {
	Hook = xmux.NewRouteGroup()
	Hook.Post("/gitlab/{name}", handle.Gitlab).DelModule(midware.CheckToken)
	Hook.Post("/gitee/{name}", handle.Gitee).DelModule(midware.CheckToken)
	Hook.Post("/github/{name}", handle.Github).DelModule(midware.CheckToken)
	Hook.Post("/random/hook", handle.RandomHook)
}
