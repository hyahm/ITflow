package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Type *xmux.GroupRoute

func init() {
	Type = xmux.NewGroupRoute()
	Type.Pattern("/type/list").Post(handle.TypeList)
	Type.Pattern("/type/update").Post(handle.TypeUpdate).Bind(&model.Data_types{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Type.Pattern("/type/add").Post(handle.TypeAdd).Bind(&model.Data_types{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)
	Type.Pattern("/type/delete").Get(handle.TypeDel).End(midware.EndLog)
	Type.Pattern("/type/get").Get(handle.GetType)

}
