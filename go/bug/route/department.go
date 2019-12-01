package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Department *xmux.GroupRoute

func init() {
	Department = xmux.NewGroupRoute()
	Department.Pattern("/department/add").Post(handle.AddBugGroup)
	Department.Pattern("/department/edit").Post(handle.EditBugGroup)
	Department.Pattern("/department/list").Post(handle.BugGroupList)
	Department.Pattern("/department/remove").Get(handle.BugGroupDel)
}
