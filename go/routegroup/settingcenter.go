package routegroup

import (
	"itflow/app/env"
	"itflow/app/important"
	"itflow/app/level"
	"itflow/app/position"
	"itflow/app/project"
	"itflow/app/status"
	"itflow/app/statusgroup"
	"itflow/app/version"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var SettingCenter *xmux.GroupRoute

func init() {
	SettingCenter = xmux.NewGroupRoute().AddPageKeys("admin").AddModule(midware.CheckRole)

	SettingCenter.AddGroup(position.Position)
	SettingCenter.AddGroup(env.Env)
	SettingCenter.AddGroup(important.Important)
	SettingCenter.AddGroup(level.Level)
	SettingCenter.AddGroup(project.Project)
	SettingCenter.AddGroup(status.Status)
	SettingCenter.AddGroup(statusgroup.StatusGroup)
	SettingCenter.AddGroup(version.Version)
}
