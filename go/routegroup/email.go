package routegroup

import (
	"itflow/handle"
	"itflow/internal/email"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Email 邮件相关路由组
var Email *xmux.GroupRoute

func init() {
	Email = xmux.NewGroupRoute().ApiCreateGroup("email", "email 相关", "email")
	Email.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Email.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Email.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Email.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Email.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Email.Pattern("/email/test").Post(handle.TestEmail).Bind(&email.Email{}).AddMidware(midware.JsonToStruct)
	Email.Pattern("/email/save").Post(handle.SaveEmail).Bind(&email.Email{}).AddMidware(midware.JsonToStruct)
	Email.Pattern("/email/get").Post(handle.GetEmail)
}
