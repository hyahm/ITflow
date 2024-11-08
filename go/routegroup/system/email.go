package system

import (
	"itflow/handle"
	"itflow/internal/email"

	"github.com/hyahm/xmux"
)

// Email 邮件相关路由组
var Email *xmux.RouteGroup

func init() {
	Email = xmux.NewRouteGroup()

	Email.Post("/email/test", handle.TestEmail).BindJson(&email.Email{})
	Email.Post("/email/save", handle.SaveEmail).BindJson(&email.Email{})
	Email.Post("/email/get", handle.GetEmail)

}
