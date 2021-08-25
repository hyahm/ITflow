package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

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
}

func DeleteStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")

	ssql := "select count(id) from jobs where sgid=?"
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
}

func GetStatusGroupName(w http.ResponseWriter, r *http.Request) {

	kns, err := model.GetStatusGroupKeyName()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	errorcode := &response.Response{
		Data: kns,
	}
	w.Write(errorcode.Marshal())
}
