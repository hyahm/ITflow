package routegroup

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"itflow/internal/restful"

	"github.com/hyahm/xmux"
)

// Api 废弃
var Api *xmux.GroupRoute

func init() {
	Api = xmux.NewGroupRoute().ApiCreateGroup("api", "与api文档相关的所有的", "接口文档")

	Api.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Api.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Api.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Api.ApiCodeField("code").ApiCodeMsg("其他错误", "请查看返回的msg")

	Api.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Api.Post("/rest/list", handle.RestList)

	Api.Post("/rest/update", handle.RestUpdate).Bind(&model.Data_restful{}).AddMidware(midware.JsonToStruct)

	Api.Post("/rest/add", handle.RestAdd).Bind(&model.Data_restful{}).AddMidware(midware.JsonToStruct)

	Api.Get("/rest/delete", handle.RestDel)

	Api.Get("/api/list", handle.ApiList)

	Api.Post("/api/update", handle.ApiUpdate).Bind(&model.Get_apilist{}).AddMidware(midware.JsonToStruct)

	Api.Post("/api/add", handle.ApiAdd).Bind(&model.Get_apilist{}).AddMidware(midware.JsonToStruct)

	Api.Get("/api/delete", handle.ApiDel)

	Api.Get("/api/one", handle.ApiOne)

	Api.Get("/edit/one", handle.EditOne)

	Api.Post("/api/resp", handle.ApiResp).Bind(&restful.Resp{}).AddMidware(midware.JsonToStruct)

	Api.Post("/header/list", handle.HeaderList)

	Api.Post("/header/add", handle.HeaderAdd).Bind(&model.Data_header{}).AddMidware(midware.JsonToStruct)

	Api.Get("/header/del", handle.HeaderDel)

	Api.Post("/header/update", handle.HeaderUpdate).
		Bind(&model.Data_header{}).AddMidware(midware.JsonToStruct)

	Api.Post("/header/get", handle.HeaderGet)

	//---------------------------------------------------------

	// Api.Pattern("/default/important",handle.DefaultImportant)

	// Api.Pattern("/default/level",handle.DefaultLevel)
}
