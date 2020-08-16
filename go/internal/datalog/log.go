package datalog

import (
	"itflow/db"
	"time"

	"github.com/hyahm/golog"
)

func InsertLog(classify, content, ip, username, action string) {
	_, err := db.Mconn.Insert("insert into log(exectime,classify,content,ip, username,action) values(?,?,?,?,?,?)",
		time.Now().Unix(), classify, content, ip, username, action,
	)
	if err != nil {
		golog.Error(err)
	}
}
