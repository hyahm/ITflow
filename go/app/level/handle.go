package level

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {

	res := response.Response{}
	Levels, err := model.GetLevelKeyNameByUid()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res.Data = Levels

	w.Write(res.Marshal())
}

func Create(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetInstance(r).Data.(*model.Level)
	err := data.Create()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID: data.Id,
	}
	w.Write(res.Marshal())
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	err := model.DeleteLevel(id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}

func Update(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetInstance(r).Data.(*model.Level)
	err := data.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}

	w.Write(response.Success())
}
