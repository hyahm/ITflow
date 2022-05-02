package env

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Env 环境路由
var Env *xmux.RouteGroup

func init() {
	Env = xmux.NewRouteGroup().AddPageKeys("env")

	Env.Post("/env/list", Read)

	Env.Post("/env/add", Create).Bind(&model.Env{}).
		AddModule(midware.JsonToStruct)

	Env.Post("/env/update", Update).Bind(&model.Env{}).
		AddModule(midware.JsonToStruct)
	Env.Get("/env/delete", Delete)
}
