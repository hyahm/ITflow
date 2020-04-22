package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Position *xmux.GroupRoute

func init() {
	Position = xmux.NewGroupRoute("posistion")
	Position.Pattern("/position/list").Post(handle.PositionGet)
	Position.Pattern("/position/add").Post(handle.PositionAdd)
	Position.Pattern("/position/del").Get(handle.PositionDel)
	Position.Pattern("/position/update").Post(handle.PositionUpdate)
	Position.Pattern("/get/hypos").Post(handle.GetHypos)
	Position.Pattern("/get/positions").Post(handle.GetPositions)
}
