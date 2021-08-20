package routegroup

import (
	"itflow/routegroup/document"

	"github.com/hyahm/xmux"
)

var Document *xmux.GroupRoute

func init() {
	Document = xmux.NewGroupRoute()
	Document.AddGroup(document.Key)
	Document.AddGroup(document.Doc)

	// Department.Post("/department/add", handle.AddBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)

	// Department.Post("/department/edit", handle.EditBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)
	// // Department.Post("/department/list", handle.BugGroupList)
	// Department.Get("/department/remove", handle.BugGroupDel)
}
