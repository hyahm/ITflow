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
	result := db.Mconn.Select(&envs, "select * from environment")
	return envs, result.Err
}

func (env *Env) Create() error {
	result := db.Mconn.InsertInterfaceWithID(env, "insert into environment($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	env.ID = result.LastInsertId
	return nil
}

func (env *Env) Update() error {
	result := db.Mconn.UpdateInterface(env, "update environment set $set where id=?", env.ID)
	return result.Err
}

func DeleteEnv(id interface{}) error {
	result := db.Mconn.Delete("delete from environment where id=?", id)
	return result.Err
}
