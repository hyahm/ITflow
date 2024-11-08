package routegroup

import (
	"itflow/handle"
	"itflow/handle/publiccontroller"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// RoleGroup 角色组路由
var Public *xmux.RouteGroup

func init() {
	Public = xmux.NewRouteGroup()
	// 我的project
	Public.Post("/project/keyname", publiccontroller.GetProjectKeyName)
	Public.Post("/version/keyname/byproject", publiccontroller.GetVersionKeyNameByProject).BindJson(publiccontroller.RequestProject{})
	// 根据项目获取用户
	Public.Post("/user/keyname/byproject", publiccontroller.GetUserKeyNameByProject).BindJson(publiccontroller.RequestProject{})
	// 获取用户
	Public.Post("/user/keyname", publiccontroller.GetUserKeyName)
	// 获取运行环境
	Public.Post("/env/keyname", publiccontroller.GetEnvKeyName)
	// 获取有限级别
	Public.Post("/level/keyname", publiccontroller.GetLevelKeyName)
	// 获取重要性
	Public.Post("/important/keyname", publiccontroller.GetImportantKeyName)
	// 职位
	Public.Post("/position/keyname", publiccontroller.GetPositionKeyName)
	// 管理者的信息
	Public.Post("/manager/keyname", publiccontroller.GetManagerKeyName)

	// 用户组
	Public.Post("/usergroup/keyname", publiccontroller.GetUserGroupKeyName)
	// 获取用户信息， 所有页面都会用到
	Public.Get("/user/info", handle.UserInfo)
	// 获取自己能管理的bug状态
	Public.Post("/get/status", handle.GetStatus)
	// 获取自己先择显示的状态
	Public.Post("/status/show", publiccontroller.ShowStatus)
	// 修改显示的状态的bug
	Bug.Post("/status/save", handle.ChangeShowStatus).BindJson(&model.User{})

}
