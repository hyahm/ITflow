package usergroup

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func List(w http.ResponseWriter, r *http.Request) {
	// 可以获取所有用户组
	uid := xmux.GetInstance(r).Get("uid").(int64)
	usergroups, err := model.GetUserGroupList(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = usergroups

}
