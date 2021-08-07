package usercontroller

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/encrypt"
	"itflow/internal/user"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}

	getuser := xmux.GetInstance(r).Data.(*user.ChangePasswod)

	uid := xmux.GetInstance(r).Get("uid").(int64)

	getaritclesql := "select count(id) from user where id=? and password=?"
	oldpassword := encrypt.PwdEncrypt(getuser.Oldpassword, cache.Salt)
	var n int
	err := db.Mconn.GetOne(getaritclesql, uid, oldpassword).Scan(&n)
	if err != nil || n != 1 {
		golog.Error(err)
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	newpassword := encrypt.PwdEncrypt(getuser.Newpassword, cache.Salt)
	chpwdsql := "update user set password=? where id=?"

	_, err = db.Mconn.Update(chpwdsql, newpassword, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}
