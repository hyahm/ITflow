package routegroup

import (
	"itflow/handle"
	"itflow/internal/user"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// User 用户操作
var User *xmux.RouteGroup

func init() {
	User = xmux.NewRouteGroup()
	User.Post("/user/login", handle.Login).BindJson(&user.Login{}).
		DelModule(midware.CheckToken)
	User.Post("/user/logout", handle.LoginOut)
	User.Post("/get/user", handle.GetUser)

	User.Post("/is/admin", handle.IsAdmin)
}
