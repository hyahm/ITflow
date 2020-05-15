package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Position *xmux.GroupRoute

func init() {
	Position = xmux.NewGroupRoute().AddMidware(midware.CheckPositionPermssion)
	Position.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Position.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Position.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Position.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")

	Position.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Position.Pattern("/position/list").Post(handle.PositionGet)

	Position.Pattern("/position/add").Post(handle.PositionAdd).Bind(&model.Data_jobs{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Position.Pattern("/position/del").Get(handle.PositionDel).End(midware.EndLog)

	Position.Pattern("/position/update").Post(handle.PositionUpdate).Bind(&model.Update_jobs{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Position.Pattern("/get/hypos").Post(handle.GetHypos)

	Position.Pattern("/get/positions").Post(handle.GetPositions)
}
