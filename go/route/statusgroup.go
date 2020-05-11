package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/status"

	"github.com/hyahm/xmux"
)

var StatusGroup *xmux.GroupRoute

func init() {
	StatusGroup = xmux.NewGroupRoute()
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
