package handle

import (
	"database/sql"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/gaencrypt"
	"itflow/network/response"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	addr := strings.Split(r.RemoteAddr, ":")
	if addr[0] != "127.0.0.1" {
		golog.Debug("only 127.0.0.1 cat request")
		w.Write(errorcode.Error("only 127.0.0.1 cat request"))
		return
	}
	password := r.FormValue("password")
	var count int
	err := db.Mconn.GetOne("select count(id) from user where rid=0").Scan(&count)

	if err != nil {
		if err == sql.ErrNoRows || count != 1 {
			golog.Debug("有且只能有一个admin账户")
			w.Write(errorcode.Error("有且只能有一个admin账户 \n"))
			return
		}
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	password = gaencrypt.PwdEncrypt(password, bugconfig.Salt)
	_, err = db.Mconn.Update("update user set password=? where rid=0", password)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// xmux.GetData(r).End = &datalog.AddLog{
	// 	Ip:       r.RemoteAddr,
	// 	Username: nickname,
	// 	Classify: "superuser",
	// 	Action:   "update",
	// }

	w.Write([]byte("修改成功 \n"))
	return

}
