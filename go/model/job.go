package model

import (
	"database/sql"
	"errors"
	"itflow/cache"
	"itflow/db"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gosql"
)

type Position struct {
	Id      int64     `json:"id" gorm:"primaryKey"`
	Name    string    `json:"name" gorm:"column:name"`
	Level   int       `json:"level" gorm:"column:level"`     // 1 是管理者， 0 是普通员工
	Hypo    int64     `json:"hypo" gorm:"column:hypo"`       //  上级id
	RoleId  int64     `json:"role_id" gorm:"column:role_id"` // 状态组0
	Uid     int64     `json:"uid" gorm:"column:uid"`         // 创建者
	Created time.Time `json:"created" gorm:"column:created" `
	Updated time.Time `json:"updated" gorm:"column:updated" `
}

func (Position) TableName() string {
	return "position"
}

func (p *Position) Create() error {
	p.Created = time.Now()
	p.Updated = time.Now()
	return db.Gorm.Create(p).Error
}

// 获取所有管理员
func (p *Position) GetManager() ([]KeyName, error) {
	ps := make([]KeyName, 0)
	err := db.Gorm.Table(p.TableName()).Select("id", "name").Where("level=1").Find(&ps).Error
	return ps, err
}

func DeleteJob(id, uid interface{}) (err error) {
	var result gosql.Result
	if uid == cache.SUPERID {
		result = db.Mconn.Delete("delete from jobs where id=?", id)
	} else {
		result = db.Mconn.Delete("delete from jobs where id=? and hypo=?", id, uid)
	}

	return result.Err
}

func (p *Position) Update() error {
	result := db.Mconn.UpdateInterface(p, "update jobs set $set where id=?", p.Id)
	return result.Err
}

func (p *Position) GetAllPositions() ([]Position, error) {
	jobs := make([]Position, 0)
	err := db.Gorm.Table(p.TableName()).Order("id desc").Find(&jobs).Error
	return jobs, err
}

func (p *Position) GetJobIdsByJobId() ([]int64, error) {
	// 通过jid 来获取 能管理的 职位 的id

	if p.Id <= 0 {
		return []int64{}, errors.New("id is zero")
	}
	var ids []int64
	err := db.Gorm.Table(p.TableName()).Where("hypo=( select hypo from position where id=?)", p.Id).Select("id").Find(&ids).Error
	return ids, err
}

type Jobs struct {
	Positions []*Position `json:"positions"`
	Code      int         `json:"code"`
}

func GetJobKeyNameByUid(uid int64) ([]KeyName, error) {
	var rows *sql.Rows
	var err error
	if uid == cache.SUPERID {
		rows, err = db.Mconn.GetRows("select id,name from jobs")
	} else {
		rows, err = db.Mconn.GetRows("select id,name from jobs where hypo=?", uid)
	}
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

func GetManagerKeyName(uid int64) ([]KeyName, error) {
	var err error
	var rows *sql.Rows
	if uid == goconfig.ReadInt64("adminid") {
		rows, err = db.Mconn.GetRows("select id,name from jobs where level=1")
	} else {
		rows, err = db.Mconn.GetRows("select id,name from jobs where level=1 and uid=?", uid)
	}

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
