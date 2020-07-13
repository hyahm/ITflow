package model

import (
	"itflow/cache"
	"itflow/db"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type Project struct {
	Id   cache.ProjectId
	Name cache.Project
	Gid  int64
}

func (p *Project) Insert() error {
	pid, err := db.Mconn.Insert("insert into project(name,gid) values(?,?)", p.Name, p.Gid)
	p.Id = cache.ProjectId(pid)
	return err
}

func NewProjectById(id interface{}) (*Project, error) {
	p := &Project{}
	err := db.Mconn.GetOne("select id,name,gid from project where id=?", id).Scan(&p.Id, &p.Name, &p.Gid)

	return p, err
}

func NewProjectListCheckId(uid int64) ([]*Project, error) {
	// 获取此用户的项目组
	ps := make([]*Project, 0)
	rows, err := db.Mconn.GetRows("select id,name,gid from project")
	if err != nil {
		golog.Info()
		return nil, err
	}
	for rows.Next() {
		p := &Project{}
		rows.Scan(&p.Id, &p.Name, &p.Gid)
		if uid == cache.SUPERID {
			ps = append(ps, p)
		} else {
			// 返回自由自己权限的
			for _, v := range strings.Split(cache.CacheGidGroup[p.Gid].Uids, ",") {
				thisUid, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					continue
				}
				if thisUid == uid {
					ps = append(ps, p)
					break
				}
			}
		}

	}
	return ps, nil
}

func (p *Project) Update() error {
	golog.Infof("%+v", p)
	_, err := db.Mconn.Update("update project set name=?,gid=? where id=?", p.Name, p.Gid, p.Id)
	return err
}

func (p *Project) Delete() error {
	_, err := db.Mconn.Update("delete from project where id=?", p.Id)
	return err
}
