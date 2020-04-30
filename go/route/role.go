package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Role *xmux.GroupRoute

func init() {
	Role = xmux.NewGroupRoute("role")
	Role.Pattern("/role/add").Post(handle.AddRole).End(midware.EndLog)
	Role.Pattern("/role/edit").Post(handle.EditRole).End(midware.EndLog)
	Role.Pattern("/role/list").Post(handle.RoleList)
	Role.Pattern("/role/remove").Get(handle.RoleDel).End(midware.EndLog)
	Role.Pattern("/role/get").Get(handle.GetRoles)
	Role.Pattern("/role/groupname").Post(handle.RoleGroupName)
}
