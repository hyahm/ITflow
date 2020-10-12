package model

import (
	"database/sql"
	"errors"
	"itflow/db"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
)

// RoleGroup: 角色组表， 由管理员分配， 管理可以操作的页面, 与用户rid关联
type RoleGroup struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Permids string `json:"permids"`
}

func NewRoleGroup(uid int64) (*RoleGroup, error) {
	// 通过uid 来获取rid
	rg := &RoleGroup{}
	err := db.Mconn.GetOne("select r.name, r.permids from rolegroup as r join user as u on u.id=? and u.rid=r.id", uid).Scan(&rg.Name, &rg.Permids)
	if err != nil {
		golog.Error(err)
		return nil, err
	}

	return rg, nil
}

func (rg *RoleGroup) Insert() error {
	var err error
	rg.ID, err = db.Mconn.Insert("insert into rolegroup(name,permids) values(?,?)", rg.Name, rg.Permids)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			return db.DuplicateErr
		}

	}
	return err
}

func (rg *RoleGroup) Update() (err error) {
	golog.Infof("%+v", rg)
	_, err = db.Mconn.Update("update rolegroup set name=? where id=?", rg.Name, rg.ID)
	if err != nil {
		return
	}
	// 查询到permids
	err = db.Mconn.GetOne("select permids from rolegroup where id=?", rg.ID).Scan(&rg.Permids)
	return
}

func CheckRoleNameInGroup(name string, rid *int64) error {
	err := db.Mconn.GetOne("select id from rolegroup where name=?",
		strings.Trim(name, " ")).Scan(rid)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("不存在此角色组")
		}
		return err
	}

	return nil
}
