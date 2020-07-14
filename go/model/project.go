package model

import (
	"errors"
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
	Uid  int64
}

func (p *Project) Insert() error {
	pid, err := db.Mconn.Insert("insert into project(name,ugid,uid) values(?,?,?)", p.Name, p.Gid, p.Uid)
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
	rows, err := db.Mconn.GetRows("select id,name,ugid,uid from project")
	if err != nil {
		golog.Info()
		return nil, err
	}
	for rows.Next() {
		p := &Project{}
		rows.Scan(&p.Id, &p.Name, &p.Gid, &p.Uid)
		// 如果是管理员或者创建者，都能看到
		if p.Participant(uid) {
			ps = append(ps, p)
		}
	}
	return ps, nil
}

func (p *Project) Participant(uid int64) bool {
	if uid == cache.SUPERID || uid == p.Uid {
		return true
	} else {
		// 返回自由自己权限的
		ug, ok := cache.CacheUGidUserGroup[p.Gid]
		if !ok {
			return false
		}
		for _, v := range strings.Split(ug.Uids, ",") {
			thisUid, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			if thisUid == uid {
				return true
			}
		}
	}
	return false
}

func (p *Project) Update() error {
	golog.Infof("%+v", p)
	_, err := db.Mconn.Update("update project set name=?,ugid=? where id=? and uid=?", p.Name, p.Gid, p.Id, p.Uid)
	return err
}

func (p *Project) Delete() error {
	count, err := db.Mconn.Update("delete from project where id=?", p.Id, p.Uid)
	if count == 0 {
		return errors.New("delete failed")
	}
	return err
}
