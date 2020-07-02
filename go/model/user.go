package model

import (
	"database/sql"
	"errors"
	"itflow/db"

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
	ShowStatus string
	Disable    bool
	BugId      int64
	Level      int
	Roleid     int64
	Jobid      int64
}

func (user *User) CheckHaveAdminUser() error {
	// 返回nil 表示存在admin账号
	var count int
	err := db.Mconn.GetOne("select count(id) from user where rid=?", goconfig.ReadInt("adminid", 0)).Scan(&count)
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
	_, err := db.Mconn.Update("update user set password=? where rid=?", password, goconfig.ReadInt("adminid", 0))
	return err

}
