package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Admin *xmux.GroupRoute

func init() {
	Admin = xmux.NewGroupRoute("admin")
	Admin.Pattern("/dashboard/usercount").Post(handle.UserCount)
	Admin.Pattern("/dashboard/projectcount").Post(handle.ProjectCount)
	Admin.Pattern("/search/log").Post(handle.SearchLog)
	Admin.Pattern("/log/classify").Post(handle.LogClassify)
	Admin.Pattern("/admin/reset").Get(handle.Reset)
	Admin.Pattern("/info/update").Post(handle.UpdateInfo)
	//router.HandleFunc("/role/update", handle.UpdateRoles)
	Admin.Pattern("/log/list").Post(handle.LogList)
}
