package project

import (
	"itflow/db"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	projects, err := model.GetAllProjects(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = projects

}

func Create(w http.ResponseWriter, r *http.Request) {

	project := xmux.GetInstance(r).Data.(*model.Project)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	project.Uid = uid
	err := project.Insert()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = project.Id
}

func Update(w http.ResponseWriter, r *http.Request) {
	project := xmux.GetInstance(r).Data.(*model.Project)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	err := project.Update(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}

func ProjectKeys(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	pkn, err := model.GetProjectKeyName(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = pkn

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	// 判断有没有bug在使用这个
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where pid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	if count > 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "使用中，无法删除"
		return
	}

	getaritclesql := "delete from project where id=?"

	result := db.Mconn.Delete(getaritclesql, id)
	if result.Err != nil {
		golog.Error(result.Err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}
