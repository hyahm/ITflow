package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var RoleGroup *xmux.GroupRoute

func init() {
	RoleGroup = xmux.NewGroupRoute().AddMidware(midware.CheckRoleGroupPermssion)

	RoleGroup.Pattern("/rolegroup/add").Post(handle.AddRole).End(midware.EndLog)

	RoleGroup.Pattern("/rolegroup/edit").Post(handle.EditRole).
		End(midware.EndLog)

	RoleGroup.Pattern("/rolegroup/list").Post(handle.RoleList)

	RoleGroup.Pattern("/rolegroup/remove").Get(handle.RoleDel).End(midware.EndLog)

	RoleGroup.Pattern("/rolegroup/get").Get(handle.GetRoleGroup)

}
