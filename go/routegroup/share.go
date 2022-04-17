package routegroup

import (
	"github.com/hyahm/xmux"
)

// Share 废弃
var Share *xmux.GroupRoute

func init() {
	Share = xmux.NewGroupRoute()
	// Share.Get("/share/list", handle.ShareList)
	// Share.Post("/share/upload", handle.ShareUpload)
	// Share.Post("/share/mkdir", handle.ShareMkdir).Bind(&model.Data_sharefile{}).AddModule(midware.JsonToStruct)
	// Share.Get("/share/remove", handle.ShareRemove)
	// Share.Post("/share/rename", handle.ShareRename).Bind(&model.Data_sharefile{}).AddModule(midware.JsonToStruct)
	// //router.HandleFunc("/share/down", handle.ShareDownload)
	// Share.Get("/share/down", handle.ShareShow)
}
