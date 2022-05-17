package routegroup

import (
	"itflow/handle"
	"itflow/handle/usercontroller"
	"itflow/internal/user"
	"itflow/midware"
	"itflow/model"
	"itflow/routegroup/usermanager"

	"github.com/hyahm/xmux"
)

// UserManager 用户管理
var UserManager *xmux.RouteGroup

func init() {
	UserManager = xmux.NewRouteGroup()

	// 用户组页面
	UserManager.AddGroup(usermanager.UserGroupPage)

	// 修改密码页面
	UserManager.Post("/password/reset", handle.ResetPwd).BindJson(&user.ResetPassword{})
	// 修改邮箱页面
	UserManager.AddGroup(usermanager.UpdateEmailPage)
	// 上传头像页面
	UserManager.Post("/upload/headimg", handle.UploadHeadImg)
	// 修改密码页面
	UserManager.Post("/password/update", usercontroller.ChangePassword).BindJson(&usercontroller.ChangePasswod{})
	// 用户列表管理页面
	UserManager.AddGroup(usermanager.UserListPage)
	// 用户创建
	// 添加用户操作
	UserManager.Post("/user/create", handle.Create).
		BindJson(&model.User{}).AddPageKeys("admin", "user").AddModule(midware.CheckRole)

}
