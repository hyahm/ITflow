package model

type Table_log struct {
	Id       int64  `json:"id"`
	Exectime int64  `json:"exectime"`
	Classify string `json:"classify"`
	Content  string `json:"content"`
	Ip       string `json:"ip"`
}

type List_log struct {
	LogList []*Table_log `json:"loglist"`
	Code    int          `json:"code"`
	Count   int          `json:"count"`
	Page    int          `json:"page"`
}
