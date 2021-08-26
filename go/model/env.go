package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Env struct {
	ID   int64  `json:"id" db:"id,default"`
	Name string `json:"name" db:"name"`
}

// 获取的就是表的所有字段
func GetAllEnv() ([]Env, error) {
	envs := make([]Env, 0)
	err := db.Mconn.Select(&envs, "select * from environment")
	return envs, err
}

func (env *Env) Create() error {
	ids, err := db.Mconn.InsertInterfaceWithID(env, "insert into environment($key) values($value)")
	if err != nil {
		golog.Error(err)
		return err
	}
	env.ID = ids[0]
	return nil
}

func (env *Env) Update() error {
	_, err := db.Mconn.UpdateInterface(env, "update environment set $set where id=?", env.ID)
	return err
}

func DeleteEnv(id interface{}) error {
	_, err := db.Mconn.Delete("delete from environment where id=?", id)
	return err
}
