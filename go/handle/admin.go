package handle

import (
	"itflow/control/local"
	"itflow/internal/response"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	addr := strings.Split(r.RemoteAddr, ":")
	if addr[0] != "127.0.0.1" {
		w.Write(errorcode.Error("only 127.0.0.1 cat request"))
		return
	}
	password := r.FormValue("password")
	err := local.ResetAdminPassword(password)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write([]byte("修改成功 \n"))
	return

}
