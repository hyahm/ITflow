package handle

import (
	"itflow/bug/bugconfig"
	"database/sql"
	"itflow/gaencrypt"
	"net/http"
	"strings"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	addr := strings.Split(r.RemoteAddr, ":")
	if addr[0] != "127.0.0.1" {
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
		return
	}
	if r.Method == http.MethodGet {
		password := r.FormValue("password")

		var count int
		err := bugconfig.Bug_Mysql.GetOne("select count(id) from user where department=?", "admin").Scan(&count)
		if err != nil {
			if err == sql.ErrNoRows || count != 1 {
				w.Write([]byte("有且只能有一个admin账户 \n"))
				return
			}
			w.Write([]byte(err.Error()))
			return
		}
		password = gaencrypt.PwdEncrypt(password, bugconfig.Salt)
		_, err = bugconfig.Bug_Mysql.Update("update user set password=? where department=?", password, "admin")
		if err != nil {
			w.Write([]byte(err.Error() + "\n"))
			return
		}
		err = insertlog( "resetadminpassword", "密码为:"+password, r)
		if err != nil {
			w.Write([]byte(err.Error() + "\n"))
			return
		}
		w.Write([]byte("修改成功 \n"))
		return
	}
}
