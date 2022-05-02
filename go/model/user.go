package model

import (
	"database/sql"
	"errors"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

type User struct {
	ID         int64   `json:"id" db:"id,default"`
	NickName   string  `json:"nickname" db:"nickname"`
	Password   string  `json:"password" db:"password"`
	Email      string  `json:"email" db:"email"`
	HeadImg    string  `json:"headimg" db:"headimg"`
	CreateTime int64   `json:"createtime" db:"createtime"`
	CreateUId  int64   `json:"createuid" db:"createuid"`
	RealName   string  `json:"realname" db:"realname"`
	ShowStatus []int64 `json:"showstatus" db:"showstatus"` // 查看的状态
	Disable    bool    `json:"disable" db:"disable"`       // 是否是垃圾箱
	Jobid      int64   `json:"jid" db:"jid"`               // 职位
}

func (user *User) UpdatePassword(old string) error {

	result := db.Mconn.Update("update user set password=? where password=? and id=?", user.Password, old, user.ID)

	return result.Err
}

func GetAllUsers(uid int64) ([]User, error) {
	us := make([]User, 0)
	if uid == cache.SUPERID {
		result := db.Mconn.Select(&us, "select * from user")
		return us, result.Err
	}
	return nil, errors.New("no permission")
}

func GetJobIdByUid(uid int64) (int64, error) {
	var jid int64
	err := db.Mconn.GetOne("select jid from user where id=?", uid).Scan(&jid)
	return jid, err
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

func GetShowStatus(uid int64) ([]int64, error) {
	user := User{}
	result := db.Mconn.Select(&user, "select showstatus from user where id=?", uid)
	if result.Err != nil {
		golog.Error(result.Err)
		return nil, result.Err
	}
	return user.ShowStatus, nil
}

func (user *User) Create() error {
	// user.HeadImg = goconfig.ReadString("defaulthead")
	result := db.Mconn.InsertInterfaceWithID(user, "insert into user($key) values($value)")
	if result.Err != nil {
		golog.Error(result.Err)
		return result.Err
	}
	user.ID = result.LastInsertId
	return nil
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
	basesql := "update user set $set where id=?"
	result := db.Mconn.UpdateInterface(user, basesql, user.ID)
	return result.Err
}

func GetUserKeyNameByProjectId(projectId int64) ([]KeyName, error) {
	// 获取用户ids
	ug := UserGroup{}
	result := db.Mconn.Select(&ug, "select uids from usergroup where id=( select ugid from project where id=?)", projectId)
	if result.Err != nil {
		golog.Error(result.Err)
		return nil, result.Err
	}
	rows, err := db.Mconn.GetRowsIn(" select id,realname from user where id in (?)", ug.Uids)
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
