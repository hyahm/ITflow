package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Role *xmux.GroupRoute

func init() {
	Role = xmux.NewGroupRoute()
	Role.Pattern("/role/add").Post(handle.AddRole)
	Role.Pattern("/role/edit").Post(handle.EditRole)
	Role.Pattern("/role/list").Post(handle.RoleList)
	Role.Pattern("/role/remove").Get(handle.RoleDel)
	Role.Pattern("/role/get").Get(handle.GetRoles)
	Role.Pattern("/role/groupname").Post(handle.RoleGroupName)
}
