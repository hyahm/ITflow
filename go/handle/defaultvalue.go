package handle

import (
	"itflow/db"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"

	"github.com/hyahm/xmux"
)

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	sl := &model.DefaultValue{}
	res := response.Response{}
	//如果是管理员的话,所有的都可以
	err := db.Mconn.GetOne("select created,completed from defaultvalue").Scan(&sl.Created, &sl.Completed)
	if err != nil {
		golog.Info(err)
		w.Write(res.ErrorE(err))
		return
	}
	res.Data = sl
	w.Write(res.Marshal())
}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetInstance(r).Data.(*model.DefaultValue)
	err := sl.Update()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	w.Write(errorcode.Success())
	return

}
