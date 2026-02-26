package model

import (
	"itflow/db"

	"github.com/hyahm/golog"
)

type UserStatus struct {
	UID    int64 `gorm:"column:uid" json:"uid"`
	Status int   `gorm:"column:status" json:"status"`
}

func (UserStatus) TableName() string {
	return "user_status"
}

func (us *UserStatus) GetShowStatus(uid int64) ([]int, error) {
	status := make([]int, 0)
	err := db.Gorm.Table(us.TableName()).Where("id=?", uid).Find(&status).Error
	if err != nil {
		golog.Error(err)
		return status, err
	}
	return status, nil
}
