package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Version *xmux.GroupRoute

func init() {
	Version = xmux.NewGroupRoute("version")
	Version.Pattern("/version/add").Post(handle.AddVersion)
	Version.Pattern("/version/list").Post(handle.VersionList)
	Version.Pattern("/version/remove").Get(handle.VersionRemove)
	Version.Pattern("/version/update").Post(handle.VersionUpdate)
}
