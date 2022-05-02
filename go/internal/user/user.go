package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/model"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux/auth"
)

// 用户登录
type ResponseLogin struct {
	ID int64 `json:"id"`
}

type Login struct {
	Username string `json:"username" type:"string" need:"是" default:"" information:"用户名"`
	Password string `json:"password"  type:"string" need:"是" default:"" information:"密码"`
}

func (login *Login) Check() (string, int64, error) {
	rl := &ResponseLogin{}
	login.Username = strings.Trim(login.Username, " ")
	enpassword := encrypt.PwdEncrypt(login.Password, cache.Salt)
	getsql := ""
	if strings.Contains(login.Username, "@") {
		getsql = "select id from user where email=? and password=? and disable=0"
	} else {
		getsql = "select id from user where nickname=? and password=? and disable=0"
	}

	err := db.Mconn.GetOne(getsql, login.Username, enpassword).Scan(&rl.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", 0, errors.New("账号或密码错误")
		}
		golog.Error(err)
		return "", 0, err
	}
	token, err := auth.MakeJwt(cache.Salt, rl)
	return token, rl.ID, err
}

type User struct {
	Id         int     `json:"id" db:"id,default"`
	Createtime int64   `json:"createtime" db:"createtime,default"`
	Realname   string  `json:"realname" db:"realname,default"`
	Nickname   string  `json:"nickname" db:"nickname,default"`
	Email      string  `json:"email" db:"email,default"`
	Headmg     string  `json:"headimg" db:"headimg,default"`
	Disable    bool    `json:"disable" db:"disable"`
	JobId      int64   `json:"jid" db:"jid"`
	ShowStatus []int64 `json:"showstatus" db:"showstatus"`
	CreateUid  int64   `json:"createuid" db:"createuid"`
}
type UserList struct {
	Userlist []*User `json:"userlist"`
	Code     int     `json:"code"`
}

type UserInfo struct {
	Roles    []string `json:"roles" type:"array" need:"否" default:"" information:"角色组"`
	Avatar   string   `json:"avatar" type:"string" need:"否" default:"" information:"个人头像地址"`
	NickName string   `json:"name" type:"string" need:"否" default:"" information:"用户昵称"`
	// Realname string   `json:"realname,omitempty" type:"string" need:"否" default:"" information:"真实姓名"`
	// Email string `json:"email,omitempty" type:"string" need:"否" default:"" information:"邮箱地址"`
	Uid int64 `json:"uid,omitempty" type:"int" need:"否" default:"" information:"用户id"`
}

func (ui *UserInfo) GetUserInfo(uid int64) error {
	ui.Roles = make([]string, 0)
	var jid int64
	err := db.Mconn.GetOne("select nickname, headimg, jid from user where id=?", uid).Scan(&ui.NickName, &ui.Avatar, &jid)
	if err != nil {
		golog.Error(err)
		return err
	}
	// 管理员
	if uid == cache.SUPERID {
		ui.Roles = append(ui.Roles, "admin")
		return nil
	}

	permids, err := model.GetPermIdsByUid(uid)
	if err != nil {
		golog.Error(err)
		return err
	}
	if len(permids) == 0 {
		ui.Roles = append(ui.Roles, "test")
		return nil
	}
	ui.Roles, err = model.GetPermsionByIds(permids)
	return err
}

// func (ui *UserInfo) Update() error {
// 	sql := "select rid, headimg from user where nickname=?"
// 	var rid string
// 	err := db.Mconn.GetOne(sql, ui.NickName).Scan(&rid, &ui.Avatar)
// 	if err != nil {
// 		golog.Error(err)
// 		return err
// 	}

// 	// 管理员
// 	if ui.Uid == cache.SUPERID {
// 		ui.Roles = append(ui.Roles, ui.NickName)
// 	} else {
// 		getrole := "select name from roles where id in (select rolelist from rolegroup where id=?)"
// 		rows, err := db.Mconn.GetRows(getrole, rid)
// 		if err != nil {
// 			golog.Error(err)
// 			return err
// 		}
// 		for rows.Next() {
// 			role := new(string)
// 			err = rows.Scan(role)
// 			if err != nil {
// 				golog.Error(err)
// 				continue
// 			}
// 			ui.Roles = append(ui.Roles, *role)
// 		}
// 		rows.Close()
// 	}

// 	return nil
// }

func (ui *UserInfo) Json() []byte {
	send, err := json.Marshal(ui)
	if err != nil {
		golog.Error(err)
	}
	return send
}
