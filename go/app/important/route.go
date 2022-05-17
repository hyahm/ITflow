package important

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Important *xmux.RouteGroup

func init() {
	Important = xmux.NewRouteGroup().AddPageKeys("important")

	Important.Post("/important/get", Read)

	Important.Post("/important/add", Create).BindJson(&model.Important{})

	Important.Get("/important/del", Delete)

	Important.Post("/important/update", Update).BindJson(&model.Important{})

	// Important.Post("/get/importants", GetImportants)

}
