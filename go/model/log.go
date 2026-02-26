package model

import (
	"itflow/db"
	"time"
)

type Log struct {
	Id         int64     `json:"id" `
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
	Classify   string    `json:"classify" `
	Ip         string    `json:"ip" `
	Uid        int64     `json:"uid"`
	Action     string    `json:"action" `
}

func (log *Log) TableName() string {
	return "log"
}

func (log *Log) Insert() error {
	log.CreateTime = time.Now()
	return db.Gorm.Table(log.TableName()).Create(log).Error
}
