package model

import (
	"database/sql"
	"errors"
	"itflow/cache"
	"itflow/db"
	"time"

	"github.com/hyahm/goconfig"
)

type User struct {
	ID         int64
	NickName   string
	Password   string
	Email      string
	HeadImg    string
	CreateTime int64
	CreateId   int64
	RealName   string
	ShowStatus cache.StoreLevelId
	Disable    bool  // 是否是垃圾箱
	BugGroupId int64 // 可以查看的bug状态
	// Level      int   // 是否是管理员， 无效
	Roleid int64 // 角色权限
	Jobid  int64 // 职位
}

func (user *User) Create() error {
	user.HeadImg = goconfig.ReadString("defaulthead")

	var err error
	createusersql := "insert into user(nickname,password,email,headimg,createtime,createuid,realname,showstatus,disable,bugsid,rid,jid) values(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	user.ID, err = db.Mconn.Insert(createusersql,
		user.NickName, user.Password, user.Email,
		user.HeadImg, time.Now().Unix(), user.CreateId,
		user.RealName, user.ShowStatus, false,
		user.BugGroupId, user.Roleid, user.Jobid,
	)

	return err
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
	_, err := db.Mconn.Update("update user set password=? where rid=?", password, goconfig.ReadInt("adminid", 1))
	return err

}

func (user *User) UpdateShowStatus(showstatus, uid interface{}) error {
	basesql := "update user set showstatus=? where id=?"
	_, err := db.Mconn.Update(basesql, showstatus, uid)
	if err != nil {
		return err
	}
	return nil
}

func NewUserById(id interface{}) (*User, error) {
	user := &User{}
	err := db.Mconn.GetOne("select id, nickname,email,headimg,createtime,createuid,realname,showstatus,disable,bugsid,rid,jid from user where id=?", id).
		Scan(&user.ID, &user.NickName, &user.Email, &user.HeadImg, &user.CreateTime, &user.CreateId, &user.RealName, &user.ShowStatus,
			&user.Disable, &user.BugGroupId, &user.Roleid, &user.Jobid)
	return user, err
}
