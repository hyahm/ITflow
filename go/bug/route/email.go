package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Email *xmux.GroupRoute

func init() {
	Email = xmux.NewGroupRoute("email")
	Email.Pattern("/email/test").Post(handle.TestEmail)
	Email.Pattern("/email/save").Post(handle.SaveEmail)
	Email.Pattern("/email/get").Post(handle.GetEmail)
}
