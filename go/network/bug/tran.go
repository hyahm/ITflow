package bug

type BugData struct {
	Importance  string   `json:"importance"`
	Title       string   `json:"title"`
	Level       string   `json:"level"`
	Version     string   `json:"version"`
	Projectname string   `json:"projectname"`
	Env         string   `json:"env"`
	Handle      []string `json:"handle"`
	Content     string   `json:"content"`
	Code        int      `json:"code"`
}
