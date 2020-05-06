package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/user"

	"github.com/hyahm/xmux"
)

var Admin *xmux.GroupRoute

func init() {
	Admin = xmux.NewGroupRoute()
	Admin.Pattern("/dashboard/usercount").Post(handle.UserCount)
	Admin.Pattern("/dashboard/projectcount").Post(handle.ProjectCount)

	Admin.Pattern("/admin/reset").Get(handle.Reset)

	Admin.Pattern("/info/update").Post(handle.UpdateInfo).Bind(&user.UserInfo{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

}
