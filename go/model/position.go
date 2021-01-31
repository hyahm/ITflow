package model

type Table_jobs struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	HypoName    string `json:"hyponame"`
	StatusGroup string `json:"statusgroup"`
	RoleGroup   string `json:"rolegroup"`
	Hid         int64  // 临时变量
}

type Data_jobs struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Hyponame string `json:"hyponame"`
	Code     int    `json:"code"`
}

type List_jobs struct {
	Positions []*Table_jobs `json:"positions"`
	Code      int           `json:"code"`
}

type Update_jobs struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	StatusGroup string `json:"statusgroup"`
	RoleGroup   string `json:"rolegroup"`
	Hyponame    string `json:"hyponame"`
}
