package user

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/internal/response"
	"itflow/model"
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

func (login *Login) Check(resp *RespLogin) []byte {
	login.Username = strings.Trim(login.Username, " ")
	errorcode := &response.Response{}
	enpassword := encrypt.PwdEncrypt(login.Password, cache.Salt)
	golog.Info(enpassword)
	getsql := "select id,nickname from user where email=? and password=? and disable=0"
	var id int64
	err := db.Mconn.GetOne(getsql, login.Username, enpassword).Scan(&id, &login.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return errorcode.LoginFailed()
		}
		golog.Error(err)
		return errorcode.ConnectMysqlFail()
	}
	golog.Info(id)
	resp.Token = encrypt.Token(login.Username, cache.Salt)
	token := &db.Token{
		Token:    resp.Token,
		NickName: login.Username,
		Id:       id,
	}
	golog.Info(token)
	// err = db.Table.Add(token, time.Second*20)
	err = db.Table.Add(token, time.Duration(goconfig.ReadInt("expiration", 120))*time.Minute)
	if err != nil {
		golog.Error(err)
		return []byte(err.Error())
	}
	golog.Infof("login seccuss, user: %s, token: %s", login.Username, resp.Token)
	resp.UserName = login.Username
	return nil
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
	Msg      string   `json:"message,omitempty" type:"string" need:"否" default:"" information:"错误信息"`
	Realname string   `json:"realname,omitempty" type:"string" need:"否" default:"" information:"真实姓名"`
	Email    string   `json:"email,omitempty" type:"string" need:"否" default:"" information:"邮箱地址"`
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
		var pids string
		permids := "select permids from rolegroup where id=(select rid from user where id=?)"
		err := db.Mconn.GetOne(permids, uid).Scan(&pids)
		if err != nil {
			golog.Error(err)
			return err
		}

		for _, v := range strings.Split(pids, ",") {
			perm, err := model.NewPerm(v)
			if err != nil {
				golog.Error(err)
				return err
			}
			if perm.Find {
				ui.Roles = append(ui.Roles, cache.CacheRidRole[perm.Rid])
			}

		}
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
	if cache.CacheNickNameUid[ui.NickName] == cache.SUPERID {
		ui.Roles = append(ui.Roles, ui.NickName)
	} else {
		var rl string
		getrole := "select rolelist from rolegroup where id=?"
		err := db.Mconn.GetOne(getrole, rid).Scan(&rl)
		if err != nil {
			golog.Error(err)
			return err
		}

		for _, v := range strings.Split(rl, ",") {
			id, _ := strconv.Atoi(v)
			ui.Roles = append(ui.Roles, cache.CacheRidRole[int64(id)])
		}
	}

	return nil
}

func (ui *UserInfo) Json() []byte {
	send, _ := json.Marshal(ui)
	return send
}
