package model

type Table_importants struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Data_importants struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code int    `json:"code"`
}

type List_importants struct {
	ImportantList []*Table_importants `json:"importantlist"`
	Code          int                 `json:"code"`
}

type Update_importants struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
