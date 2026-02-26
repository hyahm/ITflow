package handle

import (
	"itflow/classify"
	"itflow/internal/user"
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Login(w http.ResponseWriter, r *http.Request) {

	login := xmux.GetInstance(r).Data.(*user.Login)

	resp, uid, err := login.Check()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	log := model.Log{
		CreateTime: time.Now(),
		Classify:   string(classify.Login),
		Ip:         strings.Split(ip, ":")[0],
		Uid:        uid,
		Action:     "用户登录",
	}
	err = log.Insert()
	if err != nil {
		golog.Error(err)
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = resp
}

func LoginOut(w http.ResponseWriter, r *http.Request) {

	// 检查token 是否存在

}

func UserInfo(w http.ResponseWriter, r *http.Request) {

	userinfo := &user.UserInfo{}
	// userinfo.NickName = xmux.GetInstance(r).Get("nickname").(string)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	userinfo.Uid = uid
	err := userinfo.GetUserInfo(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	if len(userinfo.Roles) == 0 {
		userinfo.Roles = append(userinfo.Roles, "test")
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = userinfo
}
