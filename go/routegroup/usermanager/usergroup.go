package usermanager

import (
	"itflow/handle/usercontroller/usergroup"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var UserGroupPage *xmux.GroupRoute

func init() {
	UserGroupPage = xmux.NewGroupRoute()
	// 获取所有用户的信息关系
	UserGroupPage.Post("/alluser/keyname", usergroup.GetAllUserKeyName)
	// 获取所有用户组
	UserGroupPage.Post("/usergroup/list", usergroup.List)
	// 修改所有用户组
	UserGroupPage.Post("/usergroup/update", usergroup.Update).BindJson(&model.UserGroup{})
	// // 增加所有用户组
	UserGroupPage.Post("/usergroup/create", usergroup.Create).Bind(&model.UserGroup{}).AddModule(midware.JsonToStruct)
	// // 删除所有用户组
	UserGroupPage.Get("/usergroup/delete", usergroup.Delete).ApiDescribe("删除用户组，只有创建者和admin才能操作")

}
