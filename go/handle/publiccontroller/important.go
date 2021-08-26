package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
)

func GetImportantKeyName(w http.ResponseWriter, r *http.Request) {
	pkns, err := model.GetAllImportant()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: pkns,
	}
	w.Write(res.Marshal())
}
