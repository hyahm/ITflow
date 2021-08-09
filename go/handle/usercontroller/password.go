package usercontroller

import (
	"itflow/cache"
	"itflow/encrypt"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type ChangePasswod struct {
	Oldpassword string `json:"oldpassword"`
	Newpassword string `json:"newpassword"`
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	getuser := xmux.GetInstance(r).Data.(*ChangePasswod)
	oldpassword := encrypt.PwdEncrypt(getuser.Oldpassword, cache.Salt)
	newpassword := encrypt.PwdEncrypt(getuser.Newpassword, cache.Salt)
	user := model.User{
		ID:       uid,
		Password: newpassword,
	}
	err := user.UpdatePassword(oldpassword)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}
