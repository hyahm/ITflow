package midware

import (
	"encoding/json"
	"io/ioutil"
	"itflow/db"
	"itflow/network/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func JsonToStruct(w http.ResponseWriter, r *http.Request) bool {
	resp := &response.Response{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil || len(b) == 0 {
		w.Write(resp.ErrorE(err))
		return true
	}
	golog.Info(string(b))
	err = json.Unmarshal(b, xmux.GetData(r).Data)
	if err != nil {
		golog.Error(err)
		w.Write(resp.ErrorE(err))
		return true
	}
	return false
}

func CheckToken(w http.ResponseWriter, r *http.Request) bool {
	errorcode := &response.Response{}
	a := r.Header.Get("X-Token")
	if a == "" {
		w.Write(errorcode.TokenNotFound())
		return true
	}
	if filter, err := db.CT.Filter("Token", a); err != nil {
		w.Write(errorcode.TokenNotFound())
		return true
	} else {
		var nickname string
		filter.Get("NickName").Scan(&nickname)
		xmux.GetData(r).Set("nickname", nickname)
	}

	return false
}
