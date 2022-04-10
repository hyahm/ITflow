package document

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Doc *xmux.GroupRoute

func init() {
	Doc = xmux.NewGroupRoute()
	Doc.Post("/doc/list", handle.DocList)
	Doc.Get("/doc/check/name", handle.Name)
	// Doc.Post("/doc/create", handle.DocCreate).Bind(&model.Doc{}).AddModule(midware.JsonToStruct)
	Doc.Get("/docs/{name}/{all:path}", handle.ProxyDoc).DelModule(midware.CheckToken)
	Doc.Get("/doc/update", handle.DocUpdate)
}
