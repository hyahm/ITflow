package email

import (
	"fmt"
	"itflow/db"
	"net/http"

	"github.com/hyahm/xmux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// errorcode := &response.Response{}

	id := xmux.GetInstance(r).Get("uid")
	var email string
	err := db.Mconn.GetOne("select email from user where id=?", id).Scan(&email)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code": 2, "msg": "%s"}`, err.Error())))
		return
	}
	w.Write([]byte(fmt.Sprintf(`{"code": 0, "email": "%s"}`, email)))
	return

}
