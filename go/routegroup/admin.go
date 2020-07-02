package routegroup

import (
	"itflow/handle"
	"itflow/internal/user"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Admin *xmux.GroupRoute

func init() {
	Admin = xmux.NewGroupRoute().ApiCreateGroup("amdin", "管理员", "admin")
	Admin.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Admin.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Admin.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Admin.ApiCodeField("code").ApiCodeMsg("其他错误", "请查看返回的msg")
	Admin.Pattern("/dashboard/usercount").Post(handle.UserCount)
	Admin.Pattern("/dashboard/projectcount").Post(handle.ProjectCount)
	Admin.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Admin.Pattern("/admin/reset").Get(handle.Reset)

	Admin.Pattern("/info/update").Post(handle.UpdateInfo).Bind(&user.UserInfo{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

}
