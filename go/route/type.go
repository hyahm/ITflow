package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Type *xmux.GroupRoute

func init() {
	Type = xmux.NewGroupRoute().ApiCreateGroup("type", "接口的类型", "type")

	Type.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Type.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Type.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Type.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Type.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Type.Pattern("/type/list").Post(handle.TypeList)
	Type.Pattern("/type/update").Post(handle.TypeUpdate).Bind(&model.Data_types{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Type.Pattern("/type/add").Post(handle.TypeAdd).Bind(&model.Data_types{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)
	Type.Pattern("/type/delete").Get(handle.TypeDel).End(midware.EndLog)
	Type.Pattern("/type/get").Get(handle.GetType)

}
