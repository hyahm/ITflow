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

	Important.Pattern("/important/get").Post(handle.ImportantGet)

	Important.Pattern("/important/add").Post(handle.ImportantAdd).Bind(&model.Data_importants{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Important.Pattern("/important/del").Get(handle.ImportantDel).End(midware.EndLog)

	Important.Pattern("/important/update").Post(handle.ImportantUpdate).Bind(&model.Update_importants{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Important.Pattern("/get/importants").Post(handle.GetImportants)

}
