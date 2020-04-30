package model

type ArticleList struct {
	ID          int      `json:"id"`
	Date        int64    `json:"date"`
	Author      string   `json:"author"`
	Importance  string   `json:"importance"`
	Status      string   `json:"status"`
	Title       string   `json:"title"`
	Action      string   `json:"action"`
	Dustbin     int      `json:"dustbin"`
	Level       string   `json:"level"`
	Projectname string   `json:"projectname"`
	Env         string   `json:"env"`
	Handle      []string `json:"handle"`
}

type AllArticleList struct {
	Al    []*ArticleList `json:"articlelist"`
	Code  int            `json:"code"`
	Count int            `json:"total"`
	Page  int            `json:"page"`
}
