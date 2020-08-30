package model

import (
	"itflow/db"
)

type Version struct {
	Id         int64
	ProjectId  int64
	Name       string
	Url        string
	BakUrl     string
	CreateTime int64
}

func (v *Version) GetProjectNameByPid(pid interface{}) ([]string, error) {
	// pid 假设是唯一的，
	rows, err := db.Mconn.GetRows("select name from version where pid=? order by id desc", pid)
	if err != nil {
		return nil, err
	}
	ps := make([]string, 0)
	for rows.Next() {
		rows.Scan(
			&v.Name,
		)
		ps = append(ps, v.Name)
	}
	return ps, err
}
