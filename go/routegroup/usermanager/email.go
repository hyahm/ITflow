package usermanager

import (
	"itflow/handle/usercontroller/email"

	"github.com/hyahm/xmux"
)

var UpdateEmailPage *xmux.RouteGroup

func init() {
	UpdateEmailPage = xmux.NewRouteGroup()
	// 获取自己的邮箱
	UpdateEmailPage.Post("/my/email", email.Get)
	// 修改邮箱
	UpdateEmailPage.Post("/email/update", email.Set)
}
