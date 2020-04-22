package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Role *xmux.GroupRoute

func init() {
	Role = xmux.NewGroupRoute("role")
	Role.Pattern("/role/add").Post(handle.AddRole)
	Role.Pattern("/role/edit").Post(handle.EditRole)
	Role.Pattern("/role/list").Post(handle.RoleList)
	Role.Pattern("/role/remove").Get(handle.RoleDel)
	Role.Pattern("/role/get").Get(handle.GetRoles)
	Role.Pattern("/role/groupname").Post(handle.RoleGroupName)
}
