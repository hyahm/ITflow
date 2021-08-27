package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	// 获取显示的状态id
	// sl := xmux.GetInstance(r).Data.(*status.Status)

	uid := xmux.GetInstance(r).Get("uid").(int64)
	ss, err := model.GetShowStatus(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: ss,
	}
	w.Write(res.Marshal())
}
