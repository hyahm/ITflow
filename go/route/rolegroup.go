package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/rolegroup"

	"github.com/hyahm/xmux"
)

var RoleGroup *xmux.GroupRoute

func init() {
	RoleGroup = xmux.NewGroupRoute().AddMidware(midware.CheckRoleGroupPermssion)

	RoleGroup.Pattern("/rolegroup/add").Post(handle.AddRoleGroup).End(midware.EndLog)

	RoleGroup.Pattern("/rolegroup/edit").Post(handle.EditRoleGroup).Bind(&rolegroup.Data_roles{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	RoleGroup.Pattern("/rolegroup/list").Post(handle.RoleGroupList)

	RoleGroup.Pattern("/rolegroup/remove").Get(handle.RoleGroupDel).End(midware.EndLog)

	RoleGroup.Pattern("/roles/get").Get(handle.GetRoles)
	RoleGroup.Pattern("/rolegroup/name").Get(handle.GetRoleGroup)
}
