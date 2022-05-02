package handle

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"

	"github.com/hyahm/xmux"
)

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	//如果是管理员的话,所有的都可以
	xmux.GetInstance(r).Response.(*response.Response).Data = model.Default
}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	sl := xmux.GetInstance(r).Data.(*model.DefaultValue)
	err := sl.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	model.Default.Created = sl.Created
	model.Default.Completed = sl.Completed
	model.Default.Pass = sl.Pass
	model.Default.Receive = sl.Receive
}
