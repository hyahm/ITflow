package usergroup

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)
	ug := model.UserGroup{
		Uid: uid,
	}
	err := ug.Delete((id))
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}
