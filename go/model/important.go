package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Important struct {
	Id   int64  `json:"id" db:"id,default"`
	Name string `json:"name" db:"name"`
}

// 获取的就是表的所有字段
func GetAllImportant() ([]Important, error) {
	importants := make([]Important, 0)
	result := db.Mconn.Select(&importants, "select * from importants")
	return importants, result.Err
}

func (important *Important) Create() error {
	result := db.Mconn.InsertInterfaceWithID(important, "insert into importants($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	important.Id = result.LastInsertId
	return nil
}

func (important *Important) Update() error {
	result := db.Mconn.UpdateInterface(important, "update importants set $set where id=?", important.Id)
	return result.Err
}

func DeleteImportant(id interface{}) error {
	result := db.Mconn.Delete("delete from importants where id=?", id)
	return result.Err
}
