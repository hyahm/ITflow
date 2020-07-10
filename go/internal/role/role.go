package role

import (
	"encoding/json"
	"itflow/model"

	"github.com/hyahm/golog"
)

type Role struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	rolelist string `json:"rolelist"`
}

type RespRoles struct {
	Roles []*model.Role `json:"roles"`
	Code  int           `json:"code"`
	Msg   string        `json:"message"`
}

func (rr *RespRoles) Marshal() []byte {
	send, _ := json.Marshal(rr)
	return send
}

func (rr *RespRoles) List() []byte {
	ar, err := model.AllRole()
	if err != nil {
		rr.Msg = err.Error()
		rr.Code = 1
		golog.Error(err)
		return rr.Marshal()
	}
	rr.Roles = ar
	return rr.Marshal()
}
