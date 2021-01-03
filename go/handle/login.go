package handle

import (
	"itflow/classify"
	"itflow/internal/user"
	"itflow/model"
	"net/http"
	"sync/atomic"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

var loginSign int32

func Login(w http.ResponseWriter, r *http.Request) {
	if loginSign != 0 {
		w.Write([]byte(`{"code": 0, "msg" : "正在登录"}`))
		return
	}
	atomic.AddInt32(&loginSign, 1)
	defer atomic.AddInt32(&loginSign, -1)
	login := xmux.GetData(r).Data.(*user.Login)
	ipAddr := r.RemoteAddr

	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ipAddr = ip
	}
	golog.Info(ip)
	golog.Info(ipAddr)
	resp, err := login.Check()
	if err != nil {
		golog.Error(err)
		w.Write(resp.ErrorE(err))
		return
	}

	go model.InsertLog(classify.Login, ipAddr, "用户登录", resp.ID)
	w.Write(resp.Marshal())
	return

}

func LoginOut(w http.ResponseWriter, r *http.Request) {

	// 检查token 是否存在

}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	userinfo := &user.UserInfo{}
	userinfo.NickName = xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)
	err := userinfo.GetUserInfo(uid)
	if err != nil {
		golog.Error(err)
		userinfo.Msg = err.Error()
		userinfo.Code = 1
		w.Write(userinfo.Json())
		return
	}
	w.Write(userinfo.Json())
	return
}
