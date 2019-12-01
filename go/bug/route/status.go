package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Status *xmux.GroupRoute


func init() {
	Status = xmux.NewGroupRoute()
	Status.Pattern("/status/list").Post(handle.StatusList)
	Status.Pattern("/status/add").Post(handle.StatusAdd)
	Status.Pattern("/status/remove").Get(handle.StatusRemove)
	Status.Pattern("/status/update").Post(handle.StatusUpdate)
	Status.Pattern("/status/groupname").Post(handle.StatusGroupName)
}
