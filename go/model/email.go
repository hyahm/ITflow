package model

type Email struct {
	Host      string `json:"host"`
	Enable    bool   `json:"enable"`
	Id        int64  `json:"id"`
	Port      int    `json:"port"`
	EmailAddr string `json:"emailaddr"`
	Password  string `json:"password"`
}
