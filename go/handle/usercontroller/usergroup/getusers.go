package usergroup

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
)

func GetAllUserKeyName(w http.ResponseWriter, r *http.Request) {
	vkns, err := model.GetAllUserKeyName()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: vkns,
	}
	w.Write(res.Marshal())
}
