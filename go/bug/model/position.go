package model

type Table_jobs struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Hypo  string `json:"hypo"`
}

type Data_jobs struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Hyponame string `json:"hyponame"`
	Code     int    `json:"statuscode"`
}

type List_jobs struct {
	Positions []*Table_jobs `json:"positions"`
	Code      int           `json:"statuscode"`
}

type Update_jobs struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Level int    `json:"level"`
	Hypo  string `json:"hypo"`
}
