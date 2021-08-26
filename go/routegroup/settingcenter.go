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

	SettingCenter.AddGroup(position.Position).AddPageKeys("position")
	SettingCenter.AddGroup(env.Env).AddPageKeys("env")
	SettingCenter.AddGroup(important.Important).AddPageKeys("important")
	SettingCenter.AddGroup(level.Level).AddPageKeys("level")
	SettingCenter.AddGroup(project.Project).AddPageKeys("project")
	SettingCenter.AddGroup(status.Status).AddPageKeys("status")
	SettingCenter.AddGroup(statusgroup.StatusGroup).AddPageKeys("statusgroup")
	SettingCenter.AddGroup(version.Version).AddPageKeys("version")
}
