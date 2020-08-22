package model

import (
	"itflow/classify"
	"itflow/db"
	"strings"
	"time"

	"github.com/hyahm/golog"
)

type Log struct {
	Id       int64  `json:"id"`
	Exectime int64  `json:"exectime"`
	Classify string `json:"classify"`
	Ip       string `json:"ip"`
	Uid      int64  `json:""`
	Action   string `json:"action"`
}

// `id` bigint(20) NOT NULL AUTO_INCREMENT,
// `exectime` bigint(20) DEFAULT '0',
// `classify` varchar(30) NOT NULL DEFAULT '',
// `content` text,
// `ip` varchar(40) DEFAULT '',
// `username` varchar(50) DEFAULT '',
// `action` varchar(50) DEFAULT '',

func InsertLog(classify classify.Classify, ip, action string, uid int64) {
	ip = strings.Split(ip, ":")[0]
	_, err := db.Mconn.Insert("insert into log(exectime,classify,ip, uid,action) values(?,?,?,?,?)",
		time.Now().Unix(), classify, ip, uid, action,
	)
	if err != nil {
		golog.Error(err)
	}
}
