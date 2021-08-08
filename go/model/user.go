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

func GetJobIdByUid(uid int64) (int64, error) {
	var jid int64
	err := db.Mconn.GetOne("select jid from user where id=?", uid).Scan(&jid)
	return jid, err
}

// 获取所有用户信息
func GetUsers(jobs []int64) ([]User, error) {
	users := make([]User, 0)
	err := db.Mconn.SelectIn(&users, "select * from user where jid in (?)", jobs)
	return users, err
}

func GetShowStatus(uid int64) ([]int64, error) {
	user := User{}
	err := db.Mconn.Select(&user, "select showstatus from user where id=?", uid)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	return user.ShowStatus, nil
}

func (user *User) Create() error {
	// user.HeadImg = goconfig.ReadString("defaulthead")

	// var err error
	// createusersql := "insert into user(nickname,password,email,headimg,createtime,createuid,realname,showstatus,disable,bugsid,rid,jid) values(?,?,?,?,?,?,?,?,?,?,?,?)"
	// user.ID, err = db.Mconn.Insert(createusersql,
	// 	user.NickName, user.Password, user.Email,
	// 	user.HeadImg, time.Now().Unix(), user.CreateId,
	// 	user.RealName, user.ShowStatus, false,
	// 	user.BugGroupId, user.Roleid, user.Jobid,
	// )

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
	_, err := db.Mconn.Update("update user set password=? where id=?", password, cache.SUPERID)
	return err

}

func (user *User) Update() error {
	basesql := "update user set $set where id=?"
	_, err := db.Mconn.UpdateInterface(user, basesql, user.ID)
	return err
}

func GetUserKeyNameByProjectId(projectId int64) ([]KeyName, error) {
	// 获取用户ids
	ug := UserGroup{}
	err := db.Mconn.Select(&ug, "select uids from usergroup where id=( select ugid from project where id=?)", projectId)
	if err != nil {
		golog.Error(err)
		return nil, err
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
