package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/golog"
)

type Project struct {
	Id   cache.ProjectId
	Name cache.Project
	Gid  int64
	Uid  int64
}

func (p *Project) Insert(groupname string) error {
	pid, err := db.Mconn.Insert("insert into project(name,ugid,uid) values(?,(select ifnull(min(id),0) from usergroup where name=?),10)", p.Name, groupname, p.Uid)
	p.Id = cache.ProjectId(pid)
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
