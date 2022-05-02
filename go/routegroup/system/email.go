package system

import (
	"itflow/handle"
	"itflow/internal/email"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Email 邮件相关路由组
var Email *xmux.RouteGroup

func init() {
	Email = xmux.NewRouteGroup()

	Email.Post("/email/test", handle.TestEmail).Bind(&email.Email{}).AddModule(midware.JsonToStruct)
	Email.Post("/email/save", handle.SaveEmail).Bind(&email.Email{}).AddModule(midware.JsonToStruct)
	Email.Post("/email/get", handle.GetEmail)

}
