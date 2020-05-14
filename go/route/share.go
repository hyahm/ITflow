package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Share *xmux.GroupRoute

func init() {
	Share = xmux.NewGroupRoute()
	Share.Pattern("/share/list").Get(handle.ShareList)
	Share.Pattern("/share/upload").Post(handle.ShareUpload)
	Share.Pattern("/share/mkdir").Post(handle.ShareMkdir).Bind(&model.Data_sharefile{}).AddMidware(midware.JsonToStruct)
	Share.Pattern("/share/remove").Get(handle.ShareRemove).End(midware.EndLog)
	Share.Pattern("/share/rename").Post(handle.ShareRename).Bind(&model.Data_sharefile{}).AddMidware(midware.JsonToStruct)
	//router.HandleFunc("/share/down", handle.ShareDownload)
	Share.Pattern("/share/down").Get(handle.ShareShow)
}
