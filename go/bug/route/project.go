package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Project *xmux.GroupRoute

func init() {
	Project = xmux.NewGroupRoute("project")
	Project.Pattern("/project/list").Post(handle.ProjectList)
	Project.Pattern("/project/add").Post(handle.AddProject)
	Project.Pattern("/project/update").Post(handle.UpdateProject)
	Project.Pattern("/project/delete").Get(handle.DeleteProject)
}
