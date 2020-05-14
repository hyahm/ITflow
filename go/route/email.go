package route

import (
	"itflow/app/bugconfig"
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Email *xmux.GroupRoute

func init() {
	Email = xmux.NewGroupRoute()
	Email.Pattern("/email/test").Post(handle.TestEmail).Bind(&bugconfig.Email{}).AddMidware(midware.JsonToStruct)
	Email.Pattern("/email/save").Post(handle.SaveEmail).Bind(&bugconfig.Email{}).AddMidware(midware.JsonToStruct)
	Email.Pattern("/email/get").Post(handle.GetEmail)
}
