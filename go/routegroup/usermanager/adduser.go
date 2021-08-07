package usermanager

import (
	"itflow/handle"
	"itflow/internal/user"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var AddUserPage *xmux.GroupRoute

func init() {
	AddUserPage = xmux.NewGroupRoute()
	// 添加用户操作
	AddUserPage.Post("/user/create", handle.CreateUser).
		Bind(&user.GetAddUser{}).AddModule(midware.UserPerm).
		AddModule(midware.JsonToStruct)

}
