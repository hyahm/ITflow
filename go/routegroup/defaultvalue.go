package routegroup

import (
	"itflow/handle"
	"itflow/internal/defaults"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var DefaultValue *xmux.GroupRoute

func init() {
	DefaultValue = xmux.NewGroupRoute().ApiCreateGroup("defaultvalue", "默认值相关", "默认值")
	DefaultValue.Pattern("/default/status").Post(handle.DefaultStatus).ApiDescribe("获取默认值")

	DefaultValue.Pattern("/default/save").Post(handle.DefaultSave).Bind(&defaults.ReqDefaultValue{}).
		AddMidware(midware.JsonToStruct).ApiDescribe("保存默认值")

}
