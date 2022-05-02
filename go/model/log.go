package model

import (
	"itflow/db"
)

type Log struct {
	Id       int64  `json:"id" db:"id"`
	Exectime int64  `json:"exectime" db:"exectime"`
	Classify string `json:"classify" db:"classify"`
	Ip       string `json:"ip" db:"ip"`
	Uid      int64  `json:"uid" db:"uid"`
	Action   string `json:"action" db:"action"`
}

func (log *Log) Insert() error {
	result := db.Mconn.InsertInterfaceWithoutID(log, "insert into log($key) values($value)")
	return result.Err
}
