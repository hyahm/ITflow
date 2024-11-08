package env

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Env 环境路由
var Env *xmux.RouteGroup

func init() {
	Env = xmux.NewRouteGroup().AddPageKeys("env")

	Env.Post("/env/list", Read)

	Env.Post("/env/add", Create).BindJson(&model.Env{})
	Env.Post("/env/update", Update).BindJson(&model.Env{})
	Env.Get("/env/delete", Delete)
}
