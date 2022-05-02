package position

import (
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	jobs, err := model.GetAllPositions()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = jobs
}

func PositionGet(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}

	job := xmux.GetInstance(r).Data.(*model.Job)
	if strings.Trim(job.Name, " ") == "" {
		xmux.GetInstance(r).Response.(*response.Response).Msg = "职位名不能为空"
		return
	}
	if job.RoleGroup <= 0 {
		xmux.GetInstance(r).Response.(*response.Response).Msg = "角色组不能为空"
		return
	}

	err := job.Insert()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = job.Id

}

func Delete(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}
	id := r.FormValue("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)

	err := model.DeleteJob(id, uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}

func Update(w http.ResponseWriter, r *http.Request) {

	job := xmux.GetInstance(r).Data.(*model.Job)
	err := job.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}
