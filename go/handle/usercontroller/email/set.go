package email

import (
	"io/ioutil"
	"itflow/db"
	"itflow/response"
	"net/http"

	"github.com/hyahm/xmux"
)

func Set(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	id := xmux.GetInstance(r).Get("uid")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	result := db.Mconn.Update("update user set email=? where id=?", string(b), id)
	if result.Err != nil {
		w.Write(errorcode.ErrorE(result.Err))
		return
	}
	w.Write(errorcode.Success())
	return
}
