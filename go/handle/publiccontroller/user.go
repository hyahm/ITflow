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
		w.Write(response.ErrorE(err))
		return
	}
	golog.Info(rvkn)
	res := response.Response{
		Data: vkns,
	}
	w.Write(res.Marshal())
}

// 获取用户信息
func GetUserKeyName(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)

	kns, err := model.GetUserKeyName(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: kns,
	}
	w.Write(res.Marshal())
}
