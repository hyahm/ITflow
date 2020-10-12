package model

import (
	"itflow/db"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
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
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			return db.DuplicateErr
		}
	}

	return err
}

func (perm *Perm) Update(ids string) error {
	var err error
	db.Mconn.OpenDebug()
	_, err = db.Mconn.UpdateIn("update perm set find=?, remove=?,revise=?, increase=? where rid=? and id in (?)",
		perm.Find, perm.Remove, perm.Revise, perm.Increase, perm.Rid,
		(gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
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
