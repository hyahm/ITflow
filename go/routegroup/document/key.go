package document

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Key *xmux.RouteGroup

func init() {
	Key = xmux.NewRouteGroup()
	Key.Post("/keys/list", handle.KeyList)
	Key.Post("/keys/add", handle.AddKey).Bind(&model.Auth{}).AddModule(midware.JsonToStruct)
	Key.Get("/keys/delete", handle.DeleteKey)
	Key.Get("/keys/check/name", handle.CheckKeyName)
	Key.Post("/keys/get/me", handle.GetMykeys)
}
