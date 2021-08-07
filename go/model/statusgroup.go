package model

import "itflow/db"

type StatusGroup struct {
	ID        int64   `json:"id" db:"id,default"`
	Name      string  `json:"name" db:"name"`
	StatusIDs []int64 `json:"sids" db:"sids"`
}

func getStatusIDsByUid(uid int64) ([]int64, error) {
	sg := StatusGroup{}
	sql := "select * from statusgroup where id=(select bugsid from jobs where id=(select jid from user where id=?))"
	err := db.Mconn.Select(&sg, sql, uid)
	if err != nil {
		return nil, err
	}
	return sg.StatusIDs, nil
}
