package level

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Level *xmux.RouteGroup

func init() {
	Level = xmux.NewRouteGroup().AddPageKeys("level")

	Level.Post("/level/get", Read)

	Level.Post("/level/add", Create).
		Bind(&model.Level{}).AddModule(midware.JsonToStruct)

	Level.Get("/level/del", Delete)

	Level.Post("/level/update", Update).
		Bind(&model.Level{}).AddModule(midware.JsonToStruct)
}
