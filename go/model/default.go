package model

import (
	"itflow/cache"
	"itflow/db"
)

type DefaultValue struct {
	Created   cache.StatusId // statusid
	Completed cache.StatusId // statusid

}

func (dv *DefaultValue) Update() error {
	_, err := db.Mconn.Update("update defaultvalue set created=?,completed=?", dv.Created, dv.Completed)
	return err
}
