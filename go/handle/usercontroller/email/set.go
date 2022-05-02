package email

import (
	"io/ioutil"
	"itflow/db"
	"itflow/response"
	"net/http"

	"github.com/hyahm/xmux"
)

func Set(w http.ResponseWriter, r *http.Request) {
	id := xmux.GetInstance(r).Get("uid")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	result := db.Mconn.Update("update user set email=? where id=?", string(b), id)
	if result.Err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}
