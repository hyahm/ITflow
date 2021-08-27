package model

type Email struct {
	Host     string `json:"host" db:"host"`
	Enable   bool   `json:"enable" db:"enable"`
	Id       int64  `json:"id" db:"id,default""`
	Port     int    `json:"port" db:"port"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	NickName string `json:"nickname" db:"nickname"`
}
