package routegroup

import (
	"github.com/hyahm/xmux"
)

// Share 废弃
var Share *xmux.GroupRoute

func init() {
	Share = xmux.NewGroupRoute().ApiCreateGroup("share", "文件共享", "share file")
	Share.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Share.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Share.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Share.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Share.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	// Share.Get("/share/list", handle.ShareList)
	// Share.Post("/share/upload", handle.ShareUpload)
	// Share.Post("/share/mkdir", handle.ShareMkdir).Bind(&model.Data_sharefile{}).AddModule(midware.JsonToStruct)
	// Share.Get("/share/remove", handle.ShareRemove)
	// Share.Post("/share/rename", handle.ShareRename).Bind(&model.Data_sharefile{}).AddModule(midware.JsonToStruct)
	// //router.HandleFunc("/share/down", handle.ShareDownload)
	// Share.Get("/share/down", handle.ShareShow)
}
