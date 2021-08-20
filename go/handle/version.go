package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/version"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddVersion(w http.ResponseWriter, r *http.Request) {
	version := xmux.GetInstance(r).Data.(*model.Version)
	version.CreateUid = xmux.GetInstance(r).Get("uid").(int64)
	err := version.Create()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID:         version.Id,
		CreateTime: version.CreateTime,
	}
	w.Write(res.Marshal())

}

func VersionList(w http.ResponseWriter, r *http.Request) {
	vs, err := model.GetAllVersion()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: vs,
	}
	w.Write(res.Marshal())

}

func VersionRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	id := r.FormValue("id")
	var bugcount int

	err := db.Mconn.GetOne("select count(id) from bugs where vid=?", id).Scan(&bugcount)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if bugcount > 0 {
		golog.Errorf("vid:%s has bugs", id)
		w.Write(errorcode.IsUse())
		return
	}
	deletevl := "delete from version where id=?"
	errorcode.ID, err = db.Mconn.Update(deletevl, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// vid, err := strconv.Atoi(id)
	// if err != nil {
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// 增加日志

	send, _ := json.Marshal(errorcode)
	w.Write(send)

}

func VersionUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := xmux.GetInstance(r).Data.(*version.RespVersion)

	uid := xmux.GetInstance(r).Get("uid").(int64)
	versionsql := "update version set pid=(select id from project where name=?),name=?,urlone=?,urltwo=?,createuid=? where id=?"
	_, err := db.Mconn.Update(versionsql, data.Project, data.Name, data.Url, data.BakUrl, uid, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
}
