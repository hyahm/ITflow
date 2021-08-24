package system

import (
	"itflow/handle"
	"itflow/internal/rolegroup"
	"itflow/midware"
	"itflow/response"

	"github.com/hyahm/xmux"
)

// RoleGroup 角色组路由
var RoleGroup *xmux.GroupRoute

func init() {
	RoleGroup = xmux.NewGroupRoute().
		ApiCreateGroup("rolegroup", "角色组相关", "rolegroup").ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxx")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("0", "成功")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	RoleGroup.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	RoleGroup.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	RoleGroup.Post("/rolegroup/add", handle.AddRoleGroup).Bind(&rolegroup.ReqRoleGroup{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("添加角色组").ApiReqStruct(&rolegroup.ReqRoleGroup{}).ApiResStruct(&response.Response{})

	RoleGroup.Post("/rolegroup/edit", handle.EditRoleGroup).Bind(&rolegroup.ReqRoleGroup{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("修改角色组").ApiReqStruct(&rolegroup.ReqRoleGroup{}).ApiResStruct(&response.Response{})

	RoleGroup.Post("/rolegroup/list", handle.RoleGroupList)

	RoleGroup.Post("/rolegroup/get", handle.GetRoleGroupName)

	RoleGroup.Get("/rolegroup/remove", handle.RoleGroupDel)

	RoleGroup.Get("/roles/get", handle.GetRoles)
	// 获取编辑组的权限
	RoleGroup.Get("/rolegroup/perm/get", handle.GetRoleGroupPerm)

	RoleGroup.Post("/rolegroup/template", handle.RoleTemplate)

	// RoleGroup.Pattern("/rolegroup/name").Get(handle.GetRoleGroup)
}
