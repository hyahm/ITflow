package route

import (
	"itflow/bug/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Status *xmux.GroupRoute

func init() {
	Status = xmux.NewGroupRoute("status")
	Status.Pattern("/status/list").Post(handle.StatusList)
	Status.Pattern("/status/add").Post(handle.StatusAdd).End(midware.EndLog)
	Status.Pattern("/status/remove").Get(handle.StatusRemove).End(midware.EndLog)
	Status.Pattern("/status/update").Post(handle.StatusUpdate).End(midware.EndLog)
	Status.Pattern("/status/groupname").Post(handle.StatusGroupName)
}
