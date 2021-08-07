package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetUserKeyName(w http.ResponseWriter, r *http.Request) {
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
