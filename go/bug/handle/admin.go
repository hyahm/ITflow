package handle

import (
	"database/sql"
	"github.com/hyahm/golog"
	"itflow/bug/bugconfig"
	"itflow/db"
	"itflow/gaencrypt"
	"net/http"
	"strings"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	errorcode := &errorstruct{}
	addr := strings.Split(r.RemoteAddr, ":")
	if addr[0] != "127.0.0.1" {
		golog.Debug("only 127.0.0.1 cat request")
		w.Write(errorcode.Error("only 127.0.0.1 cat request"))
		return
	}
	password := r.FormValue("password")
	var count int
	row , err := db.Mconn.GetOne("select count(id) from user where rid=0")
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows || count != 1 {
			golog.Debug("有且只能有一个admin账户")
			w.Write(errorcode.Error("有且只能有一个admin账户 \n"))
			return
		}
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	password = gaencrypt.PwdEncrypt(password, bugconfig.Salt)
	_, err = db.Mconn.Update("update user set password=? where rid=0", password)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = insertlog("resetadminpassword", "密码为:"+password, r)
	if err != nil {
		golog.Debug(err.Error())
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	w.Write([]byte("修改成功 \n"))
	return

}
