package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Api *xmux.GroupRoute

func init() {
	Api = xmux.NewGroupRoute("api")
	Api.Pattern("/rest/list").Post(handle.RestList)
	Api.Pattern("/rest/update").Post(handle.RestUpdate).End(midware.EndLog)
	Api.Pattern("/rest/add").Post(handle.RestAdd).End(midware.EndLog)
	Api.Pattern("/rest/delete").Get(handle.RestDel).End(midware.EndLog)

	Api.Pattern("/api/list").Get(handle.ApiList)
	Api.Pattern("/api/update").Post(handle.ApiUpdate).End(midware.EndLog)
	Api.Pattern("/api/add").Post(handle.ApiAdd).End(midware.EndLog)
	Api.Pattern("/api/delete").Get(handle.ApiDel).End(midware.EndLog)
	Api.Pattern("/api/one").Get(handle.ApiOne)
	Api.Pattern("/edit/one").Get(handle.EditOne)
	Api.Pattern("/api/resp").Post(handle.ApiResp)

	Api.Pattern("/group/get").Post(handle.GroupGet)
	Api.Pattern("/group/add").Post(handle.GroupAdd).End(midware.EndLog)
	Api.Pattern("/group/del").Get(handle.GroupDel).End(midware.EndLog)
	Api.Pattern("/group/update").Post(handle.GroupUpdate).End(midware.EndLog)

	Api.Pattern("/header/list").Post(handle.HeaderList)
	Api.Pattern("/header/add").Post(handle.HeaderAdd).End(midware.EndLog)
	Api.Pattern("/header/del").Get(handle.HeaderDel).End(midware.EndLog)
	Api.Pattern("/header/update").Post(handle.HeaderUpdate).End(midware.EndLog)
	Api.Pattern("/header/get").Post(handle.HeaderGet)

	Api.Pattern("/important/get").Post(handle.ImportantGet)
	Api.Pattern("/important/add").Post(handle.ImportantAdd).End(midware.EndLog)
	Api.Pattern("/important/del").Get(handle.ImportantDel).End(midware.EndLog)
	Api.Pattern("/important/update").Post(handle.ImportantUpdate).End(midware.EndLog)
	Api.Pattern("/get/importants").Post(handle.GetImportants)

	//---------------------------------------------------------
	Api.Pattern("/default/status").Post(handle.DefaultStatus)
	Api.Pattern("/default/save").Post(handle.DefaultSave)
	Api.Pattern("/default/important").Post(handle.DefaultImportant)
	Api.Pattern("/default/level").Post(handle.DefaultLevel)
}
