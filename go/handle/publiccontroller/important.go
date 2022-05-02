package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func GetImportantKeyName(w http.ResponseWriter, r *http.Request) {
	pkns, err := model.GetAllImportant()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	xmux.GetInstance(r).Response.(*response.Response).Data = pkns
}
