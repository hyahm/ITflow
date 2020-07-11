package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Project struct {
	Id          int64
	Name        string
	Participant string
}

func (p *Project) Insert() (err error) {
	p.Id, err = db.Mconn.Insert("insert into project(name,participant) values(?,?)", p.Name, p.Participant)
	return err
}

func NewProjectById(id interface{}) (*Project, error) {
	p := &Project{}
	err := db.Mconn.GetOne("select id,name,participant from project where id=?", id).Scan(&p.Id, &p.Name, &p.Participant)

	return p, err
}

func NewProjectList() ([]*Project, error) {
	ps := make([]*Project, 0)
	rows, err := db.Mconn.GetRows("select id,name,participant from project")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := &Project{}
		rows.Scan(&p.Id, &p.Name, &p.Participant)
		ps = append(ps, p)
	}
	return ps, nil
}

func (p *Project) Update() error {
	golog.Infof("%+v", p)
	_, err := db.Mconn.Update("update project set name=?,participant=? where id=?", p.Name, p.Participant, p.Id)
	return err
}

func (p *Project) Delete() error {
	_, err := db.Mconn.Update("delete from project where id=?", p.Id)
	return err
}
