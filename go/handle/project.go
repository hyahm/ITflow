package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/project"
	"itflow/internal/response"
	"net/http"
	"strconv"

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
	send, err := data.Update(uid)
	if err != nil {
		w.Write(send)
		return
	}
	w.Write(send)

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 是否有bug使用
	var count int
	err = db.Mconn.GetOne("select count(id) from bugs where pid=?", id).Scan(&count)
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

	_, err = db.Mconn.Insert(getaritclesql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 更新缓存
	delete(cache.CacheProjectPid, cache.CachePidProject[cache.ProjectId(pid)])
	delete(cache.CachePidProject, cache.ProjectId(pid))

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
