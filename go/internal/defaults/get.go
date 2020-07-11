package defaults

import (
	"itflow/cache"
)

type RespDefaultStatus struct {
	Created   cache.Status `json:"created"`
	Completed cache.Status `json:"completed"`
	Code      int          `json:"code"`
	Msg       string       `json:"message,omitempty"`
}
