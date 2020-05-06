package user

type RespLogin struct {
	UserName string `json:"username" type:"string" need:"否" default:"" information:"用户名"`
	Token    string `json:"token" type:"string" need:"否" default:"" information:"token"`
	Code     int    `json:"code" type:"int" need:"是" default:"0" information:"返回码 0: 成功， 其他的失败"`
	Msg      string `json:"msg" type:"string" need:"是" default:"" information:"错误信息"`
}
