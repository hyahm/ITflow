package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"strings"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

// 用户登录

type Login struct {
	Username string `json:"username" type:"string" need:"是" default:"" information:"用户名"`
	Password string `json:"password"  type:"string" need:"是" default:"" information:"密码"`
}

func (login *Login) Check() (*RespLogin, error) {
	resp := &RespLogin{}
	login.Username = strings.Trim(login.Username, " ")
	enpassword := encrypt.PwdEncrypt(login.Password, cache.Salt)
	getsql := "select id,nickname from user where email=? and password=? and disable=0"
	err := db.Mconn.GetOne(getsql, login.Username, enpassword).Scan(&resp.ID, &resp.UserName)
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, errors.New("账号或密码错误")
		}
		golog.Error(err)
		return resp, err
	}
	// 这里登录信息插入缓存表
	resp.Token = encrypt.Token(resp.UserName, cache.Salt)
	token := &db.Token{
		Token:    resp.Token,
		NickName: login.Username,
		Id:       resp.ID,
	}
	err = db.Table.Add(token, goconfig.ReadDuration("expiration", time.Minute*120))
	if err != nil {
		golog.Error(err)
		return resp, err
	}
	return resp, nil
}

type User struct {
	Id          int    `json:"id"`
	Createtime  int64  `json:"createtime"`
	Realname    string `json:"realname"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Disable     int    `json:"disable"`
	StatusGroup string `json:"statusgroup"`
	RoleGroup   string `json:"rolegroup"`
	Position    string `json:"position"`
}
type UserList struct {
	Userlist []*User `json:"userlist"`
	Code     int     `json:"code"`
}

type UserInfo struct {
	Roles    []string `json:"roles" type:"array" need:"否" default:"" information:"角色组"`
	Code     int      `json:"code" type:"string" need:"是" default:"" information:"状态码， 0为成功"`
	Avatar   string   `json:"avatar" type:"string" need:"否" default:"" information:"个人头像地址"`
	NickName string   `json:"name" type:"string" need:"否" default:"" information:"用户昵称"`
	Msg      string   `json:"msg,omitempty" type:"string" need:"否" default:"" information:"错误信息"`
	Realname string   `json:"realname,omitempty" type:"string" need:"否" default:"" information:"真实姓名"`
	Email    string   `json:"email,omitempty" type:"string" need:"否" default:"" information:"邮箱地址"`
	Uid      int64    `json:"uid,omitempty" type:"int" need:"否" default:"" information:"用户id"`
}

func (ui *UserInfo) GetUserInfo(uid int64) error {
	ui.Roles = make([]string, 0)
	err := db.Mconn.GetOne("select nickname, headimg from user where id=?", uid).Scan(&ui.NickName, &ui.Avatar)
	if err != nil {
		golog.Error(err)
		return err
	}
	// 管理员
	if uid == cache.SUPERID {
		ui.Roles = append(ui.Roles, "admin")
	} else {
		var permids string
		err = db.Mconn.GetOne("select permids from rolegroup where id=(select rid from user where id=?)", uid).Scan(&permids)
		if err != nil {
			golog.Error(err)
			return err
		}
		rows, err := db.Mconn.GetRowsIn("select r.name from perm as p join roles as r on p.id in (?) and p.find=true and p.rid=r.id",
			gomysql.InArgs(strings.Split(permids, ",")).ToInArgs())

		if err != nil {
			golog.Error(err)
			return err
		}
		for rows.Next() {
			role := new(string)
			err = rows.Scan(role)
			if err != nil {
				golog.Error(err)
				continue
			}
			ui.Roles = append(ui.Roles, *role)
		}
		rows.Close()
	}
	if len(ui.Roles) == 0 {
		ui.Roles = append(ui.Roles, "test")
	}
	return nil
}

func (ui *UserInfo) Update() error {
	sql := "select rid, headimg from user where nickname=?"
	var rid string
	err := db.Mconn.GetOne(sql, ui.NickName).Scan(&rid, &ui.Avatar)
	if err != nil {
		golog.Error(err)
		return err
	}

	// 管理员
	if ui.Uid == cache.SUPERID {
		ui.Roles = append(ui.Roles, ui.NickName)
	} else {
		getrole := "select name from roles where id in (select rolelist from rolegroup where id=?)"
		rows, err := db.Mconn.GetRows(getrole, rid)
		if err != nil {
			golog.Error(err)
			return err
		}
		for rows.Next() {
			role := new(string)
			err = rows.Scan(role)
			if err != nil {
				golog.Error(err)
				continue
			}
			ui.Roles = append(ui.Roles, *role)
		}
		rows.Close()
	}

	return nil
}

func (ui *UserInfo) Json() []byte {
	send, err := json.Marshal(ui)
	if err != nil {
		golog.Error(err)
	}
	return send
}
