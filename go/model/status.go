package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Status struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (status *Status) Names() ([]string, error) {
	names := make([]string, 0)
	rows, err := db.Mconn.GetRows("select name from status")
	if err != nil {
		return names, err
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}
	return names, nil
}

func (status *Status) List() ([]*Status, error) {
	ss := make([]*Status, 0)
	rows, err := db.Mconn.GetRows("select id,name from status")
	if err != nil {
		golog.Error(err)
		return ss, err
	}
	for rows.Next() {
		st := &Status{}
		rows.Scan(&st.ID, &st.Name)
		ss = append(ss, st)
	}
	return ss, nil
}
