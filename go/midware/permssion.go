package midware

import (
	"fmt"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/xmux"
)

const (
	SELECT = 1
	REMOVE = 2
	UPDATE = 4
	CREATE = 8
)

type UserChecker interface {
	CheckUser(uid int64) error
}

func CheckUser(w http.ResponseWriter, r *http.Request) bool {

	uid := xmux.GetData(r).Get("uid").(int64)
	resp := &response.Response{}
	if xmux.GetData(r).Data == nil {
		err := fmt.Sprintf("must be bind data first %s", r.URL.RequestURI())
		w.Write(resp.Error(err))
		return true
	}
	err := xmux.GetData(r).Data.(UserChecker).CheckUser(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	return false
}
