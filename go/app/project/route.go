package project

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

//Project 项目路由
var Project *xmux.RouteGroup

func init() {
	Project = xmux.NewRouteGroup().AddPageKeys("project")
	Project.Post("/project/list", Read)

	Project.Post("/project/add", Create).BindJson(&model.Project{})

	Project.Post("/project/update", Update).BindJson(&model.Project{})

	Project.Get("/project/delete", Delete)

	Project.Post("/get/project", ProjectKeys)
	// Project.Post("/get/myproject", handle.GetMyProject).ApiDescribe("获取自己权限的项目名")
}
