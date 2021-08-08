package model

import (
	"itflow/db"
)

type DefaultValue struct {
	Created   int64 `json:"created" db:"created"`     // statusid`json:"created"`
	Completed int64 `json:"completed" db:"completed"` // statusid

}

func (dv *DefaultValue) Update() error {
	_, err := db.Mconn.UpdateInterface(dv, "update defaultvalue set $set")
	return err
}

func GetDefaultValue() (DefaultValue, error) {
	dv := DefaultValue{}
	err := db.Mconn.Select(&dv, "select * from defaultvalue")
	return dv, err
}
