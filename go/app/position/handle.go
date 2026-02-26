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
	position := model.Position{}
	positions, err := position.GetAllPositions()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = positions
}

func PositionGet(w http.ResponseWriter, r *http.Request) {

}

func Create(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}

	position := xmux.GetInstance(r).Data.(*model.Position)
	if strings.Trim(position.Name, " ") == "" {
		xmux.GetInstance(r).Response.(*response.Response).Msg = "职位名不能为空"
		return
	}
	if position.RoleId <= 0 {
		xmux.GetInstance(r).Response.(*response.Response).Msg = "角色不能为空"
		return
	}

	err := position.Create()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = position.Id

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

func ManagerList(w http.ResponseWriter, r *http.Request) {
	// errorcode := &response.Response{}
	position := model.Position{}
	ps, err := position.GetManager()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = ps
}

func Update(w http.ResponseWriter, r *http.Request) {

	job := xmux.GetInstance(r).Data.(*model.Position)
	err := job.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}
