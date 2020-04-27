package route

import (
	"itflow/bug/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Version *xmux.GroupRoute

func init() {
	Version = xmux.NewGroupRoute("version")
	Version.Pattern("/version/add").Post(handle.AddVersion).End(midware.EndLog)
	Version.Pattern("/version/list").Post(handle.VersionList)
	Version.Pattern("/version/remove").Get(handle.VersionRemove).End(midware.EndLog)
	Version.Pattern("/version/update").Post(handle.VersionUpdate).End(midware.EndLog)
}
