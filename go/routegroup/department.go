package routegroup

import (
	"itflow/handle"
	"itflow/internal/status"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Department *xmux.GroupRoute

func init() {
	Department = xmux.NewGroupRoute().ApiCreateGroup("department", "职位相关", "department")

	Department.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Department.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Department.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Department.ApiCodeField("code").ApiCodeMsg("其他错误", "请查看返回的msg")
	Department.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Department.Pattern("/department/add").Post(handle.AddBugGroup).Bind(&status.StatusGroup{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Department.Pattern("/department/edit").Post(handle.EditBugGroup).Bind(&status.StatusGroup{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)
	Department.Pattern("/department/list").Post(handle.BugGroupList)
	Department.Pattern("/department/remove").Get(handle.BugGroupDel).End(midware.EndLog)
}
