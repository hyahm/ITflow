package model

import (
	"database/sql"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/golog"
)

// keyvalue 值
type KeyName struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func GetUserKeyName(uid int64) ([]KeyName, error) {
	var rows *sql.Rows
	var err error
	if uid == cache.SUPERID {
		rows, err = db.Mconn.GetRows("select id,nickname from user")
		if err != nil {
			return nil, err
		}
	} else {
		// 拿到能操作的用户id
		rows, err = db.Mconn.GetRows("select id,realname from user where id in (select id from jobs where hypo =( select hypo from jobs where id=(select jid from user where id=?)))", uid)
		if err != nil {
			golog.Error(err)
			return nil, err
		}

	}

	defer rows.Close()
	kns := make([]KeyName, 0)
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			continue
		}
		kns = append(kns, kn)
	}
	return kns, nil
}
