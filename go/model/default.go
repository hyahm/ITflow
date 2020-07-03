package model

import (
	"itflow/cache"
	"itflow/db"
)

type DefaultValue struct {
	Status cache.StatusId // statusid
}

func (dv *DefaultValue) Update() error {
	_, err := db.Mconn.Update("update defaultvalue set status=?", dv.Status)
	return err
}
