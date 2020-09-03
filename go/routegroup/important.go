package routegroup

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Important 重要性路由
var Important *xmux.GroupRoute

func init() {
	Important = xmux.NewGroupRoute()
	Important.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Important.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Important.ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg")
	Important.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Important.Post("/important/get", handle.ImportantGet)

	Important.Post("/important/add", handle.ImportantAdd).
		Bind(&model.Data_importants{}).AddModule(midware.JsonToStruct)

	Important.Get("/important/del", handle.ImportantDel)

	Important.Post("/important/update", handle.ImportantUpdate).
		Bind(&model.Importants{}).AddModule(midware.JsonToStruct)

	Important.Post("/get/importants", handle.GetImportants)

}
