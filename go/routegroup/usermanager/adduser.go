package usermanager

import (
	"github.com/hyahm/xmux"
)

var AddUserPage *xmux.GroupRoute

func init() {
	AddUserPage = xmux.NewGroupRoute()

}
