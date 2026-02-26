package model

import (
	"errors"
	"itflow/db"
)

type Role struct {
	Id   int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}

func (Role) TableName() string {
	return "role"
}

func (r *Role) List() ([]Role, error) {
	roles := make([]Role, 0)
	err := db.Gorm.Table(r.TableName()).Find(&roles).Error
	return roles, err
}

func (r *Role) Create() error {
	return db.Gorm.Table(r.TableName()).Create(&r).Error
}

func (r *Role) Delete() error {
	if r.Id <= 0 {
		return errors.New("id not found")
	}
	return db.Gorm.Table(r.TableName()).Where("id=?", r.Id).Delete(r).Error
}
