package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/response"
	"itflow/network/rolegroup"

	"github.com/hyahm/xmux"
)

var RoleGroup *xmux.GroupRoute

func init() {
	RoleGroup = xmux.NewGroupRoute().AddMidware(midware.CheckRoleGroupPermssion).
		ApiCreateGroup("rolegroup", "角色组相关", "rolegroup").ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxx")

	RoleGroup.Pattern("/rolegroup/add").Post(handle.AddRoleGroup).Bind(&rolegroup.RoleGroup{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("添加角色组").ApiReqStruct(&rolegroup.RoleGroup{}).ApiResStruct(&response.Response{})

	RoleGroup.Pattern("/rolegroup/edit").Post(handle.EditRoleGroup).Bind(&rolegroup.RoleGroup{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("修改角色组").ApiReqStruct(&rolegroup.RoleGroup{}).ApiResStruct(&response.Response{})

	RoleGroup.Pattern("/rolegroup/list").Post(handle.RoleGroupList)

	RoleGroup.Pattern("/rolegroup/remove").Get(handle.RoleGroupDel).End(midware.EndLog)

	RoleGroup.Pattern("/roles/get").Get(handle.GetRoles)
	RoleGroup.Pattern("/rolegroup/name").Get(handle.GetRoleGroup)
}
