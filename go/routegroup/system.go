package routegroup

import (
	"itflow/routegroup/system"

	"github.com/hyahm/xmux"
)

var SystemSetting *xmux.GroupRoute

func init() {
	SystemSetting = xmux.NewGroupRoute()
	// 默认状态页面
	SystemSetting.AddGroup(system.DefaultValue)
	// 邮箱设置页面
	SystemSetting.AddGroup(system.Email)
	// 日志页面
	SystemSetting.AddGroup(system.Log)
	// 角色页面
	SystemSetting.AddGroup(system.RoleGroup)
}
