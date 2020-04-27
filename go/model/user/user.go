package user

import (
	"database/sql"
	"itflow/bug/bugconfig"
	"itflow/db"
	"itflow/gaencrypt"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

// 用户登录

type Login struct {
	Username string `json:"username" type:"string" need:"是" default:"" information:"用户名"`
	Password string `json:"password"  type:"string" need:"是" default:"" information:"密码"`
}

func (login *Login) Check(resp *RespLogin) error {
	login.Username = strings.Trim(login.Username, " ")

	enpassword := gaencrypt.PwdEncrypt(login.Password, bugconfig.Salt)
	getsql := "select nickname from user where email=? and password=? and disable=0"

	row, err := db.Mconn.GetOne(getsql, login.Username, enpassword)
	if err != nil {
		golog.Error(err)
		return err
	}
	err = row.Scan(&login.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			golog.Error("username or password error")
			return err
		}
		golog.Error(err)
		return err
	}

	resp.Token = gaencrypt.Token(login.Username, bugconfig.Salt)
	_, err = db.RSconn.Set(resp.Token, login.Username,
		time.Duration(goconfig.ReadInt("expiration", 120))*time.Minute)
	if err != nil {
		golog.Error(err)
		return err
	}
	resp.UserName = login.Username
	return nil
}

type UserInfo struct {
	Roles  []string `json:"roles" type:"array" need:"是" default:"" information:"角色组"`
	Code   int      `json:"code" type:"string" need:"是" default:"" information:"状态码"`
	Avatar string   `json:"avatar" type:"string" need:"是" default:"" information:"个人头像地址"`
	Name   string   `json:"nickname" type:"string" need:"是" default:"" information:"用户昵称"`
}

func (ui *UserInfo) GetUserInfo() error {
	sql := "select rid, headimg from user where nickname=?"
	var rid string
	row, err := db.Mconn.GetOne(sql, ui.Name)
	if err != nil {
		golog.Error(err)
		return err
	}
	err = row.Scan(&rid, &ui.Avatar)
	if err != nil {
		golog.Error(err)
		return err
	}
	// 管理员
	if bugconfig.CacheNickNameUid[ui.Name] == bugconfig.SUPERID {
		ui.Roles = append(ui.Roles, goconfig.ReadString("adminuser", "admin"))
	} else {
		var rl string
		getrole := "select rolelist from rolegroup where id=?"
		row, err := db.Mconn.GetOne(getrole, rid)
		if err != nil {
			golog.Error(err)
			return err
		}
		err = row.Scan(&rl)
		if err != nil {
			golog.Error(err)
			return err
		}
		for _, v := range strings.Split(rl, ",") {
			id, _ := strconv.Atoi(v)
			ui.Roles = append(ui.Roles, bugconfig.CacheRidRole[int64(id)])
		}
	}
	return nil
}
