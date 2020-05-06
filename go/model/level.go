package model

type Table_level struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Data_level struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

type List_levels struct {
	Levels []*Table_level `json:"levels"`
	Code   int            `json:"code"`
}

type Update_level struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	OldName string `json:"oldname"`
}
