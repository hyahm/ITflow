package routegroup

import (
	"github.com/hyahm/xmux"
)

// Department 部门路由组
var Department *xmux.RouteGroup

func init() {
	Department = xmux.NewRouteGroup()

	// Department.Post("/department/add", handle.AddBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)

	// Department.Post("/department/edit", handle.EditBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)
	// // Department.Post("/department/list", handle.BugGroupList)
	// Department.Get("/department/remove", handle.BugGroupDel)
}
