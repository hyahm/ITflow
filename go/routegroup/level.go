package routegroup

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Level bug级别路由
var Level *xmux.GroupRoute

func init() {
	Level = xmux.NewGroupRoute().AddModule(midware.LevelPermModule)
	Level.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Level.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Level.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Level.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Level.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Level.Post("/level/get", handle.LevelGet)

	Level.Post("/level/add", handle.LevelAdd).
		Bind(&model.Data_level{}).AddModule(midware.JsonToStruct)

	Level.Get("/level/del", handle.LevelDel)

	Level.Post("/level/update", handle.LevelUpdate).
		Bind(&model.Update_level{}).AddModule(midware.JsonToStruct)
	Level.Post("/get/levels", handle.GetLevels)
}
