package routegroup

import (
	"itflow/handle"
	"itflow/internal/defaults"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// DefaultValue 默认值
var DefaultValue *xmux.GroupRoute

func init() {
	DefaultValue = xmux.NewGroupRoute().ApiCreateGroup("defaultvalue", "默认值相关", "默认值")
	DefaultValue.Post("/default/status", handle.DefaultStatus).ApiDescribe("获取默认值")

	DefaultValue.Post("/default/save", handle.DefaultSave).Bind(&defaults.ReqDefaultValue{}).
		AddMidware(midware.JsonToStruct).ApiDescribe("保存默认值")

}
