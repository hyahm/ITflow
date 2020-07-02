package routegroup

import (
	"itflow/handle"
	"itflow/internal/status"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var StatusGroup *xmux.GroupRoute

func init() {
	StatusGroup = xmux.NewGroupRoute().ApiCreateGroup("statusgroup", "状态组操作", "bug status group")

	StatusGroup.ApiCodeField("code").ApiCodeMsg("0", "成功")
	StatusGroup.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	StatusGroup.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	StatusGroup.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	StatusGroup.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	StatusGroup.Pattern("/statusgroup/add").Post(handle.AddStatusGroup).Bind(&status.StatusGroup{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	StatusGroup.Pattern("/statusgroup/edit").Post(handle.EditStatusGroup).Bind(&status.StatusGroup{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("编辑statusgroup").ApiReqStruct(&status.StatusGroup{}).
		ApiSupplement("如果没有状态组名，就删除")

	StatusGroup.Pattern("/statusgroup/list").Post(handle.StatusGroupList)

	StatusGroup.Pattern("/statusgroup/remove").Get(handle.DeleteStatusGroup).End(midware.EndLog)
	StatusGroup.Pattern("/statusgroup/name").Post(handle.GetStatusGroupName)
}
