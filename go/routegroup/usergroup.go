package routegroup

import (
	"itflow/handle"
	"itflow/internal/usergroup"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var UserGroup *xmux.GroupRoute

func init() {
	UserGroup = xmux.NewGroupRoute().ApiCreateGroup("usergroup", "用户组相关", "user group")
	UserGroup.ApiCodeField("code").ApiCodeMsg("0", "成功")
	UserGroup.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	UserGroup.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	UserGroup.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	UserGroup.Pattern("/group/get").Post(handle.UserGroupGet).ApiResponseTemplate(`"grouplist":[],"code":0`).ApiDescribe("获取自己创建的用户组")

	UserGroup.Pattern("/groupnames/get").Post(handle.GroupNamesGet).ApiResponseTemplate(`"groupnames":[],"code":0`).ApiDescribe("获取自己创建的用户组")

	UserGroup.Pattern("/group/add").Post(handle.GroupAdd).Bind(&usergroup.RespUserGroup{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	UserGroup.Pattern("/group/del").Get(handle.GroupDel).End(midware.EndLog).ApiDescribe("删除用户组，只有创建者和admin才能操作")

	UserGroup.Pattern("/group/update").Post(handle.GroupUpdate).Bind(&usergroup.RespUpdateUserGroup{}).
		AddMidware(midware.JsonToStruct).AddMidware(midware.CheckUser).
		End(midware.EndLog).ApiDescribe("编辑用户组，只有创建者和admin才能操作")
}
