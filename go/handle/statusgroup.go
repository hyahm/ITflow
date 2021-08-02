package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/perm"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Insert {
		w.Write(errorcode.Error("no perm"))
		return
	}
	data := xmux.GetInstance(r).Data.(*model.StatusGroup)
	if data.ID < 0 {
		golog.Error("id not found")
		w.Write(errorcode.Error("id not found"))
		return
	}
	isql := "insert into statusgroup($key) values($value)"
	ids, err := db.Mconn.InsertInterfaceWithID(data, isql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	errorcode.ID = ids[0]
	w.Write(errorcode.Marshal())

}

func EditStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Update {
		w.Write(errorcode.Error("no perm"))
		return
	}
	sg := xmux.GetInstance(r).Data.(*model.StatusGroup)

	isql := "update statusgroup set $set where id = ?"
	_, err := db.Mconn.UpdateInterface(sg, isql, sg.ID)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)

}

func StatusGroupList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Select {
		w.Write(errorcode.Error("no perm"))
		return
	}
	DepartmentList := make([]*model.StatusGroup, 0)
	s := "select id,name,sids from statusgroup"
	err := db.Mconn.Select(&DepartmentList, s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	res := &response.Response{
		Data: DepartmentList,
	}
	w.Write(res.Marshal())
	return

}

func DeleteStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Delete {
		w.Write(errorcode.Error("no perm"))
		return
	}
	id := r.FormValue("id")

	ssql := "select count(id) from jobs where bugsid=?"
	var count int
	err := db.Mconn.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("使用中"))
		return
	}
	isql := "delete from  statusgroup where id = ?"
	_, err = db.Mconn.Update(isql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	//更新缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GetStatusGroupName(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := &struct {
		Names []string `json:"names"`
		Code  int      `json:"code"`
	}{
		Names: make([]string, 0),
	}
	s := "select name from statusgroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		data.Names = append(data.Names, name)

	}
	rows.Close()
	send, _ := json.Marshal(data)
	golog.Info(string(send))
	w.Write(send)
}
