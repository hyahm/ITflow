package model

import (
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

type UserGroup struct {
	Id   int64   `json:"id" db:"id,default"`
	Name string  `json:"name" db:"name"`
	Uids []int64 `json:"uids" db:"uids"`
	Uid  int64   `json:"uid" db:"uid"`
}

func (ug *UserGroup) Delete(id interface{}) (err error) {

	if cache.SUPERID == ug.Uid {
		gsql := "delete from usergroup where id=? "
		_, err = db.Mconn.Delete(gsql, id)
	} else {
		gsql := "delete from  usergroup where id=? and uid=?"
		_, err = db.Mconn.Delete(gsql, id, ug.Uid)
	}
	return
}

func (ug *UserGroup) Update() (err error) {

	if cache.SUPERID == ug.Uid {
		gsql := "update usergroup set $set where id=? "
		_, err = db.Mconn.UpdateInterface(ug, gsql, ug.Id)
	} else {
		gsql := "update usergroup set $set where id=? and uid=?"
		_, err = db.Mconn.UpdateInterface(ug, gsql, ug.Id, ug.Uid)
	}
	return
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

func GetUserGroupIds(uid int64) ([]int64, error) {

	gsql := "select id from usergroup where uid=? or uid=? or json_contains(uids,json_array(?))"
	rows, err := db.Mconn.GetRows(gsql, uid, goconfig.ReadInt64("adminid", 1), uid)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0)
	defer rows.Close()
	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (ug *UserGroup) GetUserIds() error {
	return db.Mconn.Select(&ug, "select uids from usergroup where id=?", ug.Id)
}

func GetUserGroupKeyNameByUid() ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,name from usergroup")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kns := make([]KeyName, 0)
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		kns = append(kns, kn)
	}
	return kns, nil
}
