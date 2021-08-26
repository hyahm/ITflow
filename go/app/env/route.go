package env

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Env 环境路由
var Env *xmux.GroupRoute

func init() {
	Env = xmux.NewGroupRoute()

	Env.Post("/env/list", Read)

	Env.Get("/env/add", Create)

	Env.Post("/env/update", Update).Bind(&model.Env{}).
		AddModule(midware.JsonToStruct)
	Env.Get("/env/delete", Delete)
}
