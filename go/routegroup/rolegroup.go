package routegroup

import (
	"itflow/handle"
	"itflow/internal/response"
	"itflow/internal/rolegroup"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var RoleGroup *xmux.GroupRoute

func init() {
	RoleGroup = xmux.NewGroupRoute().
		// AddMidware(midware.CheckRoleGroupPermssion).
		ApiCreateGroup("rolegroup", "角色组相关", "rolegroup").ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxx")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("0", "成功")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	RoleGroup.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	RoleGroup.Pattern("/rolegroup/add").Post(handle.AddRoleGroup).Bind(&rolegroup.ReqRoleGroup{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("添加角色组").ApiReqStruct(&rolegroup.ReqRoleGroup{}).ApiResStruct(&response.Response{})

	RoleGroup.Pattern("/rolegroup/edit").Post(handle.EditRoleGroup).Bind(&rolegroup.ReqRoleGroup{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("修改角色组").ApiReqStruct(&rolegroup.ReqRoleGroup{}).ApiResStruct(&response.Response{})

	RoleGroup.Pattern("/rolegroup/list").Post(handle.RoleGroupList)
	RoleGroup.Pattern("/rolegroup/get").Post(handle.GetRoleGroupName)

	RoleGroup.Pattern("/rolegroup/remove").Get(handle.RoleGroupDel).End(midware.EndLog)

	RoleGroup.Pattern("/roles/get").Get(handle.GetRoles)

	RoleGroup.Pattern("/rolegroup/template").Post(handle.RoleTemplate)
	// RoleGroup.Pattern("/rolegroup/name").Get(handle.GetRoleGroup)
}
