package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Share *xmux.GroupRoute

func init() {
	Share = xmux.NewGroupRoute("share")
	Share.Pattern("/share/list").Get(handle.ShareList)
	Share.Pattern("/share/upload").Post(handle.ShareUpload)
	Share.Pattern("/share/mkdir").Post(handle.ShareMkdir)
	Share.Pattern("/share/remove").Get(handle.ShareRemove).End(midware.EndLog)
	Share.Pattern("/share/rename").Post(handle.ShareRename)
	//router.HandleFunc("/share/down", handle.ShareDownload)
	Share.Pattern("/share/down").Get(handle.ShareShow)
}
