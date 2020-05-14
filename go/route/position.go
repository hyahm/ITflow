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

	Position.Pattern("/position/list").Post(handle.PositionGet)

	Position.Pattern("/position/add").Post(handle.PositionAdd).Bind(&model.Data_jobs{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Position.Pattern("/position/del").Get(handle.PositionDel).End(midware.EndLog)

	Position.Pattern("/position/update").Post(handle.PositionUpdate).Bind(&model.Update_jobs{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Position.Pattern("/get/hypos").Post(handle.GetHypos)

	Position.Pattern("/get/positions").Post(handle.GetPositions)
}
