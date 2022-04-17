package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/gomysql"
)

type Project struct {
	Id   int64  `json:"id" db:"id,default"`
	Name string `json:"name" db:"name"`
	UGid int64  `json:"ugid" db:"ugid"`
	Uid  int64  `json:"uid" db:"uid"`
}

func GetProjectKeyName(uid int64) ([]KeyName, error) {
	ugids, err := GetUserGroupIds(uid)
	if err != nil {
		return nil, nil
	}
	rows, err := db.Mconn.GetRowsIn("select id,name from project where ugid in (?) or uid=?", ugids, uid)
	if err != nil {
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

func (p *Project) Insert() error {

	result := db.Mconn.InsertInterfaceWithID(p, "insert into project($key) values($value)")
	if result.Err != nil {
		return result.Err
	}
	p.Id = result.LastInsertId
	return nil
}

func NewProjectById(id interface{}) (*Project, error) {
	p := &Project{}
	err := db.Mconn.GetOne("select id,name,ugid,uid from project where id=?",
		id).Scan(&p.Id, &p.Name, &p.UGid, &p.Uid)

	return p, err
}

func GetAllProjects(uid int64) ([]*Project, error) {
	// 获取此用户的项目组
	ps := make([]*Project, 0)
	if uid == cache.SUPERID {
		result := db.Mconn.Select(&ps, `select * from project`)
		return ps, result.Err
	} else {
		// 如果是管理员或者创建者，都能看到
		result := db.Mconn.Select(&ps, `select * from project where uid=? or uid=? or 
ugid in (select id from usergroup where json_contains(ugid, json_array(?)))`, uid, cache.SUPERID, uid)

		return ps, result.Err
	}

}

func (p *Project) Update(uid int64) error {
	var result gomysql.Result
	if uid == cache.SUPERID {
		result = db.Mconn.UpdateInterface(p, "update project set $set where id=?", p.Id)
	} else {
		result = db.Mconn.UpdateInterface(p, "update project set $set where id=? and uid=?", p.Id, uid)
	}
	return result.Err
}

func (p *Project) Delete() error {
	result := db.Mconn.Update("delete from project where id=?", p.Id, p.Uid)
	if result.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return result.Err
}

func GetUserGroupId(pid interface{}) (int64, error) {
	var ugid int64
	err := db.Mconn.GetOne("select ugid from project where id=?", pid).Scan(&ugid)
	return ugid, err
}
