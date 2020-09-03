package routegroup

import (
	"github.com/hyahm/xmux"
)

// Type 废弃
var Type *xmux.GroupRoute

func init() {
	Type = xmux.NewGroupRoute().ApiCreateGroup("type", "接口的类型", "type")

	Type.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Type.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Type.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Type.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Type.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	// Type.Post("/type/list", handle.TypeList)
	// Type.Post("/type/update", handle.TypeUpdate).Bind(&model.Data_types{}).AddModule(midware.JsonToStruct)

	// Type.Post("/type/add", handle.TypeAdd).Bind(&model.Data_types{}).AddModule(midware.JsonToStruct)
	// Type.Get("/type/delete", handle.TypeDel)
	// Type.Get("/type/get", handle.GetType)

}
