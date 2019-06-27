package handle

import (
	"bug/asset"
	"bug/bugconfig"
	"bug/buglog"
	"database/sql"
	"encoding/json"
	"gadb"
	"gaencrypt"
	"galog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type resLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Code     int    `json:"statuscode"`
	Avatar   string `json:"avatar"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		login := &resLogin{}

		err = json.Unmarshal(s, login)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		conn, err := gadb.NewSqlConfig().ConnDB()

		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		// 解密
		username, err := gaencrypt.RsaDecrypt(login.Username, bugconfig.PrivateKey, true)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorRsa())
			return
		}

		tmp, err := gaencrypt.RsaDecrypt(login.Password, bugconfig.PrivateKey, true)
		if err != nil {

			galog.Error(err.Error())
			w.Write(errorcode.ErrorRsa())
			return
		}

		login.Username = ""
		login.Password = ""
		login.Token = gaencrypt.Token(string(username), bugconfig.Salt)
		enpassword := gaencrypt.PwdEncrypt(string(tmp), bugconfig.Salt)

		getsql := "select nickname from user where email=? and password=? and disable=0"

		err = conn.GetOne(getsql, string(username), enpassword).Scan(&login.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				galog.Error("username or password error")
				w.Write(errorcode.ErrorUserNameOrPassword())
				return
			}
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		err = asset.Setkey(login.Token, login.Username)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentRedis())
			return
		}

		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "login",
		}
		err = il.Login("nickname: %s has login", login.Username)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		resp, _ := json.Marshal(login)
		w.Write(resp)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func Loginout(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		conn, err := gadb.NewSqlConfig().ConnDB()
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()

		token := r.FormValue("token")
		nickname, err := asset.Getvalue(token)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentRedis())
			return
		}
		asset.Delkey(token)
		il := buglog.AddLog{
			Conn: conn,
			Ip:   strings.Split(r.RemoteAddr, ":")[0],
		}
		err = il.Login("nickname: %s has logout", nickname)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
	}
}

type userInfo struct {
	Roles  []string `json:"roles"`
	Code   int      `json:"statuscode"`
	Avatar string   `json:"avatar"`
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		userinfo := &userInfo{}

		sql := "select rid,level,headimg from user where nickname=?"
		var rid string
		var level int
		err = conn.GetOne(sql, nickname).Scan(&rid, &level, &userinfo.Avatar)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			userinfo.Roles = append(userinfo.Roles, "admin")
		} else {
			var rl string
			getrole := "select rolelist from rolegroup where id=?"
			conn.GetOne(getrole, rid).Scan(&rl)
			for _, v := range strings.Split(rl, ",") {
				id, _ := strconv.Atoi(v)
				userinfo.Roles = append(userinfo.Roles, bugconfig.CacheRidRole[int64(id)])
			}
		}

		send, _ := json.Marshal(userinfo)
		w.Write(send)

	}
}
