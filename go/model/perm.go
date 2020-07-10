package model

import (
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

	// 更新的时候 rid 也是固定的， 不需要修改
	effect, err := db.Mconn.UpdateIn("update perm set find=?, remove=?,revise=?, increase=? where rid=? and id in (?)",
		perm.Find, perm.Remove, perm.Revise, perm.Increase, perm.Rid, strings.Split(ids, ","))
	golog.Info(effect)
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
