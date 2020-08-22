package user

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type RespLogin struct {
	ID       int64  `json:"id"`
	UserName string `json:"username" type:"string" need:"否" default:"" information:"用户名"`
	Token    string `json:"token" type:"string" need:"否" default:"" information:"token"`
	Code     int    `json:"code" type:"int" need:"是" default:"0" information:"返回码 0: 成功， 其他的失败"`
	Msg      string `json:"msg" type:"string" need:"是" default:"" information:"错误信息"`
}

func (rl *RespLogin) Marshal() []byte {
	send, err := json.Marshal(rl)
	if err != nil {
		golog.Error(err)

	}
	return send
}

func (rl *RespLogin) Error(msg string) []byte {
	rl.Code = 1
	rl.Msg = msg
	return rl.Marshal()
}

func (rl *RespLogin) ErrorE(err error) []byte {
	return rl.Error(err.Error())
}
