package routegroup

import (
	"github.com/hyahm/xmux"
)

// Share 废弃
var Share *xmux.RouteGroup

func init() {
	Share = xmux.NewRouteGroup()
	// Share.Get("/share/list", handle.ShareList)
	// Share.Post("/share/upload", handle.ShareUpload)
	// Share.Post("/share/mkdir", handle.ShareMkdir).BindJson(&model.Data_sharefile{})
	// Share.Get("/share/remove", handle.ShareRemove)
	// Share.Post("/share/rename", handle.ShareRename).BindJson(&model.Data_sharefile{})
	// //router.HandleFunc("/share/down", handle.ShareDownload)
	// Share.Get("/share/down", handle.ShareShow)
}
