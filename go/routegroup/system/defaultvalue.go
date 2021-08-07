package system

import (
	"itflow/handle"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// DefaultValue 默认值
var DefaultValue *xmux.GroupRoute

func init() {
	DefaultValue = xmux.NewGroupRoute().ApiCreateGroup("defaultvalue", "默认值相关", "默认值")
	DefaultValue.Post("/default/status", handle.DefaultStatus).ApiDescribe("获取默认值(已完成)")

	DefaultValue.Post("/default/save", handle.DefaultSave).BindJson(&model.DefaultValue{})

}
