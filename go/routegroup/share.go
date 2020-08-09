package routegroup

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

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
	Share.Pattern("/share/list").Get(handle.ShareList)
	Share.Pattern("/share/upload").Post(handle.ShareUpload)
	Share.Pattern("/share/mkdir").Post(handle.ShareMkdir).Bind(&model.Data_sharefile{}).AddMidware(midware.JsonToStruct)
	Share.Pattern("/share/remove").Get(handle.ShareRemove).End(midware.EndLog)
	Share.Pattern("/share/rename").Post(handle.ShareRename).Bind(&model.Data_sharefile{}).AddMidware(midware.JsonToStruct)
	//router.HandleFunc("/share/down", handle.ShareDownload)
	Share.Pattern("/share/down").Get(handle.ShareShow)
}
