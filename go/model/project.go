package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
)

type Project struct {
	Id   int64
	Name string
	Gid  int64
	Uid  int64
}

func (p *Project) Insert(groupname string) error {
	rows, err := db.Mconn.GetRows("select id from usergroup where name=?")
	if err != nil {
		return err
	}
	ids := make([]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			golog.Error(err)
			continue
		}
		ids = append(ids, id)
	}
	rows.Close()
	p.Id, err = db.Mconn.Insert("insert into project(name,ugid,uid) values(?,?,?)", p.Name, strings.Join(ids, ","), p.Uid)
	return err
}

func NewProjectById(id interface{}) (*Project, error) {
	p := &Project{}
	err := db.Mconn.GetOne("select id,name,ugid,uid from project where id=?", id).Scan(&p.Id, &p.Name, &p.Gid, &p.Uid)

	return p, err
}

func NewProjectListCheckId(uid int64) ([]*Project, error) {
	// 获取此用户的项目组
	ps := make([]*Project, 0)
	// 如果是管理员或者创建者，都能看到
	rows, err := db.Mconn.GetRows(`select id,name,ugid,uid from project where uid=? or uid=? or 
		ugid in (select ids from usergroup where ugid=id)`, uid, cache.SUPERID)
	if err != nil {
		golog.Info()
		return nil, err
	}
	for rows.Next() {
		p := &Project{}
		rows.Scan(&p.Id, &p.Name, &p.Gid, &p.Uid)

		ps = append(ps, p)
	}
	rows.Close()
	return ps, nil
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
