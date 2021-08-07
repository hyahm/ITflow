package routegroup

import (
	"itflow/handle"
	"itflow/handle/publiccontroller"
	"itflow/internal/user"

	"github.com/hyahm/xmux"
)

// RoleGroup 角色组路由
var Public *xmux.GroupRoute

func init() {
	Public = xmux.NewGroupRoute()
	// 根据项目获取版本
	Public.Post("/version/keyname", publiccontroller.GetVersionKeyName).BindJson(publiccontroller.RequestProject{})
	// 根据项目获取用户
	Public.Post("/user/keyname", publiccontroller.GetUserKeyName).BindJson(publiccontroller.RequestProject{})
	// 获取运行环境
	Public.Post("/env/keyname", publiccontroller.GetEnvKeyName)
	// 获取有限级别
	Public.Post("/level/keyname", publiccontroller.GetLevelKeyName)
	// 获取重要性
	Public.Post("/important/keyname", publiccontroller.GetImportantKeyName)
	// 职位
	Public.Post("/position/keyname", publiccontroller.GetPositionKeyName)
	// 获取用户信息， 所有页面都会用到
	Public.Get("/user/info", handle.UserInfo).
		ApiDescribe("获取用户信息").
		ApiResStruct(user.UserInfo{})
	// 获取自己能管理的bug状态
	Public.Post("/get/status", handle.GetStatus)
	// 获取自己先择显示的状态
	Public.Post("/status/show", publiccontroller.ShowStatus).
		ApiDescribe("查询可以查看的状态")
}
