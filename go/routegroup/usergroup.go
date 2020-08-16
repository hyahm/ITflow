package routegroup

import (
	"itflow/handle"
	"itflow/internal/usergroup"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// UserGroup 用户组
var UserGroup *xmux.GroupRoute

func init() {
	UserGroup = xmux.NewGroupRoute().ApiCreateGroup("usergroup", "用户组相关", "user group")
	UserGroup.ApiCodeField("code").ApiCodeMsg("0", "成功")
	UserGroup.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	UserGroup.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	UserGroup.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	UserGroup.Post("/group/get", handle.UserGroupGet).ApiResponseTemplate(`{"grouplist":["aa"],"code":0}`).ApiDescribe("获取自己创建的用户组")

	UserGroup.Post("/groupnames/get", handle.GroupNamesGet).ApiResponseTemplate(`{"groupnames":["bb"],"code":0}`).ApiDescribe("获取自己创建的用户组")

	UserGroup.Post("/group/add", handle.GroupAdd).Bind(&usergroup.RespUserGroup{}).AddMidware(midware.JsonToStruct)
	UserGroup.Get("/group/del", handle.GroupDel).ApiDescribe("删除用户组，只有创建者和admin才能操作")

	UserGroup.Post("/group/update", handle.GroupUpdate).Bind(&usergroup.RespUpdateUserGroup{}).
		AddMidware(midware.JsonToStruct).AddMidware(midware.CheckUser).
		ApiDescribe("编辑用户组，只有创建者和admin才能操作")
}
