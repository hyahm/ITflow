package routegroup

import (
	"itflow/handle"
	"itflow/internal/response"
	"itflow/internal/user"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// User 用户操作
var User *xmux.GroupRoute

func init() {
	User = xmux.NewGroupRoute()
	User.ApiCreateGroup("user", "用户相关的", "user")
	User.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	User.ApiCodeField("code").ApiCodeMsg("0", "成功")
	User.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	User.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	User.Post("/user/login", handle.Login).Bind(&user.Login{}).
		DelModule(midware.CheckToken).AddModule(midware.JsonToStruct).
		ApiDescribe("用户登录接口").
		ApiDelReqHeader("X-Token").
		ApiReqStruct(user.Login{}).
		ApiResStruct(user.RespLogin{}).
		ApiRequestTemplate(`{"username":"admin", "password": "123456"}`).
		ApiResponseTemplate(`{"username":"admin","token":"sdfhdffffsdfgasdfasdf", "code": 0}`)
	User.Post("/user/logout", handle.LoginOut).
		ApiDescribe("用户退出接口").
		ApiResStruct(response.Response{}).
		ApiSupplement("返回码是大部分公用的")

	User.Get("/user/info", handle.UserInfo).
		ApiDescribe("获取用户信息").
		ApiResStruct(user.UserInfo{}).
		ApiResponseTemplate(`{"roles": ["admin"], "code": 0, "avatar":"http://xxxx/aaaa.png", "nickname": "admin"}`)

	User.Post("/user/create", handle.CreateUser).Bind(&user.GetAddUser{}).
		AddModule(midware.JsonToStruct).ApiDescribe("添加用户").
		ApiRequestTemplate(`{"nickname":"cander","email":"yifan@uupoweremail.com","password":"123456","repassword":"123456","realname":"cander","rolegroup":"all","statusgroup":"aaa","level":2,"position":"aaaaa"}`)

	User.Post("/password/update", handle.ChangePassword).Bind(&user.ChangePasswod{}).
		AddModule(midware.JsonToStruct)

	User.Post("/get/user", handle.GetUser).ApiDescribe("获取所有用户的真实名")
	User.Get("/get/project/user", handle.GetProjectUser).ApiDescribe("根据项目获取关联用户和版本")
}
