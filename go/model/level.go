package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Level struct {
	Id   int64  `json:"id" db:"id,default"`
	Name string `json:"name" db:"name"`
}

func GetLevelKeyNameByUid() ([]KeyName, error) {
	rows, err := db.Mconn.GetRows("select id,name from level")
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

func (level *Level) Create() error {
	result := db.Mconn.InsertInterfaceWithID(level, "insert into level($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	level.Id = result.LastInsertId
	return nil
}

func (level *Level) Update() error {
	result := db.Mconn.UpdateInterface(level, "update level set $set where id=?", level.Id)
	return result.Err
}

func DeleteLevel(id interface{}) error {
	result := db.Mconn.Delete("delete from level where id=?", id)
	return result.Err
}
