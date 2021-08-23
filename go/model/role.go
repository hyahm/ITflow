package model

import (
	"itflow/db"
)

type Role struct {
	Id   int64  `json:"id,omitempty" db:"id,default"`
	Role string `json:"role" db:"role,default"`
	Info string `json:"info" db:"info,default"`
}

func AllRole() ([]Role, error) {
	roles := make([]Role, 0)
	err := db.Mconn.Select(&roles, "select * from roles")
	return roles, err
}
