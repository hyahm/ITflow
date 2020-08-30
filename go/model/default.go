package model

import (
	"itflow/db"
)

type DefaultValue struct {
	Created   string `json:"created"`   // statusid`json:"created"`
	Completed string `json:"completed"` // statusid

}

func (dv *DefaultValue) Update() (int64, int64, error) {
	var createdId, completedId int64
	_, err := db.Mconn.Update(`update defaultvalue set created=(select ifnull(min(id),0) from status where name=?),
	completed=(select ifnull(min(id),0) from status where name=?) `, dv.Created, dv.Completed)

	err = db.Mconn.GetOne("select created,completed from defaultvalue").Scan(&createdId, &completedId)
	return createdId, completedId, err
}
