package usermanager

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var UserListPage *xmux.RouteGroup

func init() {
	UserListPage = xmux.NewRouteGroup().AddModule(midware.JobAuth).AddPageKeys("admin", "user").AddModule(midware.CheckRole)
	// 获取所有用户
	UserListPage.Post("/user/list", handle.Read)
	// 删除用户
	UserListPage.Get("/user/remove", handle.Delete)
	// 禁用用户
	UserListPage.Get("/user/disable", handle.DisableUser)
	// 用户信息修改
	UserListPage.Post("/user/update", handle.Update).BindJson(model.User{})
}
