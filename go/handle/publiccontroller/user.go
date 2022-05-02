package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetUserKeyNameByProject(w http.ResponseWriter, r *http.Request) {
	rvkn := xmux.GetInstance(r).Data.(*RequestProject)
	vkns, err := model.GetUserKeyNameByProjectId(rvkn.ProjectId)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = vkns
}

// 获取用户信息
func GetUserKeyName(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)

	kns, err := model.GetUserKeyName(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = kns
}
