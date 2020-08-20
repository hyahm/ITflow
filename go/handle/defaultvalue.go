package handle

import (
	"itflow/db"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/golog"

	"itflow/internal/defaults"

	"github.com/hyahm/xmux"
)

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	sl := &defaults.RespDefaultStatus{}
	//如果是管理员的话,所有的都可以

	err := db.Mconn.GetOne("select s.name from defaultvalue as d join status as s on created=s.id ").Scan(&sl.Created)
	err = db.Mconn.GetOne("select s.name from defaultvalue as d join status as s on completed=s.id ").Scan(&sl.Completed)
	if err != nil {
		golog.Error(err)
		w.Write(sl.ErrorE(err))
		return
	}
	w.Write(sl.Marshal())
	return

}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetData(r).Data.(*defaults.ReqDefaultValue)
	golog.Info(sl)
	err := sl.Update()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

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
