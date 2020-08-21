package handle

import (
	"itflow/db"
	"itflow/internal/project"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ProjectList(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetData(r).Get("uid").(int64)
	w.Write(project.GetList(uid))
	return

}

func AddProject(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetData(r).Data.(*project.ReqProject)
	uid := xmux.GetData(r).Get("uid").(int64)
	send, err := data.Add(uid)
	if err != nil {
		w.Write(send)
		return
	}
	w.Write(send)

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	data := xmux.GetData(r).Data.(*project.ReqProject)
	uid := xmux.GetData(r).Get("uid").(int64)
	errorcode := &response.Response{}
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
