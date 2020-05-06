package model

type Table_buggroup struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Sids string `json:"sids"`
}

type Table_status struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type List_StatusName struct {
	StatusList []string `json:"statuslist"`
	Code       int      `json:"code"`
}
