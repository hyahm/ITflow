package handle

import (
	"itflow/db"
	"itflow/internal/perm"
	"itflow/internal/project"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ProjectList(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Select {
		w.Write(errorcode.Error("no perm"))
		return
	}
	w.Write(project.GetList(uid))
	return

}

func AddProject(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetInstance(r).Data.(*project.ReqProject)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Insert {
		w.Write(errorcode.Error("no perm"))
		return
	}
	send, err := data.Add(uid)
	if err != nil {
		w.Write(send)
		return
	}
	w.Write(send)

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	data := xmux.GetInstance(r).Data.(*project.ReqProject)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Update {
		w.Write(errorcode.Error("no perm"))
		return
	}
	_, err := db.Mconn.Update("update project set name=?, ugid=(select ifnull(min(id),0) from usergroup where name=?) where id=? and uid=?",
		data.ProjectName, data.GroupName, data.Id, uid)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	w.Write(errorcode.Success())

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Delete {
		w.Write(errorcode.Error("no perm"))
		return
	}
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

	_, err = db.Mconn.Delete(getaritclesql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write(errorcode.Success())
	return

}
