package handle

import (
	"itflow/cache"
	"itflow/encrypt"
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	addr := strings.Split(r.RemoteAddr, ":")
	if addr[0] != "127.0.0.1" {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "only 127.0.0.1 cat request"
		return
	}
	password := r.FormValue("password")
	user := model.User{}

	enpassword := encrypt.PwdEncrypt(password, cache.Salt)

	err := user.UpdateAdminPassword(enpassword)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}
