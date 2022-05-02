package routegroup

import (
	"itflow/handle"

	"github.com/hyahm/xmux"
)

// DashBoard 主页上的
var DashBoard *xmux.RouteGroup

func init() {
	DashBoard = xmux.NewRouteGroup()
	DashBoard.Post("/dashboard/bugcount", handle.BugCount)

	DashBoard.Post("/dashboard/projectcount", handle.ProjectCount)
	DashBoard.Post("/dashboard/usercount", handle.UserCount)
}
