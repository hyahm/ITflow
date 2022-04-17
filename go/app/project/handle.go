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
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: projects,
	}
	w.Write(res.Marshal())

}

func Create(w http.ResponseWriter, r *http.Request) {

	project := xmux.GetInstance(r).Data.(*model.Project)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	project.Uid = uid
	err := project.Insert()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID: project.Id,
	}
	w.Write(res.Marshal())

}

func Update(w http.ResponseWriter, r *http.Request) {
	project := xmux.GetInstance(r).Data.(*model.Project)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	err := project.Update(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}

func ProjectKeys(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	pkn, err := model.GetProjectKeyName(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: pkn,
	}
	w.Write(res.Marshal())

}

func Delete(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	id := r.FormValue("id")
	golog.Info(id)
	// 判断有没有bug在使用这个
	var count int
	err := db.Mconn.GetOne("select count(id) from bugs where pid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.IsUse())
		return
	}

	getaritclesql := "delete from project where id=?"

	result := db.Mconn.Delete(getaritclesql, id)
	if result.Err != nil {
		golog.Error(result.Err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write(errorcode.Success())
}
