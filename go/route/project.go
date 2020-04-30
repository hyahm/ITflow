package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Project *xmux.GroupRoute

func init() {
	Project = xmux.NewGroupRoute("project")
	Project.Pattern("/project/list").Post(handle.ProjectList)
	Project.Pattern("/project/add").Post(handle.AddProject).End(midware.EndLog)
	Project.Pattern("/project/update").Post(handle.UpdateProject).End(midware.EndLog)
	Project.Pattern("/project/delete").Get(handle.DeleteProject).End(midware.EndLog)
}
