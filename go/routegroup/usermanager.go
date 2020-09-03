package routegroup

import (
	"itflow/handle"
	"itflow/internal/user"
	"itflow/midware"

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

	UserManager.Post("/password/reset", handle.ResetPwd).Bind(&user.ResetPassword{}).
		AddModule(midware.JsonToStruct).ApiDescribe("修改密码")

	UserManager.Get("/user/remove", handle.RemoveUser).ApiDescribe("删除用户")

	UserManager.Get("/user/disable", handle.DisableUser).ApiDescribe("禁用用户")
	UserManager.Post("/user/list", handle.UserList).ApiDescribe("获取用户列表")
	UserManager.Post("/user/update", handle.UserUpdate).Bind(&user.User{}).AddModule(midware.JsonToStruct).
		ApiDescribe("修改用户").ApiRequestTemplate(`{"id":3,"createtime":1594094022,"realname":"test","nickname":"test","email":"test@qq.com","disable":0,"statusgroup":"验证","rolegroup":"test","position":"python"}`)
}
