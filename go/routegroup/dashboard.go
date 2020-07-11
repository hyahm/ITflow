package routegroup

import (
	"github.com/hyahm/xmux"
	"itflow/handle"
)
var DashBoard *xmux.GroupRoute

func init() {
	DashBoard = xmux.NewGroupRoute().ApiCreateGroup("dashboard", "面板相关接口", "面板").
	ApiCodeField("code").ApiCodeMsg("0", "成功").
	ApiCodeField("code").ApiCodeMsg("20", "token过期").
	ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg").
	ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	DashBoard.Pattern("/dashboard/bugcount").Post(handle.BugCount).
	ApiDescribe("计算7天的创建和完成的bug数量")
}