package model

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type Table_level struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Data_level struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

type List_levels struct {
	Levels []*Table_level `json:"levels"`
	Code   int            `json:"code"`
	Msg    string         `json:"msg"`
}

type Update_level struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	OldName string `json:"oldname"`
}

func (ll *List_levels) Marshal() []byte {
	send, err := json.Marshal(ll)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (ll *List_levels) Error(msg string) []byte {
	ll.Code = 1
	ll.Msg = msg
	return ll.Marshal()
}

func (ll *List_levels) ErrorE(err error) []byte {
	return ll.Error(err.Error())
}
