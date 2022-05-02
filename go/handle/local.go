package handle

import (
	"itflow/db"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetExpire(w http.ResponseWriter, r *http.Request) {
	token := xmux.Var(r)["token"]
	golog.Info(token)
	filter, err := db.Table.Filter("Token", token)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	time := filter.TTL()
	xmux.GetInstance(r).Response.(*response.Response).UpdateTime = int64(time)
}
