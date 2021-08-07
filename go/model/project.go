package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
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

func (p *Project) Insert(groupname string) error {

	rows, err := db.Mconn.GetRows("select id from usergroup where name=?", groupname)
	if err != nil {
		golog.Error(err)
		return err
	}
	ids := make([]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			golog.Info(err)
			continue
		}
		ids = append(ids, id)
	}
	rows.Close()
	p.Id, err = db.Mconn.Insert("insert into project(name,ugid,uid) values(?,?,?)", p.Name, strings.Join(ids, ","), p.Uid)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			return db.DuplicateErr
		}

	}
	return err
}

func NewProjectById(id interface{}) (*Project, error) {
	p := &Project{}
	err := db.Mconn.GetOne("select id,name,ugid,uid from project where id=?",
		id).Scan(&p.Id, &p.Name, &p.UGid, &p.Uid)

	return p, err
}

func NewProjectListCheckId(uid int64) ([]*Project, error) {
	// 获取此用户的项目组
	ps := make([]*Project, 0)
	// 如果是管理员或者创建者，都能看到
	err := db.Mconn.Select(&ps, `select * from project where uid=? or uid=? or 
		ugid in (select id from usergroup where json_contains(ugid, json_array(?)))`, uid, cache.SUPERID, uid)

	return ps, err
}

func (p *Project) Update(groupname string) error {
	golog.Infof("%+v", p)
	_, err := db.Mconn.Update("update project set name=?,ugid=(select ifnull(min(id),0) from usergroup where name=?) where id=? and uid=?", p.Name, groupname, p.Id, p.Uid)
	return err
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
