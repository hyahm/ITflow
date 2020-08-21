package usergroup

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type RespUserGroupList struct {
	UserGroupList []*RespUserGroup `json:"usergrouplist"`
	Code          int              `json:"code"`
	Msg           string           `json:"msg,omitempty"`
}

func (rugl *RespUserGroupList) Marshal() []byte {
	send, err := json.Marshal(rugl)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (rugl *RespUserGroupList) Error(msg string) []byte {
	rugl.Code = 1
	rugl.Msg = msg
	return rugl.Marshal()
}

func (rugl *RespUserGroupList) ErrorE(err error) []byte {
	return rugl.Error(err.Error())
}
