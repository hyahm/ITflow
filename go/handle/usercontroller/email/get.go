package email

import (
	"itflow/db"
	"itflow/response"
	"net/http"

	"github.com/hyahm/xmux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	// errorcode := &response.Response{}

	id := xmux.GetInstance(r).Get("uid")
	var email string
	err := db.Mconn.GetOne("select email from user where id=?", id).Scan(&email)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = email
}
