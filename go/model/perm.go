package model

import (
	"errors"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
)

// `id` int(11) NOT NULL AUTO_INCREMENT,
//   `rgid` int(11) NOT NULL DEFAULT '0',
//   `find` tinyint(1) NOT NULL DEFAULT '0',
//   `remove` tinyint(1) NOT NULL DEFAULT '0',
//   `revise` tinyint(1) NOT NULL DEFAULT '0',
//   `increase` tinyint(1) NOT NULL DEFAULT '0',
//   `rid` bigint(20) NOT NULL DEFAULT '0',
type Perm struct {
	Id       int64
	Find     bool
	Remove   bool
	Revise   bool
	Increase bool
	Rid      int64
}

func (perm *Perm) Insert() error {
	var err error
	perm.Id, err = db.Mconn.Insert("insert into perm(rid, find, remove,revise, increase) values(?,?,?,?,?)",
		perm.Rid, perm.Find, perm.Remove, perm.Revise, perm.Increase)
	return err
}

func (perm *Perm) Update(ids string) error {
	db.Mconn.OpenDebug()
	golog.Debug(ids)
	var err error
	if strings.Index(ids, ",") > 0 {
		if strings.Count(ids, ",") == 1 {
			_, err = db.Mconn.Update("update perm set find=?, remove=?,revise=?, increase=? where rid=? and id=?",
				perm.Find, perm.Remove, perm.Revise, perm.Increase, perm.Rid, ids)
		} else {
			_, err = db.Mconn.Update("update perm set find=?, remove=?,revise=?, increase=? where rid=? and id in (?)",
				perm.Find, perm.Remove, perm.Revise, perm.Increase, perm.Rid, ids)
		}
	} else {
		return errors.New("data errror")
	}
	// 更新的时候 rid 也是固定的， 不需要修改

	golog.Info(db.Mconn.GetSql())
	return err
}

func NewPerm(id interface{}) (*Perm, error) {
	var err error
	perm := &Perm{}
	err = db.Mconn.GetOne("select id,rid, find, remove,revise, increase from perm where id=?", id).Scan(
		&perm.Id, &perm.Rid, &perm.Find, &perm.Remove, &perm.Revise, &perm.Increase)
	return perm, err
}

func (perm *Perm) Delete() error {
	var err error
	_, err = db.Mconn.Delete("delete from perm where id=?", perm.Id)
	return err
}
