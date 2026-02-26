package model

import (
	"errors"
	"itflow/cache"
	"itflow/db"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/gosql"
)

type Project struct {
	Id      int64     `json:"id" gorm:"primaryKey" form:"id"`
	Name    string    `json:"name" gorm:"column:name"`
	Created time.Time `json:"created" gorm:"column:created"`
	Updated time.Time `json:"updated" gorm:"column:updated"`
	Uid     int64     `json:"uid" gorm:"column:uid"`
}

func (Project) TableName() string {
	return "project"
}

func GetProjectKeyName(uid int64) ([]KeyName, error) {
	ugm := ProjectUserMap{
		Uid: uid,
	}
	projects, err := ugm.GetUserGroupIds()
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	if len(projects) == 0 {
		return make([]KeyName, 0), nil
	}
	kns := make([]KeyName, 0)
	err = db.Gorm.Table("project").Select("id", "name").Where("id in ?", projects).Find(&kns).Error
	return kns, err
}

func (p *Project) Insert() error {
	p.Created = time.Now()
	p.Updated = time.Now()
	return db.Gorm.Create(p).Error
}

func NewProjectById(id interface{}) (*Project, error) {
	p := &Project{}
	// err := db.Mconn.GetOne("select id,name,ugid,uid from project where id=?",
	// 	id).Scan(&p.Id, &p.Name, &p.UGid, &p.Uid)

	return p, nil
}

func (p *Project) GetAllProjects(uid int64) ([]Project, error) {
	ps := make([]Project, 0)
	if p.Uid <= 0 {
		return ps, errors.New("uid not found")
	}
	// 获取此用户的项目组

	// query := db.Gorm.Table(p.TableName())
	query := db.Gorm.Table(p.TableName())
	if uid != cache.SUPERID {
		pum := UserGroupMap{}
		pids, err := pum.GetProjectIdsByUid(uid)
		if err != nil {
			return ps, err
		}
		if len(pids) == 0 {
			return ps, err
		}
		query = query.Where("id in ?", pids)
	}
	err := query.Order("created desc").Find(&ps).Error
	golog.Error(err)
	return ps, err

}

func (p *Project) Update(uid int64) error {
	var result gosql.Result
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
