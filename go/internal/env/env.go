package env

import (
	"encoding/json"
	"itflow/cache"

	"github.com/hyahm/golog"
)

type Envlist struct {
	Elist []*Env `json:"envlist"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
}
type Env struct {
	Id      cache.EnvId `json:"id"`
	EnvName cache.Env   `json:"envname"`
}

func (el *Envlist) Marshal() []byte {
	send, err := json.Marshal(el)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (el *Envlist) Error(msg string) []byte {
	el.Code = 1
	el.Msg = msg
	return el.Marshal()
}

func (el *Envlist) ErrorE(err error) []byte {
	return el.Error(err.Error())
}
