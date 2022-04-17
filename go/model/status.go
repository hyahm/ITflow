package model

import (
	"database/sql"
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
)

type Status struct {
	ID   int64  `json:"id" db:"id,default"`
	Name string `json:"name" db:"name"`
}

func (status *Status) Names() ([]string, error) {
	names := make([]string, 0)
	rows, err := db.Mconn.GetRows("select name from status")
	if err != nil {
		if err == sql.ErrNoRows {
			return names, nil
		}
		return names, err
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		names = append(names, name)
	}
	rows.Close()
	return names, nil
}

func GetStatusList() ([]*Status, error) {
	ss := make([]*Status, 0)
	result := db.Mconn.Select(&ss, "select * from status")

	return ss, result.Err
}

func GetStatusIDsByUid(uid int64) ([]*Status, error) {

	if uid == cache.SUPERID {
		return GetStatusList()
	} else {
		// 获取statusids 通过uid
		sids, err := getStatusIDsByUid(uid)
		if err != nil {
			golog.Error(err)
			return nil, err
		}
		ss := make([]*Status, 0)
		result := db.Mconn.SelectIn(&ss, "select * from status where id in (?)", sids)

		return ss, result.Err

	}
}

func GetMyStatusList(id interface{}) ([]string, error) {
	var sids string
	err := db.Mconn.GetOne("select showstatus from user where id=?", id).Scan(&sids)
	if err != nil {
		return nil, err
	}
	return strings.Split(sids, ","), nil
}

// 获取的就是表的所有字段
func GetAllStatus() ([]Status, error) {
	statuss := make([]Status, 0)
	result := db.Mconn.Select(&statuss, "select * from status")
	return statuss, result.Err
}

func (status *Status) Create() error {
	result := db.Mconn.InsertInterfaceWithID(status, "insert into status($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	status.ID = result.LastInsertId
	return nil
}

func (status *Status) Update() error {
	result := db.Mconn.UpdateInterface(status, "update status set $set where id=?", status.ID)
	return result.Err
}

func DeleteStatus(id interface{}) error {
	result := db.Mconn.Delete("delete from status where id=?", id)
	return result.Err
}
