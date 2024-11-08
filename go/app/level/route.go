package level

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Level *xmux.RouteGroup

func init() {
	Level = xmux.NewRouteGroup().AddPageKeys("level")

	Level.Post("/level/get", Read)

	Level.Post("/level/add", Create).BindJson(&model.Level{})

	Level.Get("/level/del", Delete)

	Level.Post("/level/update", Update).BindJson(&model.Level{})
}
