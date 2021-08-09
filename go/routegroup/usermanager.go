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
var UserManager *xmux.GroupRoute

func init() {
	UserManager = xmux.NewGroupRoute().ApiCreateGroup("usermanager", "用户管理相关接口", "用户管理").
		ApiCodeField("code").ApiCodeMsg("0", "成功").
		ApiCodeField("code").ApiCodeMsg("20", "token过期").
		ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg").
		ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	// 用户组页面
	UserManager.AddGroup(usermanager.UserGroupPage)

	// 修改密码页面
	UserManager.Post("/password/reset", handle.ResetPwd).Bind(&user.ResetPassword{}).AddModule(midware.UserPerm).
		AddModule(midware.JsonToStruct).ApiDescribe("修改密码")
	// 修改邮箱页面
	UserManager.AddGroup(usermanager.UpdateEmailPage)
	// 上传头像页面
	UserManager.Post("/upload/headimg", handle.UploadHeadImg)
	// 修改密码页面
	UserManager.Post("/password/update", usercontroller.ChangePassword).Bind(&usercontroller.ChangePasswod{}).
		AddModule(midware.JsonToStruct)
	// 用户列表管理页面
	UserManager.AddGroup(usermanager.UserListPage)
	// 用户创建
	// 添加用户操作
	UserManager.Post("/user/create", handle.CreateUser).
		Bind(&model.User{}).AddModule(midware.UserPerm).
		AddModule(midware.JsonToStruct)

}
