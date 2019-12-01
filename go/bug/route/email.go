package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Email *xmux.GroupRoute

func init() {
	Email = xmux.NewGroupRoute()
	Email.Pattern("/email/test").Post(handle.TestEmail)
	Email.Pattern("/email/save").Post(handle.SaveEmail)
	Email.Pattern("/email/get").Post(handle.GetEmail)
}
