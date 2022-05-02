package user

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type RespLogin struct {
	ID       int64  `json:"id"`
	UserName string `json:"username" type:"string" need:"否" default:"" information:"用户名"`
	Token    string `json:"token" type:"string" need:"否" default:"" information:"token"`
}

func (rl *RespLogin) Marshal() []byte {
	send, err := json.Marshal(rl)
	if err != nil {
		golog.Error(err)

	}
	return send
}

func (rl *RespLogin) Error(msg string) []byte {
	return rl.Marshal()
}

func (rl *RespLogin) ErrorE(err error) []byte {
	return rl.Error(err.Error())
}
