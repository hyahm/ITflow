package handle

import (
	"fmt"
	"itflow/db"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetExpire(w http.ResponseWriter, r *http.Request) {
	token := xmux.Var(r)["token"]
	golog.Info(token)
	filter, err := db.Table.Filter("Token", token)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	time := filter.TTL()

	w.Write([]byte(fmt.Sprintf("%f", time)))
}
