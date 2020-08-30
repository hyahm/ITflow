package model

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type Importants struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Data_importants struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

type List_importants struct {
	ImportantList []*Importants `json:"importantlist"`
	Code          int           `json:"code"`
	Msg           string        `json:"msg"`
}

func (li *List_importants) Marshal() []byte {
	send, err := json.Marshal(li)
	if err != nil {
		golog.Error(err)
	}
	return send
}
func (li *List_importants) Error(msg string) []byte {
	li.Code = 1
	li.Msg = msg
	return li.Marshal()
}

func (li *List_importants) ErrorE(err error) []byte {

	return li.Error(err.Error())
}
