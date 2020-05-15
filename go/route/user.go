package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/response"
	"itflow/network/user"

	"github.com/hyahm/xmux"
)

var User *xmux.GroupRoute

func init() {
	User = xmux.NewGroupRoute()
	User.ApiCreateGroup("user", "用户相关的", "user")
	User.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	User.ApiCodeField("code").ApiCodeMsg("0", "成功")
	User.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	User.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	User.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")

	User.Pattern("/user/login").Post(handle.Login).Bind(&user.Login{}).
		DelMidware(midware.CheckToken).AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("用户登录接口").
		ApiDelReqHeader("X-Token").
		ApiReqStruct(user.Login{}).
		ApiResStruct(user.RespLogin{}).
		ApiRequestTemplate(`{"username":"admin", "password": "123456"}`).
		ApiResponseTemplate(`{"username":"admin","token":"sdfhdffffsdfgasdfasdf", "code": 0}`)
	User.Pattern("/user/logout").Post(handle.LoginOut).
		End(midware.EndLog).
		ApiDescribe("用户退出接口").
		ApiResStruct(response.Response{}).
		ApiSupplement("返回码是大部分公用的")

	User.Pattern("/user/info").Get(handle.UserInfo).
		ApiDescribe("获取用户信息").ApiCodeMsg("10", "token 过期").ApiCodeMsg("0", "成功").
		ApiResStruct(user.UserInfo{}).
		ApiCodeMsg("10", "token 过期").
		ApiResponseTemplate(`{"roles": ["admin"], "code": 0, "avatar":"http://xxxx/aaaa.png", "nickname": "admin"}`)

	User.Pattern("/user/list").Post(handle.UserList).Bind(&user.UserList{}).AddMidware(midware.JsonToStruct)

	User.Pattern("/user/update").Post(handle.UserUpdate).Bind(&user.User{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	User.Pattern("/user/create").Post(handle.CreateUser).Bind(&user.GetAddUser{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	User.Pattern("/user/remove").Get(handle.RemoveUser).End(midware.EndLog)

	User.Pattern("/user/disable").Get(handle.DisableUser).End(midware.EndLog)

	User.Pattern("/password/update").Post(handle.ChangePassword).Bind(&user.ChangePasswod{}).
		AddMidware(midware.JsonToStruct)

	User.Pattern("/password/reset").Post(handle.ResetPwd).Bind(&user.ResetPassword{}).AddMidware(midware.JsonToStruct)

}
