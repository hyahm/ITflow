package model

import (
	"itflow/db"

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
	Id  int64 `json:"id" db:"id,default"`
	PV  uint8 `json:"pv" db:"pv,default"`
	Rid int64 `json:"rid" db:"rid,default"`
}

// func (perm *Perm) Insert() error {
// 	var err error
// 	perm.Id, err = db.Mconn.Insert("insert into perm(rid, find, remove,revise, increase) values(?,?,?,?,?)",
// 		perm.Rid, perm.Find, perm.Remove, perm.Revise, perm.Increase)
// 	if err != nil {
// 		if err.(*mysql.MySQLError).Number == 1062 {
// 			return db.DuplicateErr
// 		}
// 	}
// 	return err
// }

func InsertManyPerm(perms []Perm) ([]int64, error) {
	return db.Mconn.InsertInterfaceWithID(perms, "insert into perm($key) values($value)")
}

func (perm *Perm) Update() error {
	_, err := db.Mconn.UpdateInterface(perm, "update perm set $set where id=?", perm.Id)
	return err
}

func DeletePerms(ids ...int64) error {
	_, err := db.Mconn.DeleteIn("delete from perm where id in (?)", ids)
	return err
}

func GetPermsionPageAndPVById(permids []int64) (map[string]uint8, error) {
	rows, err := db.Mconn.GetRowsIn("select name,pv from perm join roles on perm.rid=roles.id and perm.id in (?) ", permids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pagekey := make(map[string]uint8)
	for rows.Next() {
		var name string
		var pv uint8
		err = rows.Scan(&name, &pv)
		if err != nil {
			golog.Error(err)
			continue
		}
		pagekey[name] = pv
	}
	return pagekey, nil
}

func GetPermsionByIds(permids []int64) ([]string, error) {
	rows, err := db.Mconn.GetRowsIn("select name from perm join roles on perm.rid=roles.id and perm.id in (?) and pv>0 ", permids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	names := make([]string, 0)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Error(err)
			continue
		}
		names = append(names, name)
	}
	return names, nil
}
