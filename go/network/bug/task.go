package bug

type SearchParam struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Level   string   `json:"level"`
	Project string   `json:"project"`
	Title   string   `json:"title"`
	Status  []string `json:"status"`
}
