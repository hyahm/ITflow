package routegroup

import (
	"github.com/hyahm/xmux"
)

// Department 部门路由组
var Department *xmux.GroupRoute

func init() {
	Department = xmux.NewGroupRoute().ApiCreateGroup("department", "职位相关", "department")

	Department.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Department.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Department.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Department.ApiCodeField("code").ApiCodeMsg("其他错误", "请查看返回的msg")
	Department.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	// Department.Post("/department/add", handle.AddBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)

	// Department.Post("/department/edit", handle.EditBugGroup).
	// 	Bind(&status.StatusGroup{}).AddModule(midware.JsonToStruct)
	// // Department.Post("/department/list", handle.BugGroupList)
	// Department.Get("/department/remove", handle.BugGroupDel)
}
