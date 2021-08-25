package routegroup

import (
	"itflow/midware"
	"itflow/routegroup/system"

	"github.com/hyahm/xmux"
)

var SystemSetting *xmux.GroupRoute

func init() {
	SystemSetting = xmux.NewGroupRoute().AddPageKeys("admin").AddModule(midware.MustBeSuperAdmin)
	// 默认状态页面
	SystemSetting.AddGroup(system.DefaultValue)
	// 邮箱设置页面
	SystemSetting.AddGroup(system.Email)
	// 日志页面
	SystemSetting.AddGroup(system.Log).AddPageKeys("log").DelModule(midware.MustBeSuperAdmin).AddModule(midware.CheckRole)
	// 角色页面
	SystemSetting.AddGroup(system.RoleGroup)
}
