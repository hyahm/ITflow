package routegroup

import (
	"itflow/handle"
	"itflow/internal/user"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Admin 管理员组
var Admin *xmux.GroupRoute

func init() {
	Admin = xmux.NewGroupRoute().ApiCreateGroup("amdin", "管理员", "admin")
	Admin.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Admin.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Admin.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Admin.ApiCodeField("code").ApiCodeMsg("其他错误", "请查看返回的msg")
	Admin.Post("/dashboard/usercount", handle.UserCount)

	Admin.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Admin.Get("/admin/reset", handle.Reset)

	Admin.Post("/info/update", handle.UpdateInfo).Bind(&user.UserInfo{}).
		AddMidware(midware.JsonToStruct)

}
