package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Api *xmux.GroupRoute

func init() {
	Api = xmux.NewGroupRoute("api")
	Api.Pattern("/rest/list").Post(handle.RestList)
	Api.Pattern("/rest/update").Post(handle.RestUpdate)
	Api.Pattern("/rest/add").Post(handle.RestAdd)
	Api.Pattern("/rest/delete").Get(handle.RestDel)

	Api.Pattern("/api/list").Get(handle.ApiList)
	Api.Pattern("/api/update").Post(handle.ApiUpdate)
	Api.Pattern("/api/add").Post(handle.ApiAdd)
	Api.Pattern("/api/delete").Get(handle.ApiDel)
	Api.Pattern("/api/one").Get(handle.ApiOne)
	Api.Pattern("/edit/one").Get(handle.EditOne)
	Api.Pattern("/api/resp").Post(handle.ApiResp)

	Api.Pattern("/type/list").Post(handle.TypeList)
	Api.Pattern("/type/update").Post(handle.TypeUpdate)
	Api.Pattern("/type/add").Post(handle.TypeAdd)
	Api.Pattern("/type/delete").Get(handle.TypeDel)
	Api.Pattern("/type/get").Get(handle.GetType)

	Api.Pattern("/group/get").Post(handle.GroupGet)
	Api.Pattern("/group/add").Post(handle.GroupAdd)
	Api.Pattern("/group/del").Get(handle.GroupDel)
	Api.Pattern("/group/update").Post(handle.GroupUpdate)

	Api.Pattern("/header/list").Post(handle.HeaderList)
	Api.Pattern("/header/add").Post(handle.HeaderAdd)
	Api.Pattern("/header/del").Get(handle.HeaderDel)
	Api.Pattern("/header/update").Post(handle.HeaderUpdate)
	Api.Pattern("/header/get").Post(handle.HeaderGet)

	Api.Pattern("/important/get").Post(handle.ImportantGet)
	Api.Pattern("/important/add").Post(handle.ImportantAdd)
	Api.Pattern("/important/del").Get(handle.ImportantDel)
	Api.Pattern("/important/update").Post(handle.ImportantUpdate)
	Api.Pattern("/get/importants").Post(handle.GetImportants)

	Api.Pattern("/level/get").Post(handle.LevelGet)
	Api.Pattern("/level/add").Post(handle.LevelAdd)
	Api.Pattern("/level/del").Get(handle.LevelDel)
	Api.Pattern("/level/update").Post(handle.LevelUpdate)
	Api.Pattern("/get/levels").Post(handle.GetLevels)

	//---------------------------------------------------------
	Api.Pattern("/default/status").Post(handle.DefaultStatus)
	Api.Pattern("/default/save").Post(handle.DefaultSave)
	Api.Pattern("/default/important").Post(handle.DefaultImportant)
	Api.Pattern("/default/level").Post(handle.DefaultLevel)
}
