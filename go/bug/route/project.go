package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Project *xmux.GroupRoute

func init() {
	Project = xmux.NewGroupRoute()
	Project.Pattern("/project/list").Post(handle.ProjectList)
	Project.Pattern("/project/add").Post(handle.AddProject)
	Project.Pattern("/project/update").Post(handle.UpdateProject)
	Project.Pattern("/project/delete").Get(handle.DeleteProject)
}
