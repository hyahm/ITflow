package routegroup

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Position 职位路由
var Position *xmux.GroupRoute

func init() {
	Position = xmux.NewGroupRoute()
	// AddModule(midware.CheckPositionPermssion)
	Position.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Position.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Position.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Position.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")

	Position.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Position.Post("/position/list", handle.PositionGet)

	Position.Post("/position/add", handle.PositionAdd).Bind(&model.Job{}).AddModule(midware.JsonToStruct)

	Position.Get("/position/del", handle.PositionDel)

	Position.Post("/position/update", handle.PositionUpdate).
		Bind(&model.Job{}).AddModule(midware.JsonToStruct)

	Position.Get("/get/hypos", handle.GetHypos)

	Position.Post("/get/positions", handle.GetPositions).AddModule(midware.UserPerm)
}
