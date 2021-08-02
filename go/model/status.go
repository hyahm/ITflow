package model

import (
	"database/sql"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
)

type Status struct {
	ID   int64  `json:"id" db:"id,default"`
	Name string `json:"name" db:"name"`
}

func (status *Status) Names() ([]string, error) {
	names := make([]string, 0)
	rows, err := db.Mconn.GetRows("select name from status")
	if err != nil {
		if err == sql.ErrNoRows {
			return names, nil
		}
		return names, err
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}
	rows.Close()
	return names, nil
}

func GetStatusList() ([]*Status, error) {
	ss := make([]*Status, 0)
	err := db.Mconn.Select(&ss, "select * from status")
	if err != nil {
		golog.Error(err)
		return ss, err
	}
	return ss, nil
}

func GetMyStatusList(id interface{}) ([]string, error) {
	var sids string
	err := db.Mconn.GetOne("select showstatus from user where id=?", id).Scan(&sids)
	if err != nil {
		return nil, err
	}
	return strings.Split(sids, ","), nil
}
