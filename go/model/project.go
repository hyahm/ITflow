package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"
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

	ids, err := db.Mconn.InsertInterfaceWithID(p, "insert into project($key) values($value)")
	if err != nil {
		return err
	}
	p.Id = ids[0]
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
		err := db.Mconn.Select(&ps, `select * from project`)
		return ps, err
	} else {
		// 如果是管理员或者创建者，都能看到
		err := db.Mconn.Select(&ps, `select * from project where uid=? or uid=? or 
ugid in (select id from usergroup where json_contains(ugid, json_array(?)))`, uid, cache.SUPERID, uid)

		return ps, err
	}

}

func (p *Project) Update(uid int64) (err error) {
	if uid == cache.SUPERID {
		_, err = db.Mconn.UpdateInterface(p, "update project set $set where id=?", p.Id)
	} else {
		_, err = db.Mconn.UpdateInterface(p, "update project set $set where id=? and uid=?", p.Id, uid)
	}
	return
}

func (p *Project) Delete() error {
	count, err := db.Mconn.Update("delete from project where id=?", p.Id, p.Uid)
	if count == 0 {
		return errors.New("delete failed")
	}
	return err
}

func GetUserGroupId(pid interface{}) (int64, error) {
	var ugid int64
	err := db.Mconn.GetOne("select ugid from project where id=?", pid).Scan(&ugid)
	return ugid, err
}
