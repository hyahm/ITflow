package routegroup

import (
	"itflow/handle"

	"github.com/hyahm/xmux"
)

// DashBoard 主页上的
var DashBoard *xmux.GroupRoute

func init() {
	DashBoard = xmux.NewGroupRoute().ApiCreateGroup("dashboard", "面板相关接口", "面板").
		ApiCodeField("code").ApiCodeMsg("0", "成功").
		ApiCodeField("code").ApiCodeMsg("20", "token过期").
		ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg").
		ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	DashBoard.Post("/dashboard/bugcount", handle.BugCount).
		ApiDescribe("计算7天的创建和完成的bug数量(待测试)")

	DashBoard.Post("/dashboard/projectcount", handle.ProjectCount).ApiDescribe("总任务和总完成的任务数(待测试)")
	DashBoard.Post("/dashboard/usercount", handle.UserCount).ApiDescribe("用户总数（已完成）")
}
