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
	err := db.Mconn.Select(&importants, "select * from importants")
	return importants, err
}

func (important *Important) Create() error {
	ids, err := db.Mconn.InsertInterfaceWithID(important, "insert into importants($key) values($value)")
	if err != nil {
		golog.Error(err)
		return err
	}
	important.Id = ids[0]
	return nil
}

func (important *Important) Update() error {
	_, err := db.Mconn.UpdateInterface(important, "update importants set $set where id=?", important.Id)
	return err
}

func DeleteImportant(id interface{}) error {
	_, err := db.Mconn.Delete("delete from importants where id=?", id)
	return err
}
