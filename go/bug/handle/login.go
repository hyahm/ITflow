package handle

import (
	"database/sql"
	"encoding/json"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/db"
	"itflow/gaencrypt"
	"itflow/model/response"
	"itflow/model/user"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Login(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	login := xmux.GetData(r).Data.(*user.Login)
	login.Username = strings.Trim(login.Username, " ")

	enpassword := gaencrypt.PwdEncrypt(login.Password, bugconfig.Salt)

	getsql := "select nickname from user where email=? and password=? and disable=0"

	row, err := db.Mconn.GetOne(getsql, login.Username, enpassword)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&login.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			golog.Error("username or password error")
			w.Write(errorcode.Error("username or password error"))
			return
		}
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	resp := user.RespLogin{}

	resp.Token = gaencrypt.Token(login.Username, bugconfig.Salt)
	_, err = db.RSconn.Set(resp.Token, login.Username,
		time.Duration(goconfig.ReadInt("expiration", 120))*time.Minute)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "login",
	}
	err = il.Login("nickname: %s has login", login.Username)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	resp.UserName = login.Username

	send, _ := json.Marshal(resp)
	w.Write(send)
	return

}

func Loginout(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	token := r.FormValue("token")
	nickname, err := db.RSconn.Get(token)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	il := buglog.AddLog{
		Ip: strings.Split(r.RemoteAddr, ":")[0],
	}
	err = il.Login("nickname: %s has logout", nickname)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

}

type userInfo struct {
	Roles  []string `json:"roles"`
	Code   int      `json:"code"`
	Avatar string   `json:"avatar"`
	Name   string   `json:"name"`
}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname, err := logtokenmysql(r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	userinfo := &userInfo{}

	sql := "select rid,level,headimg,nickname from user where nickname=?"
	var rid string
	var level int
	row, err := db.Mconn.GetOne(sql, nickname)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&rid, &level, &userinfo.Avatar, &userinfo.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		userinfo.Roles = append(userinfo.Roles, "admin")
	} else {
		var rl string
		getrole := "select rolelist from rolegroup where id=?"
		row, err := db.Mconn.GetOne(getrole, rid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		err = row.Scan(&rl)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for _, v := range strings.Split(rl, ",") {
			id, _ := strconv.Atoi(v)
			userinfo.Roles = append(userinfo.Roles, bugconfig.CacheRidRole[int64(id)])
		}
	}

	send, _ := json.Marshal(userinfo)
	w.Write(send)

}
