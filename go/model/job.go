package model

import (
	"database/sql"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/golog"
)

type Job struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Level       int    `json:"level"`  // 1 是管理者， 2 是普通员工
	HypoName    int64  `json:"hypo"`   //  管理者id
	StatusGroup int64  `json:"bugsid"` // 状态组
	RoleGroup   int64  `json:"rid"`    // 角色组
}

func GetJobIdsByJobId(jid int64) ([]int64, error) {
	// 通过jid 来获取 能管理的 职位 的id
	rows, err := db.Mconn.GetRows("select id from jobs where hypo=( select hypo from jobs where id=?)", jid)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	defer rows.Close()
	jobs := make([]int64, 0)
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			golog.Error(err)
			continue
		}
		jobs = append(jobs, id)
	}
	return jobs, nil
}

type Jobs struct {
	Positions []*Job `json:"positions"`
	Code      int    `json:"code"`
}

func GetJobKeyNameByUid(uid int64) ([]KeyName, error) {
	var rows *sql.Rows
	var err error
	if uid == cache.SUPERID {
		rows, err = db.Mconn.GetRows("select id,name from jobs")
	} else {
		rows, err = db.Mconn.GetRows("select id,name from jobs where hypo=?", uid)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kns := make([]KeyName, 0)
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		kns = append(kns, kn)
	}
	return kns, nil
}
