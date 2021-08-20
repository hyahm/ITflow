package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetManagerKeyName(w http.ResponseWriter, r *http.Request) {
	res := response.Response{
		Data: make([]model.KeyName, 0),
	}
	uid := xmux.GetInstance(r).Get("uid").(int64)

	kns, err := model.GetManagerKeyName(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res.Data = kns
	w.Write(res.Marshal())
}
