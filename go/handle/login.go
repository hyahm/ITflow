package handle

import (
	"itflow/classify"
	"itflow/internal/user"
	"itflow/model"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Login(w http.ResponseWriter, r *http.Request) {

	login := xmux.GetData(r).Data.(*user.Login)

	resp, err := login.Check()
	if err != nil {
		golog.Error(err)
		w.Write(resp.ErrorE(err))
		return
	}
	go model.InsertLog(classify.Login, r.RemoteAddr, "用户登录", resp.ID)
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
