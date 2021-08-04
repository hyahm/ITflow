package model

import (
	"itflow/db"

	"github.com/hyahm/goconfig"
)

type UserGroup struct {
	Id   int64   `json:"id" db:"id,default"`
	Name string  `json:"name" db:"name"`
	Uids []int64 `json:"uids" db:"uids"`
	Uid  int64   `json:"uid" db:"uid"`
}

func (ug *UserGroup) Update() error {
	gsql := "update usergroup set $set where id=? "
	_, err := db.Mconn.UpdateInterface(ug, gsql, ug.Id)
	return err
}

func (ug *UserGroup) Create() error {
	gsql := "insert into usergroup($key) values($value) "
	ids, err := db.Mconn.InsertInterfaceWithID(ug, gsql)
	if err != nil {
		return err
	}
	ug.Id = ids[0]
	return err
}

func GetUserGroupList(uid int64) ([]UserGroup, error) {
	ug := make([]UserGroup, 0)
	gsql := "select * from usergroup where uid=? or uid=? or json_contains(uids,json_array(?))"
	err := db.Mconn.Select(&ug, gsql, uid, goconfig.ReadInt64("adminid", 1), uid)
	return ug, err
}
