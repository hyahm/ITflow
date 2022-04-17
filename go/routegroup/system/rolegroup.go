package system

import (
	"itflow/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// RoleGroup 角色组路由
var RoleGroup *xmux.GroupRoute

func init() {
	RoleGroup = xmux.NewGroupRoute()

	RoleGroup.Post("/rolegroup/add", handle.AddRoleGroup).Bind(&handle.RequestRoleGroup{}).
		AddModule(midware.JsonToStruct)
	RoleGroup.Post("/rolegroup/edit", handle.EditRoleGroup).Bind(&handle.RequestRoleGroup{}).
		AddModule(midware.JsonToStruct)
	RoleGroup.Post("/rolegroup/list", handle.RoleGroupList)

	RoleGroup.Post("/rolegroup/get", handle.GetRoleGroupName)

	RoleGroup.Get("/rolegroup/remove", handle.RoleGroupDel)

	RoleGroup.Get("/roles/get", handle.GetRoles)
	// 获取编辑组的权限
	RoleGroup.Get("/rolegroup/perm/get", handle.GetRoleGroupPerm)

	// RoleGroup.Post("/rolegroup/template", handle.RoleTemplate)

	// RoleGroup.Pattern("/rolegroup/name").Get(handle.GetRoleGroup)
}
