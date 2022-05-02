package routegroup

import (
	"itflow/routegroup/document"

	"github.com/hyahm/xmux"
)

var Document *xmux.RouteGroup

func init() {
	Document = xmux.NewRouteGroup()
	Document.AddGroup(document.Key)
	Document.AddGroup(document.Doc)

	// Department.Post("/department/add", handle.AddBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)

	// Department.Post("/department/edit", handle.EditBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)
	// // Department.Post("/department/list", handle.BugGroupList)
	// Department.Get("/department/remove", handle.BugGroupDel)
}
