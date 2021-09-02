package handle

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"

	"github.com/hyahm/xmux"
)

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	res := response.Response{}
	//如果是管理员的话,所有的都可以
	res.Data = model.Default
	w.Write(res.Marshal())
}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetInstance(r).Data.(*model.DefaultValue)
	golog.Infof("%#v", sl)
	err := sl.Update()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	model.Default.Created = sl.Created
	model.Default.Completed = sl.Completed
	model.Default.Pass = sl.Pass
	model.Default.Receive = sl.Receive
	w.Write(errorcode.Success())
}
