package version

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	version := xmux.GetInstance(r).Data.(*model.Version)
	version.CreateUid = xmux.GetInstance(r).Get("uid").(int64)
	err := version.Create()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = version.Id
	xmux.GetInstance(r).Response.(*response.Response).CreateTime = version.CreateTime

}

func Read(w http.ResponseWriter, r *http.Request) {
	vs, err := model.GetAllVersion()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = vs

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	err := model.DeleteVersion(id)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}

func Update(w http.ResponseWriter, r *http.Request) {

	version := xmux.GetInstance(r).Data.(*model.Version)
	err := version.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func GetVersion(w http.ResponseWriter, r *http.Request) {

}
