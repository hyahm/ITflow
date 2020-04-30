package route

import (
	"itflow/bug/handle"
	"itflow/midware"
	"itflow/model/response"
	"itflow/model/user"

	"github.com/hyahm/xmux"
)

var User *xmux.GroupRoute

func init() {
	User = xmux.NewGroupRoute("user").ApiCreateGroup("user", "用户相关的", "user").ApiReqHeader("X-Token", "asdfasdfasdfasdfsdf")
	User.Pattern("/user/login").Post(handle.Login).
		DelMidware(midware.CheckToken).
		ApiDescribe("用户登录接口").
		ApiCodeMsg("10", "token 过期").
		ApiCodeMsg("0", "成功").
		ApiReqStruct(user.Login{}).
		Bind(&user.Login{}).
		AddMidware(midware.JsonToStruct).ApiResStruct(user.RespLogin{}).
		ApiRequestTemplate(`{"username":"admin", "password": "123456"}`).
		ApiResponseTemplate(`{"username":"admin","token":"sdfhdffffsdfgasdfasdf", "code": 0}`).
		End(midware.EndLog).ApiDelReqHeader("X-Token")

	User.Pattern("/user/logout").Post(handle.LoginOut).End(midware.EndLog).
		ApiDescribe("用户退出接口").ApiCodeMsg("10", "token 过期").ApiCodeMsg("0", "成功").
		ApiResStruct(response.Response{}).
		ApiSupplement("返回码是大部分公用的")

	User.Pattern("/user/info").Get(handle.UserInfo).
		ApiDescribe("获取用户信息").ApiCodeMsg("10", "token 过期").ApiCodeMsg("0", "成功").
		ApiResStruct(user.UserInfo{}).
		ApiCodeMsg("10", "token 过期").
		ApiResponseTemplate(`{"roles": ["admin"], "code": 0, "avatar":"http://xxxx/aaaa.png", "nickname": "admin"}`)

	User.Pattern("/user/list").Post(handle.UserList)
	User.Pattern("/user/update").Post(handle.UserUpdate).End(midware.EndLog)
	User.Pattern("/user/create").Post(handle.CreateUser).End(midware.EndLog)
	User.Pattern("/user/remove").Get(handle.RemoveUser).End(midware.EndLog)
	User.Pattern("/user/disable").Get(handle.DisableUser).End(midware.EndLog)
	User.Pattern("/password/update").Post(handle.ChangePassword)
	User.Pattern("/password/reset").Post(handle.ResetPwd)

}
