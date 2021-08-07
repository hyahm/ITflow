package usergroup

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Update(w http.ResponseWriter, r *http.Request) {
	// 可以获取所有用户组
	uid := xmux.GetInstance(r).Get("uid").(int64)
	ug := xmux.GetInstance(r).Data.(*model.UserGroup)
	ug.Uid = uid
	err := ug.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}
