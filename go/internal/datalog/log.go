package datalog

import (
	"itflow/db"
	"time"

	"github.com/hyahm/golog"
)

type AddLog struct {
	Classify string
	Ip       string
	Username string
	Msg      string
	Action   string
}

func (al *AddLog) Insert() {
	_, err := db.Mconn.Insert("insert into log(exectime,classify,content,ip, username,action) values(?,?,?,?,?,?)",
		time.Now().Unix(), al.Classify, "", al.Ip, al.Username, al.Action,
	)
	if err != nil {
		golog.Error(err)
	}
}
