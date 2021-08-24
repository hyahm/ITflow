package setting

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// StatusGroup 状态组路由
var StatusGroup *xmux.GroupRoute

func init() {
	StatusGroup = xmux.NewGroupRoute()

	StatusGroup.Post("/statusgroup/add", handle.AddStatusGroup).BindJson(&model.StatusGroup{})

	StatusGroup.Post("/statusgroup/edit", handle.EditStatusGroup).Bind(&model.StatusGroup{}).
		AddModule(midware.JsonToStruct)
	StatusGroup.Post("/statusgroup/list", handle.StatusGroupList)

	StatusGroup.Get("/statusgroup/remove", handle.DeleteStatusGroup)
	StatusGroup.Post("/statusgroup/keyname", handle.GetStatusGroupName)
}
