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
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = statusGroup.ID
}

func Update(w http.ResponseWriter, r *http.Request) {
	sg := xmux.GetInstance(r).Data.(*model.StatusGroup)
	err := sg.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}

func Read(w http.ResponseWriter, r *http.Request) {

	sgs, err := model.GetAllStatusGroup()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = sgs
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	err := model.DeleteStatusGroup(id)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func GetStatusGroupName(w http.ResponseWriter, r *http.Request) {

	kns, err := model.GetStatusGroupKeyName()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = kns
}
