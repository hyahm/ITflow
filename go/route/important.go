package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Important *xmux.GroupRoute

func init() {
	Important = xmux.NewGroupRoute().AddMidware(midware.CheckImportantPermssion)

	Important.Pattern("/important/get").Post(handle.ImportantGet)

	Important.Pattern("/important/add").Post(handle.ImportantAdd).End(midware.EndLog)

	Important.Pattern("/important/del").Get(handle.ImportantDel).End(midware.EndLog)

	Important.Pattern("/important/update").Post(handle.ImportantUpdate).End(midware.EndLog)

	Important.Pattern("/get/importants").Post(handle.GetImportants)

}
