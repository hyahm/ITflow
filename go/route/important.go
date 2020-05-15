package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Important *xmux.GroupRoute

func init() {
	Important = xmux.NewGroupRoute().AddMidware(midware.CheckImportantPermssion)
	Important.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Important.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Important.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Important.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Important.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Important.Pattern("/important/get").Post(handle.ImportantGet)

	Important.Pattern("/important/add").Post(handle.ImportantAdd).Bind(&model.Data_importants{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Important.Pattern("/important/del").Get(handle.ImportantDel).End(midware.EndLog)

	Important.Pattern("/important/update").Post(handle.ImportantUpdate).Bind(&model.Update_importants{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Important.Pattern("/get/importants").Post(handle.GetImportants)

}
