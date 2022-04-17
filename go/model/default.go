package model

import (
	"itflow/db"
)

type DefaultValue struct {
	Created   int64 `json:"created" db:"created,default"`     // statusid`json:"created"`
	Completed int64 `json:"completed" db:"completed,default"` // statusid
	Pass      int64 `json:"pass" db:"pass,default"`           // statusid
	Receive   int64 `json:"receive" db:"receive,default"`     // statusid

}

func (dv *DefaultValue) Update() error {
	result := db.Mconn.UpdateInterface(dv, "update defaultvalue set $set")
	return result.Err
}

func GetDefaultValue() (DefaultValue, error) {
	dv := DefaultValue{}
	result := db.Mconn.Select(&dv, "select * from defaultvalue")
	return dv, result.Err
}
