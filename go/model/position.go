package model

type Job struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`
	HypoName    string `json:"hyponame"`
	StatusGroup string `json:"statusgroup"`
	RoleGroup   string `json:"rolegroup"`
	Hid         int64  // 临时变量
}

type Jobs struct {
	Positions []*Job `json:"positions"`
	Code      int    `json:"code"`
}
