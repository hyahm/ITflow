package routegroup

import (
	"itflow/handle"
	"itflow/internal/status"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// StatusGroup 状态组路由
var StatusGroup *xmux.GroupRoute

func init() {
	StatusGroup = xmux.NewGroupRoute().ApiCreateGroup("statusgroup", "状态组操作", "bug status group")

	StatusGroup.ApiCodeField("code").ApiCodeMsg("0", "成功")
	StatusGroup.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	StatusGroup.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	StatusGroup.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	StatusGroup.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	StatusGroup.Post("/statusgroup/add", handle.AddStatusGroup).Bind(&status.StatusGroup{}).
		AddModule(midware.JsonToStruct)

	StatusGroup.Post("/statusgroup/edit", handle.EditStatusGroup).Bind(&status.StatusGroup{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("编辑statusgroup").ApiReqStruct(&status.StatusGroup{}).
		ApiSupplement("如果没有状态组名，就删除")

	StatusGroup.Post("/statusgroup/list", handle.StatusGroupList)

	StatusGroup.Get("/statusgroup/remove", handle.DeleteStatusGroup)
	StatusGroup.Post("/statusgroup/name", handle.GetStatusGroupName)
}
