package handle

import (
	"itflow/cache"
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
		golog.Info(err)
		// w.Write(sl.Marshal())
		// return
	}
	w.Write(sl.Marshal())
	return

}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetInstance(r).Data.(*defaults.ReqDefaultValue)
	var err error
	cache.DefaultCreateSid, cache.DefaultCompleteSid, err = sl.Update()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info(cache.DefaultCreateSid)
	w.Write(errorcode.Success())
	return

}
