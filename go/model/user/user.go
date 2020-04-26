package user

// 用户登录

type Login struct {
	Username string `json:"username" type:"string" need:"是" default:"" information:"用户名"`
	Password string `json:"password"  type:"string" need:"是" default:"" information:"密码"`
}

type RespLogin struct {
	UserName string `json:"username" type:"string" need:"否" default:"" information:"用户名"`
	Token    string `json:"token" type:"string" need:"否" default:"" information:"token"`
	Code     int    `json:"code" type:"int" need:"是" default:"0" information:"返回码 0: 成功， 其他的失败"`
	Msg      string `json:"msg" type:"string" need:"是" default:"" information:"错误信息"`
}
