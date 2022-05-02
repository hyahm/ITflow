package important

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Important *xmux.RouteGroup

func init() {
	Important = xmux.NewRouteGroup().AddPageKeys("important")

	Important.Post("/important/get", Read)

	Important.Post("/important/add", Create).
		Bind(&model.Important{}).AddModule(midware.JsonToStruct)

	Important.Get("/important/del", Delete)

	Important.Post("/important/update", Update).
		Bind(&model.Important{}).AddModule(midware.JsonToStruct)

	// Important.Post("/get/importants", GetImportants)

}
