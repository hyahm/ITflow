package document

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Doc *xmux.RouteGroup

func init() {
	Doc = xmux.NewRouteGroup()
	Doc.Post("/doc/list", handle.DocList)
	Doc.Get("/doc/check/name", handle.Name)
	// Doc.Post("/doc/create", handle.DocCreate).BindJson(&model.Doc{})
	Doc.Get("/docs/{name}/{all:path}", handle.ProxyDoc).DelModule(midware.CheckToken)
	Doc.Get("/doc/update", handle.DocUpdate)
}
