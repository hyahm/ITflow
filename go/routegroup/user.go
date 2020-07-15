package routegroup

import (
	"itflow/handle"
	"itflow/internal/response"
	"itflow/internal/user"
	"itflow/midware"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func InitUser() *xmux.GroupRoute {
	golog.Info(3333333)
	User := xmux.NewGroupRoute()
	User.ApiCreateGroup("user", "用户相关的", "user")
	User.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	User.ApiCodeField("code").ApiCodeMsg("0", "成功")
	User.ApiCodeField("code").ApiCodeMsg("20", "token过期")
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
		ApiDescribe("获取用户信息").
		ApiResStruct(user.UserInfo{}).
		ApiResponseTemplate(`{"roles": ["admin"], "code": 0, "avatar":"http://xxxx/aaaa.png", "nickname": "admin"}`)

	User.Pattern("/user/create").Post(handle.CreateUser).Bind(&user.GetAddUser{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	User.Pattern("/password/update").Post(handle.ChangePassword).Bind(&user.ChangePasswod{}).
		AddMidware(midware.JsonToStruct)

	User.Pattern("/get/user").Post(handle.GetUser)
	User.Pattern("/get/project/user").Get(handle.GetProjectUser)
	return User
}
