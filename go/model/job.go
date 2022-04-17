package model

import (
	"database/sql"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

type Job struct {
	Id          int64  `json:"id" db:"id,default"`
	Name        string `json:"name" db:"name"`
	Level       int    `json:"level" db:"level"`      // 1 是管理者， 2 是普通员工
	HypoName    int64  `json:"hypo" db:"hypo"`        //  管理者id
	StatusGroup int64  `json:"statusgroup" db:"sgid"` // 状态组
	RoleGroup   int64  `json:"rolegroup" db:"rgid"`   // 角色组
}

func (job *Job) Insert() error {
	result := db.Mconn.InsertInterfaceWithID(job, "insert into jobs($key) values($value)")
	if result.Err != nil {
		return result.Err
	}
	job.Id = result.LastInsertId
	return nil
}

func DeleteJob(id, uid interface{}) (err error) {
	var result gomysql.Result
	if uid == cache.SUPERID {
		result = db.Mconn.Delete("delete from jobs where id=?", id)
	} else {
		result = db.Mconn.Delete("delete from jobs where id=? and hypo=?", id, uid)
	}

	return result.Err
}

func (job *Job) Update() error {
	result := db.Mconn.UpdateInterface(job, "update jobs set $set where id=?", job.Id)
	return result.Err
}

func GetAllPositions() ([]Job, error) {
	jobs := make([]Job, 0)
	result := db.Mconn.Select(&jobs, "select * from jobs")
	return jobs, result.Err
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

func GetManagerKeyName(uid int64) ([]KeyName, error) {
	var err error
	var rows *sql.Rows
	if uid == goconfig.ReadInt64("adminid") {
		rows, err = db.Mconn.GetRows("select id,name from jobs where level=1")
	} else {
		rows, err = db.Mconn.GetRows("select id,name from jobs where level=1 and uid=?", uid)
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
