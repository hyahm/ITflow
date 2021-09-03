package level

import (
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Level *xmux.GroupRoute

func init() {
	Level = xmux.NewGroupRoute().AddPageKeys("level")
	Level.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Level.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Level.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Level.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Level.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Level.Post("/level/get", Read)

	Level.Post("/level/add", Create).
		Bind(&model.Level{}).AddModule(midware.JsonToStruct)

	Level.Get("/level/del", Delete)

	Level.Post("/level/update", Update).
		Bind(&model.Level{}).AddModule(midware.JsonToStruct)
}
