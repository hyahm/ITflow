package defaults

import (
	"itflow/cache"
)

type RespDefaultStatus struct {
	Status cache.Status `json:"defaultstatus"`
	Code   int          `json:"code"`
	Msg    string       `json:"message,omitempty"`
}
