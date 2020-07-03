package status

import (
	"itflow/cache"
)

type Status struct {
	CheckStatus cache.StatusList `json:"checkstatus"`
	Code        int              `json:"code"`
}

type ChangeStatus struct {
	Status []string `json:"status"`
	Code   int      `json:"code"`
}
