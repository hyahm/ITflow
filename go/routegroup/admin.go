package routegroup

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Admin 管理员组
var Admin *xmux.RouteGroup

func init() {
	Admin = xmux.NewRouteGroup()

	Admin.Get("/admin/reset", handle.Reset).DelModule(midware.CheckToken)

	// Admin.Post("/info/update", handle.UpdateInfo).BindJson(&user.UserInfo{})

}
