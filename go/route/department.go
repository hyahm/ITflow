package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Department *xmux.GroupRoute

func init() {
	Department = xmux.NewGroupRoute()
	Department.Pattern("/department/add").Post(handle.AddBugGroup).End(midware.EndLog)
	Department.Pattern("/department/edit").Post(handle.EditBugGroup).End(midware.EndLog)
	Department.Pattern("/department/list").Post(handle.BugGroupList)
	Department.Pattern("/department/remove").Get(handle.BugGroupDel).End(midware.EndLog)
}
