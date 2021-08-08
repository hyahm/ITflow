package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type RequestProject struct {
	ProjectId int64 `json:"project_id"`
}

func GetVersionKeyNameByProject(w http.ResponseWriter, r *http.Request) {
	rvkn := xmux.GetInstance(r).Data.(*RequestProject)
	vkns, err := model.GetVersionKeyNameByProjectId(rvkn.ProjectId)
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
