package usermanager

import (
	"github.com/hyahm/xmux"
)

var AddUserPage *xmux.RouteGroup

func init() {
	AddUserPage = xmux.NewRouteGroup()

}
