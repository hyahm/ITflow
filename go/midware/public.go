package midware

import (
	"encoding/json"
	"io/ioutil"
	"itflow/db"
	"itflow/model/response"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func JsonToStruct(w http.ResponseWriter, r *http.Request) bool {
	resp := &response.Response{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	golog.Infof("%+v", xmux.Bridge)
	golog.Infof("%+v", xmux.GetData(r))
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
	golog.Info(a)
	_, err := db.RSconn.Get(a)
	if err != nil {
		if err == redis.Nil {
			// token 没找到或过期
			errorcode.Code = 10
			golog.Info("ffff")
			w.Write(errorcode.ErrorE(err))
			return true
		}
		errorcode.Code = 11
		w.Write(errorcode.ErrorE(err))
		return true

	}
	return false
}
