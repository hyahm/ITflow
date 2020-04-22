package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Department *xmux.GroupRoute

func init() {
	Department = xmux.NewGroupRoute("department")
	Department.Pattern("/department/add").Post(handle.AddBugGroup)
	Department.Pattern("/department/edit").Post(handle.EditBugGroup)
	Department.Pattern("/department/list").Post(handle.BugGroupList)
	Department.Pattern("/department/remove").Get(handle.BugGroupDel)
}
