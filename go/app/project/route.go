package project

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

//Project 项目路由
var Project *xmux.GroupRoute

func init() {
	Project = xmux.NewGroupRoute()
	Project.Post("/project/list", Read).ApiDescribe("获取所有列表").ApiSupplement("只有跟项目有关的参与者才能看到项目")

	Project.Post("/project/add", Create).Bind(&model.Project{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("增加项目")

	Project.Post("/project/update", Update).
		Bind(&model.Project{}).AddModule(midware.JsonToStruct).
		ApiDescribe("修改项目")

	Project.Get("/project/delete", Delete).
		ApiDescribe("删除项目")

	Project.Post("/get/project", ProjectKeys).ApiDescribe("获取所有项目名")
	// Project.Post("/get/myproject", handle.GetMyProject).ApiDescribe("获取自己权限的项目名")
}
