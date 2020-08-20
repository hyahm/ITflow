package defaults

import (
	"encoding/json"
	"itflow/cache"

	"github.com/hyahm/golog"
)

type RespDefaultStatus struct {
	Created   cache.Status `json:"created"`
	Completed cache.Status `json:"completed"`
	Code      int          `json:"code"`
	Msg       string       `json:"message,omitempty"`
}

func (rds *RespDefaultStatus) Marshal() []byte {
	send, err := json.Marshal(rds)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (rds *RespDefaultStatus) Error(msg string) []byte {
	rds.Code = 1
	rds.Msg = msg
	return rds.Marshal()
}

func (rds *RespDefaultStatus) ErrorE(err error) []byte {
	return rds.Error(err.Error())
}
