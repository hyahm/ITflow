package bug

type GetArticle struct {
	Status      string   `json:"status"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Id          int      `json:"id"`
	Selectusers []string `json:"selectuser"`
	Important   string   `json:"important"`
	Level       string   `json:"level"`
	Projectname string   `json:"projectname"`
	Envname     string   `json:"envname"`
	Version     string   `json:"version"`
}
