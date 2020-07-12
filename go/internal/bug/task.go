package bug

import (
	"itflow/cache"
)

type SearchParam struct {
	Page    int           `json:"page"`
	Limit   int           `json:"limit"`
	Level   cache.Level   `json:"level"`
	Project cache.Project `json:"project"`
	Title   string        `json:"title"`
	Status  []string      `json:"status"`
}
