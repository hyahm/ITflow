package model

import "itflow/db"

type Role struct {
	Id   int64  `json:"id,omitempty"`
	Role string `json:"role"`
	Info string `json:"info"`
}

func AllRole() ([]*Role, error) {
	rs := make([]*Role, 0)
	rows, err := db.Mconn.GetRows("select name,info from roles")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rl := &Role{}
		rows.Scan(&rl.Role, &rl.Info)
		rs = append(rs, rl)
	}
	return rs, nil
}
