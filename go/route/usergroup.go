package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var UserGroup *xmux.GroupRoute

func init() {
	UserGroup = xmux.NewGroupRoute().AddMidware(midware.CheckUserGroupPermssion)

	UserGroup.Pattern("/group/get").Post(handle.GroupGet)

	UserGroup.Pattern("/group/add").Post(handle.GroupAdd).Bind(&model.Get_groups{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	UserGroup.Pattern("/group/del").Get(handle.GroupDel).End(midware.EndLog)

	UserGroup.Pattern("/group/update").Post(handle.GroupUpdate).Bind(&model.Get_groups{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)
}
