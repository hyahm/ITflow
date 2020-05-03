package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Type *xmux.GroupRoute

func init() {
	Type = xmux.NewGroupRoute()
	Type.Pattern("/type/list").Post(handle.TypeList)
	Type.Pattern("/type/update").Post(handle.TypeUpdate).End(midware.EndLog)
	Type.Pattern("/type/add").Post(handle.TypeAdd).End(midware.EndLog)
	Type.Pattern("/type/delete").Get(handle.TypeDel).End(midware.EndLog)
	Type.Pattern("/type/get").Get(handle.GetType)

}
