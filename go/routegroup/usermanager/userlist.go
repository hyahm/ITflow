package usermanager

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var UserListPage *xmux.GroupRoute

func init() {
	UserListPage = xmux.NewGroupRoute().AddModule(midware.JobAuth).AddPageKeys("admin", "user").AddModule(midware.CheckRole)
	// 获取所有用户
	UserListPage.Post("/user/list", handle.Read)
	// 删除用户
	UserListPage.Get("/user/remove", handle.Delete).ApiDescribe("删除用户")
	// 禁用用户
	UserListPage.Get("/user/disable", handle.DisableUser).ApiDescribe("禁用用户")
	// 用户信息修改
	UserListPage.Post("/user/update", handle.Update).Bind(model.User{}).
		AddModule(midware.JsonToStruct)
}
