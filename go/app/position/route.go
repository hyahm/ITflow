package position

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Position *xmux.RouteGroup

func init() {
	Position = xmux.NewRouteGroup().AddPageKeys("position")

	Position.Post("/position/list", Read)
	Position.Post("/position/add", Create).Bind(&model.Job{}).AddModule(midware.JsonToStruct)

	Position.Get("/position/del", Delete)

	Position.Post("/position/update", Update).
		Bind(&model.Job{}).AddModule(midware.JsonToStruct)

	// Position.Post("/get/positions", handle.GetPositions).AddModule(midware.UserPerm)
}
