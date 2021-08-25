package routegroup

import (
	"itflow/midware"
	"itflow/routegroup/setting"

	"github.com/hyahm/xmux"
)

var SettingCenter *xmux.GroupRoute

func init() {
	SettingCenter = xmux.NewGroupRoute().AddPageKeys("admin").AddModule(midware.CheckRole)

	SettingCenter.AddGroup(setting.Position).AddPageKeys("position")
	SettingCenter.AddGroup(setting.Env).AddPageKeys("env")
	SettingCenter.AddGroup(setting.Important).AddPageKeys("important")
	SettingCenter.AddGroup(setting.Level).AddPageKeys("level")
	SettingCenter.AddGroup(setting.Project).AddPageKeys("project")
	SettingCenter.AddGroup(setting.Status).AddPageKeys("status")
	SettingCenter.AddGroup(setting.StatusGroup).AddPageKeys("statusgroup")
	SettingCenter.AddGroup(setting.Version).AddPageKeys("version")
}
