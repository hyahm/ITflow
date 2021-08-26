package bug

type BugManager struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
