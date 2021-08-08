package model

import (
	"itflow/db"
	"time"

	"github.com/hyahm/golog"
)

type Version struct {
	Id         int64  `json:"id" db:"id,default"`
	ProjectId  int64  `json:"pid" db:"pid"`
	Name       string `json:"name" db:"name"`
	Url        string `json:"urlone" db:"urlone"`
	BakUrl     string `json:"urltwo" db:"urltwo"`
	CreateTime int64  `json:"createtime" db:"createtime"`
	CreateUid  int64  `json:"createuid" db:"createuid"`
}

func GetAllVersion() ([]Version, error) {
	vs := make([]Version, 0)
	err := db.Mconn.Select(&vs, "select * from version")
	return vs, err
}

func (v *Version) Update() error {
	_, err := db.Mconn.UpdateInterface(v, "update version set $set where id=?", v.Id)
	return err
}

func (v *Version) Create() error {
	v.CreateTime = time.Now().Unix()
	ids, err := db.Mconn.InsertInterfaceWithID(v, "insert into version($key) values($value)")
	if err != nil {
		return err
	}
	v.Id = ids[0]
	return nil
}

func (v *Version) GetProjectNameByPid(pid interface{}) ([]string, error) {
	// pid 假设是唯一的，
	rows, err := db.Mconn.GetRows("select name from version where pid=? order by id desc", pid)
	if err != nil {
		return nil, err
	}
	ps := make([]string, 0)
	for rows.Next() {
		rows.Scan(
			&v.Name,
		)
		ps = append(ps, v.Name)
	}
	rows.Close()
	return ps, err
}

func GetVersionIdsByProjectId(pid interface{}) ([]int64, error) {

	rows, err := db.Mconn.GetRows("select id from version where pid=?", pid)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	defer rows.Close()
	vids := make([]int64, 0)
	for rows.Next() {
		var vid int64
		err = rows.Scan(&vid)
		if err != nil {
			continue
		}
		vids = append(vids, vid)
	}
	return vids, nil
}

func GetVersionKeyNameByProjectId(pid interface{}) ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,name from version where pid=?", pid)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	defer rows.Close()
	kns := make([]KeyName, 0)
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			continue
		}
		kns = append(kns, kn)
	}
	return kns, nil
}
