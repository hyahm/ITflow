package system

import (
	"itflow/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

// DefaultValue 默认值
var DefaultValue *xmux.GroupRoute

func init() {
	DefaultValue = xmux.NewGroupRoute().DelModule(midware.MustBeSuperAdmin)

	DefaultValue.Post("/default/status", handle.DefaultStatus)
	DefaultValue.Post("/default/save", handle.DefaultSave).BindJson(&model.DefaultValue{})

}
