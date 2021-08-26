package statusgroup

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Create(w http.ResponseWriter, r *http.Request) {

	statusGroup := xmux.GetInstance(r).Data.(*model.StatusGroup)
	err := statusGroup.Insert()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	resp := &response.Response{
		ID: statusGroup.ID,
	}

	w.Write(resp.Marshal())

}

func Update(w http.ResponseWriter, r *http.Request) {
	sg := xmux.GetInstance(r).Data.(*model.StatusGroup)
	err := sg.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}

func Read(w http.ResponseWriter, r *http.Request) {

	sgs, err := model.GetAllStatusGroup()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := &response.Response{
		Data: sgs,
	}
	w.Write(res.Marshal())
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	err := model.DeleteStatusGroup(id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}

	//更新缓存
	w.Write(response.Success())
}

func GetStatusGroupName(w http.ResponseWriter, r *http.Request) {

	kns, err := model.GetStatusGroupKeyName()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	errorcode := &response.Response{
		Data: kns,
	}
	w.Write(errorcode.Marshal())
}
