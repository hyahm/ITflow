package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type DefaultValue struct {
	Created   string `json:"created"`   // statusid`json:"created"`
	Completed string `json:"completed"` // statusid

}

func (dv *DefaultValue) Update() error {
	db.Mconn.OpenDebug()
	_, err := db.Mconn.Update(`update defaultvalue set created=(select ifnull(min(id),0) from status where name=?),
	completed=(select ifnull(min(id),0) from status where name=?) `, dv.Created, dv.Completed)
	golog.Info(db.Mconn.GetSql())
	return err
}
