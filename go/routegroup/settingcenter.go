package routegroup

import (
	"itflow/routegroup/setting"

	"github.com/hyahm/xmux"
)

var SettingCenter *xmux.GroupRoute

func init() {
	SettingCenter = xmux.NewGroupRoute()
	SettingCenter.AddGroup(setting.Position)
	SettingCenter.AddGroup(setting.Env)
	SettingCenter.AddGroup(setting.Important)
	SettingCenter.AddGroup(setting.Level)
	SettingCenter.AddGroup(setting.Project)
	SettingCenter.AddGroup(setting.Status)
	SettingCenter.AddGroup(setting.StatusGroup)
	SettingCenter.AddGroup(setting.Version)
}
