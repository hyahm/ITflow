package important

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	res := response.Response{}
	importants, err := model.GetAllImportant()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res.Data = importants
	w.Write(res.Marshal())
}

func Create(w http.ResponseWriter, r *http.Request) {
	data := xmux.GetInstance(r).Data.(*model.Important)
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
	err := model.DeleteImportant(id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}

func Update(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetInstance(r).Data.(*model.Important)

	err := data.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}
