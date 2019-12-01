package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Version *xmux.GroupRoute

func init() {
	Version = xmux.NewGroupRoute()
	Version.Pattern("/version/add").Post(handle.AddVersion)
	Version.Pattern("/version/list").Post(handle.VersionList)
	Version.Pattern("/version/remove").Get(handle.VersionRemove)
	Version.Pattern("/version/update").Post(handle.VersionUpdate)
}
