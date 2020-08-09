package model

import (
	"itflow/db"
)

type UserGroups struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Ids  string `json:"ids"`
	Uid  int64
}

func (ug *UserGroups) Update() error {
	gsql := "update usergroup set name=?,ids=? where id=? "
	_, err := db.Mconn.Update(gsql, ug.Name, ug.Ids, ug.Id)
	if err != nil {
		return err
	}

	return nil
}
