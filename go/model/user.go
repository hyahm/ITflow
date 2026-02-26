package model

import (
	"database/sql"
	"errors"
	"itflow/cache"
	"itflow/db"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

var ErrIdNotFound = errors.New("id not found")

type User struct {
	Id         int64     `json:"id" gorm:"primaryKey"`
	NickName   string    `json:"nickname" gorm:"column:nickname"`
	Password   string    `json:"password" gorm:"column:password"`
	Email      string    `json:"email" gorm:"column:email"`
	HeadImg    string    `json:"headimg" gorm:"column:headimg"`
	Created    time.Time `json:"created" gorm:"column:created"`
	Updated    time.Time `json:"updated" gorm:"column:updated"`
	CreateId   int64     `json:"create_id" gorm:"column:create_id"`
	RealName   string    `json:"realname" gorm:"column:realname"`
	Disable    bool      `json:"disable" gorm:"column:disable"`         // 是否是垃圾箱
	PositionId int64     `json:"position_id" gorm:"column:position_id"` // 职位
}

func (User) TableName() string {
	return "user"
}

func (user *User) GetUserKeyName(uid int64) ([]KeyName, error) {
	kns := make([]KeyName, 0)
	err := db.Gorm.Table(user.TableName()).Select("id as id", "nickname as name").Find(&kns).Error
	golog.Info(kns)
	return kns, err
}

func (user *User) UpdatePassword(old string) error {
	return db.Gorm.Table(user.TableName()).Where("password=? and id=?", old, user.Id).Select("password").Updates(user).Error
}

func GetAllUsers(uid int64) ([]User, error) {
	us := make([]User, 0)
	err := db.Gorm.Find(&us).Error
	return us, err
}

func (user *User) GetKeyNameByUids(uids []int64) ([]KeyName, error) {
	userinfo := make([]KeyName, 0)
	err := db.Gorm.Table(user.TableName()).Where("id in ?", uids).Select("id", "realname as name").Find(&userinfo).Error
	return userinfo, err
}
func (user *User) GetJobIdByUid(uid int64) (int64, error) {
	var positionId int64
	err := db.Gorm.Table(user.TableName()).Select("position_id").Where("id=?", uid).Scan(&positionId).Error
	return positionId, err
}
func DeleteUser(id interface{}) error {
	result := db.Mconn.Delete("delete from user where id=? ", id)
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	if result.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// 获取所有用户信息
func GetUsers(jobs []int64) ([]User, error) {
	users := make([]User, 0)
	result := db.Mconn.SelectIn(&users, "select * from user where jid in (?)", jobs)
	return users, result.Err

}

func (user *User) Create() error {
	// user.HeadImg = goconfig.ReadString("defaulthead")
	user.Created = time.Now()
	user.Updated = time.Now()
	return db.Gorm.Table(user.TableName()).Create(user).Error
}

func (user *User) CheckHaveAdminUser() error {
	// 返回nil 表示存在admin账号
	var count int
	err := db.Mconn.GetOne("select count(id) from user where rid=?", goconfig.ReadInt("adminid", 1)).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows || count != 1 {
			return errors.New("有且只能有一个admin账户 \n")
		}

		return err
	}
	return nil
}

func (user *User) UpdateAdminPassword(password string) error {
	// 修改密码
	result := db.Mconn.Update("update user set password=? where id=?", password, cache.SUPERID)
	return result.Err

}

func (user *User) Update() error {
	if user.Id <= 0 {
		return ErrIdNotFound
	}
	return db.Gorm.Table(user.TableName()).Where("id=?", user.Id).Update("headimg", user.HeadImg).Error
	// result := db.Mconn.UpdateInterface(user, basesql, user.Id)
	// return result.Err
}

func (user *User) GetUserKeyNameByProjectId(projectId int64) ([]KeyName, error) {
	// 获取用户ids
	uids := make([]int64, 0)
	err := db.Gorm.Table(user.TableName()).Where("id=?", projectId).Find(&uids).Error
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	kns := make([]KeyName, 0)
	err = db.Gorm.Table(user.TableName()).Select("id as value", "nickname as label").Where("id in ?", uids).Find(&kns).Error
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	return kns, nil
}

func GetAllUserKeyName() ([]KeyName, error) {
	// 获取用户ids

	rows, err := db.Mconn.GetRowsIn(" select id,realname from user")
	if err != nil {
		golog.Error(err)
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
