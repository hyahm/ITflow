package status

import (
	"encoding/json"

	"github.com/hyahm/golog"
)

type Status struct {
	CheckStatus []string `json:"checkstatus"`
	Code        int      `json:"code"`
	Msg         string   `json:"msg"`
}

func (s *Status) Marshal() []byte {
	send, err := json.Marshal(s)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (s *Status) Error(msg string) []byte {
	s.Code = 1
	s.Msg = msg
	return s.Marshal()
}

func (s *Status) ErrorE(err error) []byte {
	return s.Error(err.Error())
}

type ChangeStatus struct {
	Status []string `json:"status"`
	Code   int      `json:"code"`
}
