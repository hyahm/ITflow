package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)
var User *xmux.GroupRoute


func init() {
	User = xmux.NewGroupRoute()
	User.Pattern("/user/login").Post(handle.Login)
	User.Pattern("/user/logout").Post(handle.Loginout)
	User.Pattern("/user/info").Get(handle.UserInfo)
	User.Pattern("/user/list").Post(handle.UserList)
	User.Pattern("/user/update").Post(handle.UserUpdate)
	User.Pattern("/user/create").Post(handle.CreateUser)
	User.Pattern("/user/remove").Get(handle.RemoveUser)
	User.Pattern("/user/disable").Get(handle.DisableUser)
	User.Pattern("/password/update").Post(handle.ChangePassword)
	User.Pattern("/password/reset").Post(handle.ResetPwd)

}
