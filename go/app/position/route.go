package position

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Position *xmux.RouteGroup

func init() {
	Position = xmux.NewRouteGroup().AddPageKeys("position")

	Position.Post("/position/list", Read)
	Position.Post("/position/add", Create).BindJson(&model.Job{})

	Position.Get("/position/del", Delete)

	Position.Post("/position/update", Update).BindJson(&model.Job{})

	// Position.Post("/get/positions", handle.GetPositions).AddModule(midware.UserPerm)
}
