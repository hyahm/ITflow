package handle

import (
	"database/sql"
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/gadb"
	"itflow/gaencrypt"
	"net/http"
	"strconv"
	"strings"
)

type resLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Code     int    `json:"code"`
	Avatar   string `json:"avatar"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	errorcode := &errorstruct{}
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	login := &resLogin{}

	err = json.Unmarshal(s, login)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 解密
	//username, err := gaencrypt.RsaDecrypt(login.Username, bugconfig.PrivateKey, true)
	//if err != nil {
	//	golog.Error(err.Error())
	//	w.Write(errorcode.ErrorRsa())
	//	return
	//}
	//
	//tmp, err := gaencrypt.RsaDecrypt(login.Password, bugconfig.PrivateKey, true)
	//if err != nil {
	//
	//	golog.Error(err.Error())
	//	w.Write(errorcode.ErrorRsa())
	//	return
	//}
	login.Username = strings.Trim(login.Username, " ")

	login.Token = gaencrypt.Token(login.Username, bugconfig.Salt)
	enpassword := gaencrypt.PwdEncrypt(login.Password, bugconfig.Salt)

	getsql := "select nickname from user where email=? and password=? and disable=0"

	err = bugconfig.Bug_Mysql.GetOne(getsql, login.Username, enpassword).Scan(&login.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			golog.Error("username or password error")
			w.Write(errorcode.Error("username or password error"))
			return
		}
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = asset.Setkey(login.Token, login.Username)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "login",
	}
	err = il.Login("nickname: %s has login", login.Username)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	resp, _ := json.Marshal(login)
	w.Write(resp)
	return

}

func Loginout(w http.ResponseWriter, r *http.Request) {

	errorcode := &errorstruct{}
	conn, err := gadb.NewSqlConfig().ConnDB()
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	defer conn.Db.Close()

	token := r.FormValue("token")
	nickname, err := asset.Getvalue(token)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	asset.Delkey(token)
	il := buglog.AddLog{
		Ip: strings.Split(r.RemoteAddr, ":")[0],
	}
	err = il.Login("nickname: %s has logout", nickname)
	if err != nil {
		golog.Error(err.Error())
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

	errorcode := &errorstruct{}
	nickname, err := logtokenmysql(r)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	userinfo := &userInfo{}

	sql := "select rid,level,headimg,nickname from user where nickname=?"
	var rid string
	var level int
	err = bugconfig.Bug_Mysql.GetOne(sql, nickname).Scan(&rid, &level, &userinfo.Avatar, &userinfo.Name)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		userinfo.Roles = append(userinfo.Roles, "admin")
	} else {
		var rl string
		getrole := "select rolelist from rolegroup where id=?"
		bugconfig.Bug_Mysql.GetOne(getrole, rid).Scan(&rl)
		for _, v := range strings.Split(rl, ",") {
			id, _ := strconv.Atoi(v)
			userinfo.Roles = append(userinfo.Roles, bugconfig.CacheRidRole[int64(id)])
		}
	}

	send, _ := json.Marshal(userinfo)
	w.Write(send)

}
