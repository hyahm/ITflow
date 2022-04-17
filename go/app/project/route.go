package project

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

//Project 项目路由
var Project *xmux.GroupRoute

func init() {
	Project = xmux.NewGroupRoute().AddPageKeys("project")
	Project.Post("/project/list", Read)

	Project.Post("/project/add", Create).Bind(&model.Project{}).
		AddModule(midware.JsonToStruct)

	Project.Post("/project/update", Update).
		Bind(&model.Project{}).AddModule(midware.JsonToStruct)

	Project.Get("/project/delete", Delete)

	Project.Post("/get/project", ProjectKeys)
	// Project.Post("/get/myproject", handle.GetMyProject).ApiDescribe("获取自己权限的项目名")
}
