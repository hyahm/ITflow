package routegroup

import (
	"itflow/handle"
	"itflow/internal/project"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

//Project 项目路由
var Project *xmux.GroupRoute

func init() {
	Project = xmux.NewGroupRoute().ApiCreateGroup("project", "项目相关接口（建议给项目发起者添加操作）", "项目管理")
	Project.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Project.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Project.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Project.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Project.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Project.Post("/project/list", handle.ProjectList).ApiDescribe("获取所有列表").ApiSupplement("只有跟项目有关的参与者才能看到项目")

	Project.Post("/project/add", handle.AddProject).Bind(&project.ReqProject{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("增加项目")

	Project.Post("/project/update", handle.UpdateProject).
		Bind(&project.ReqProject{}).AddModule(midware.JsonToStruct).
		ApiDescribe("修改项目")

	Project.Get("/project/delete", handle.DeleteProject).
		ApiDescribe("删除项目")

	Project.Post("/get/project", handle.GetProject).ApiDescribe("获取所有项目名")
	Project.Post("/get/myproject", handle.GetMyProject).ApiDescribe("获取自己权限的项目名")
}
