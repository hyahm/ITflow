package routegroup

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Admin 管理员组
var Admin *xmux.GroupRoute

func init() {
	Admin = xmux.NewGroupRoute().ApiCreateGroup("amdin", "管理员", "admin")

	Admin.Get("/admin/reset", handle.Reset).ApiDescribe("修改管理员密码，只能本地使用(已完成)").ApiSupplement(`
	curl http://127.0.0.1:10001/admin/reset?password=123
	`).DelModule(midware.CheckToken)

	// Admin.Post("/info/update", handle.UpdateInfo).Bind(&user.UserInfo{}).
	// 	AddModule(midware.JsonToStruct).ApiDescribe("只做保留，没有用到")

}
