package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Position *xmux.GroupRoute

func init() {
	Position = xmux.NewGroupRoute()
	Position.Pattern("/position/list").Post(handle.PositionGet)
	Position.Pattern("/position/add").Post(handle.PositionAdd)
	Position.Pattern("/position/del").Get(handle.PositionDel)
	Position.Pattern("/position/update").Post(handle.PositionUpdate)
	Position.Pattern("/get/hypos").Post(handle.GetHypos)
	Position.Pattern("/get/positions").Post(handle.GetPositions)
}
