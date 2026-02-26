package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type Env struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}

func (Env) TableName() string {
	return "environment"
}

// 获取的就是表的所有字段
func (e *Env) GetAllEnv() ([]Env, error) {
	envs := make([]Env, 0)
	err := db.Gorm.Table(e.TableName()).Order("id asc").Find(&envs).Error
	return envs, err
}

func (env *Env) Create() error {
	result := db.Mconn.InsertInterfaceWithID(env, "insert into environment($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	env.Id = result.LastInsertId
	return nil
}

func (env *Env) Update() error {
	result := db.Mconn.UpdateInterface(env, "update environment set $set where id=?", env.Id)
	return result.Err
}

func DeleteEnv(id interface{}) error {
	result := db.Mconn.Delete("delete from environment where id=?", id)
	return result.Err
}
