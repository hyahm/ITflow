package routegroup

import (
	"itflow/handle"

	"github.com/hyahm/xmux"
)

// DashBoard 主页上的
var DashBoard *xmux.GroupRoute

func init() {
	DashBoard = xmux.NewGroupRoute()
	DashBoard.Post("/dashboard/bugcount", handle.BugCount)

	DashBoard.Post("/dashboard/projectcount", handle.ProjectCount)
	DashBoard.Post("/dashboard/usercount", handle.UserCount)
}
