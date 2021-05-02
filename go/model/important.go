package model

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type Important struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ResposeImportant struct {
	ImportantList []*Important `json:"importantlist"`
	Code          int          `json:"code"`
	Msg           string       `json:"msg"`
}

func (li *ResposeImportant) Marshal() []byte {
	send, err := json.Marshal(li)
	if err != nil {
		golog.Error(err)
	}
	return send
}
func (li *ResposeImportant) Error(msg string) []byte {
	li.Code = 1
	li.Msg = msg
	return li.Marshal()
}

func (li *ResposeImportant) ErrorE(err error) []byte {

	return li.Error(err.Error())
}
