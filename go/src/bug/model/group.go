package model

type Table_groups struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Ids  string `json:"ids"`
}

type Get_groups struct {
	Id    int64    `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
	Code  int      `json:"statuscode"`
}

type Send_groups struct {
	GroupList []*Get_groups `json:"grouplist"`
	Code      int           `json:"statuscode"`
}
