package model

import (
	"itflow/cache"
	"itflow/db"
)

type UserGroup struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Ids  string `json:"ids"`
	Uid  int64
}

func (ug *UserGroup) Update(uid int64) error {
	if uid == cache.SUPERID {
		gsql := "update usergroup set name=?,ids=? where id=? "
		_, err := db.Mconn.Update(gsql, ug.Name, ug.Ids, ug.Id)
		if err != nil {
			return err
		}
	} else {
		gsql := "update usergroup set name=?,ids=? where id=? and uid=?"
		_, err := db.Mconn.Update(gsql, ug.Name, ug.Ids, ug.Id, uid)
		if err != nil {
			return err
		}
	}
	return nil
}

// type Send_groups struct {
// 	GroupList []*Get_groups `json:"grouplist" type:"array" need:"是" information:"返回的grouplist"`
// 	Code      int           `json:"code" type:"int" need:"是" information:"错误码"`
// }
