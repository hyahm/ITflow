package route

import (
	"itflow/bug/handle"
	"itflow/midware"
	"itflow/model/user"

	"github.com/hyahm/xmux"
)

var User *xmux.GroupRoute

func init() {
	User = xmux.NewGroupRoute("user")
	User.Pattern("/user/login").Post(handle.Login).
		ApiDescribe("用户登录接口").
		ApiReqStruct(user.Login{}).
		Bind(&user.Login{}).AddMidware(midware.JsonToStruct).ApiResStruct(user.RespLogin{}).
		ApiRequestTemplate(`{"username":"admin", "password": "123456"}`).
		ApiResponseTemplate(`{"username":"admin","token":"sdfhdffffsdfgasdfasdf", "code": 0}`)
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
