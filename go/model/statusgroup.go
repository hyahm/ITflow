package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type StatusGroup struct {
	ID        int64   `json:"id" db:"id,default"`
	Name      string  `json:"name" db:"name"`
	StatusIDs []int64 `json:"sids" db:"sids"`
}

func getStatusIDsByUid(uid int64) ([]int64, error) {
	sg := StatusGroup{}
	sql := "select * from statusgroup where id=(select sgid from jobs where id=(select jid from user where id=?))"
	result := db.Mconn.Select(&sg, sql, uid)

	return sg.StatusIDs, result.Err
}

func GetStatusGroupKeyName() ([]KeyName, error) {
	s := "select id,name from statusgroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)

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

func (sg *StatusGroup) Insert() error {
	result := db.Mconn.InsertInterfaceWithID(sg, "insert into statusgroup($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	sg.ID = result.LastInsertId
	return nil
}

func (sg *StatusGroup) Update() error {
	result := db.Mconn.UpdateInterface(sg, "update statusgroup set $set where id = ?", sg.ID)
	return result.Err

}

func GetAllStatusGroup() ([]StatusGroup, error) {
	sgs := make([]StatusGroup, 0)
	result := db.Mconn.Select(&sgs, "select * from statusgroup")
	return sgs, result.Err
}

func DeleteStatusGroup(id interface{}) error {
	result := db.Mconn.Delete("delete from statusgroup where id=?", id)
	return result.Err
}
