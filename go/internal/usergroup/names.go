package usergroup

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type RespUserGroupName struct {
	UserGroupNames []string `json:"usergroupnames"`
	Code           int      `json:"code"`
	Msg            string   `json:"msg"`
}

func (rugn *RespUserGroupName) Marshal() []byte {
	send, err := json.Marshal(rugn)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (rugn *RespUserGroupName) Error(msg string) []byte {
	rugn.Code = 1
	rugn.Msg = msg
	return rugn.Marshal()
}

func (rugn *RespUserGroupName) ErrorE(err error) []byte {
	return rugn.Error(err.Error())
}
