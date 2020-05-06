package midware

import (
	"encoding/json"
	"io/ioutil"
	"itflow/db"
	"itflow/network/response"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func JsonToStruct(w http.ResponseWriter, r *http.Request) bool {
	golog.Info("aaaasdgsdf")
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
		errorcode.Code = 10
		w.Write(errorcode.ErrorE(redis.Nil))
		return true
	}
	nickname, err := db.RSconn.Get(a)
	if err != nil {
		if err == redis.Nil {
			// token 没找到或过期
			errorcode.Code = 10
			w.Write(errorcode.ErrorE(err))
			return true
		}
		errorcode.Code = 11
		w.Write(errorcode.ErrorE(err))
		return true

	}
	xmux.GetData(r).Set("nickname", nickname)
	return false
}
