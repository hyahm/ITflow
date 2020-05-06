package user

import (
	"database/sql"
	"encoding/json"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/gaencrypt"
	"itflow/network/response"
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
	enpassword := gaencrypt.PwdEncrypt(login.Password, bugconfig.Salt)
	getsql := "select nickname from user where email=? and password=? and disable=0"

	row, err := db.Mconn.GetOne(getsql, login.Username, enpassword)
	if err != nil {
		golog.Error(err)
		return errorcode.ConnectMysqlFail()
	}
	err = row.Scan(&login.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return errorcode.LoginFailed()
		}
		golog.Error(err)
		return errorcode.ConnectMysqlFail()
	}

	resp.Token = gaencrypt.Token(login.Username, bugconfig.Salt)
	_, err = db.RSconn.Set(resp.Token, login.Username,
		time.Duration(goconfig.ReadInt("expiration", 120))*time.Minute)
	if err != nil {
		golog.Error(err)
		return errorcode.ConnectRedisFail()
	}
	golog.Infof("login seccuss, user: %s, token: %s", login.Username, resp.Token)
	resp.UserName = login.Username
	return nil
}

type UserInfo struct {
	Roles    []string `json:"roles" type:"array" need:"否" default:"" information:"角色组"`
	Code     int      `json:"code" type:"string" need:"是" default:"" information:"状态码， 0为成功"`
	Avatar   string   `json:"avatar" type:"string" need:"否" default:"" information:"个人头像地址"`
	NickName string   `json:"nickname" type:"string" need:"否" default:"" information:"用户昵称"`
	Msg      string   `json:"msg,omitempty" type:"string" need:"否" default:"" information:"错误信息"`
	Realname string   `json:"realname,omitempty" type:"string" need:"否" default:"" information:"真实姓名"`
	Email    string   `json:"email,omitempty" type:"string" need:"否" default:"" information:"邮箱地址"`
}

func (ui *UserInfo) GetUserInfo() error {
	sql := "select rid, headimg from user where nickname=?"
	var rid string
	row, err := db.Mconn.GetOne(sql, ui.NickName)
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
	if bugconfig.CacheNickNameUid[ui.NickName] == bugconfig.SUPERID {
		ui.Roles = append(ui.Roles, ui.NickName)
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

func (ui *UserInfo) Json() []byte {
	send, _ := json.Marshal(ui)
	return send
}
