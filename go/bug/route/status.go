package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Status *xmux.GroupRoute

func init() {
	Status = xmux.NewGroupRoute("status")
	Status.Pattern("/status/list").Post(handle.StatusList)
	Status.Pattern("/status/add").Post(handle.StatusAdd)
	Status.Pattern("/status/remove").Get(handle.StatusRemove)
	Status.Pattern("/status/update").Post(handle.StatusUpdate)
	Status.Pattern("/status/groupname").Post(handle.StatusGroupName)
}
