package usergroup

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// 可以获取所有用户组
	ug := xmux.GetInstance(r).Data.(*model.UserGroup)
	ug.Uid = xmux.GetInstance(r).Get("uid").(int64)
	err := ug.Create()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = ug.Id
}
