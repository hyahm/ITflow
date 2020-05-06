package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/network/datalog"
	"itflow/network/response"
	"itflow/network/user"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Login(w http.ResponseWriter, r *http.Request) {

	login := xmux.GetData(r).Data.(*user.Login)

	resp := &user.RespLogin{}
	errresp := login.Check(resp)
	if errresp != nil {
		golog.Error(string(errresp))
		w.Write(errresp)
		return
	}

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: login.Username,
		Classify: "login",
		Action:   "login",
	}
	send, _ := json.Marshal(resp)
	w.Write(send)
	return

}

func LoginOut(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	// 检查token 是否存在
	token := r.Header.Get("X-Token")
	nickname, err := db.RSconn.Get(token)
	if err != nil {

		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "loginout",
	}

}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	userinfo := &user.UserInfo{}
	nickname, err := logtokenmysql(r)
	if err != nil {
		userinfo.Msg = err.Error()
		userinfo.Code = 1
		w.Write(userinfo.Json())
		return
	}
	userinfo.Name = nickname

	err = userinfo.GetUserInfo()
	if err != nil {
		golog.Error(err)
		userinfo.Msg = err.Error()
		userinfo.Code = 2
		w.Write(userinfo.Json())
		return
	}
	w.Write(userinfo.Json())

}
