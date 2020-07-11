package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/golog"

	"itflow/internal/defaults"

	"github.com/hyahm/xmux"
)

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	sl := &defaults.RespDefaultStatus{}
	//如果是管理员的话,所有的都可以
	sl.Created = cache.DefaultCreateSid.Name()
	sl.Completed = cache.DefaultCompleteSid.Name()
	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetData(r).Data.(*defaults.ReqDefaultValue)

	err := sl.Save()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// cache.CacheDefault["important"] = iid
	// cache.CacheDefault["level"] = lid
	w.Write(errorcode.Success())
	return

}

// type defaultImportant struct {
// 	DefaultImportant string `json:"defaultimportant"`
// 	Code             int    `json:"code"`
// }

// func DefaultImportant(w http.ResponseWriter, r *http.Request) {

// 	data := &defaultImportant{}
// 	data.DefaultImportant = cache.CacheIidImportant[cache.CacheDefault["important"]]
// 	send, _ := json.Marshal(data)
// 	w.Write(send)
// 	return

// }

// type defaultLevel struct {
// 	DefaultLevel string `json:"defaultlevel"`
// 	Code         int    `json:"code"`
// }

// func DefaultLevel(w http.ResponseWriter, r *http.Request) {

// 	data := &defaultLevel{}
// 	data.DefaultLevel = cache.CacheLidLevel[cache.CacheDefault["level"]]
// 	send, _ := json.Marshal(data)
// 	w.Write(send)
// 	return

// }
