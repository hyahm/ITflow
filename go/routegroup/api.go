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
	Api.Pattern("/rest/list").Post(handle.RestList)

	Api.Pattern("/rest/update").Post(handle.RestUpdate).Bind(&model.Data_restful{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Api.Pattern("/rest/add").Post(handle.RestAdd).Bind(&model.Data_restful{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Api.Pattern("/rest/delete").Get(handle.RestDel).End(midware.EndLog)

	Api.Pattern("/api/list").Get(handle.ApiList)

	Api.Pattern("/api/update").Post(handle.ApiUpdate).Bind(&model.Get_apilist{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Api.Pattern("/api/add").Post(handle.ApiAdd).Bind(&model.Get_apilist{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Api.Pattern("/api/delete").Get(handle.ApiDel).End(midware.EndLog)

	Api.Pattern("/api/one").Get(handle.ApiOne)

	Api.Pattern("/edit/one").Get(handle.EditOne)

	Api.Pattern("/api/resp").Post(handle.ApiResp).Bind(&restful.Resp{}).AddMidware(midware.JsonToStruct)

	Api.Pattern("/header/list").Post(handle.HeaderList)

	Api.Pattern("/header/add").Post(handle.HeaderAdd).Bind(&model.Data_header{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Api.Pattern("/header/del").Get(handle.HeaderDel).End(midware.EndLog)

	Api.Pattern("/header/update").Post(handle.HeaderUpdate).Bind(&model.Data_header{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Api.Pattern("/header/get").Post(handle.HeaderGet)

	//---------------------------------------------------------

	// Api.Pattern("/default/important").Post(handle.DefaultImportant)

	// Api.Pattern("/default/level").Post(handle.DefaultLevel)
}
