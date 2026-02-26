package email

import (
	"itflow/db"
	"itflow/response"
	"net/http"

	"github.com/hyahm/xmux"
)

func Set(w http.ResponseWriter, r *http.Request) {
	id := xmux.GetInstance(r).Get("uid")
	email := r.FormValue("email")
	result := db.Mconn.Update("update user set email=? where id=?", email, id)
	if result.Err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = result.Err.Error()
		return
	}
}
