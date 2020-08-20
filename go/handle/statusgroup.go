package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/response"
	"itflow/internal/status"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
	"github.com/hyahm/xmux"
)

func AddStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	golog.Infof("%+v", *data)
	sids, err := data.GetIds()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	isql := "insert into statusgroup(name,sids) values(?,?)"
	errorcode.Id, err = db.Mconn.Insert(isql, data.Name, sids)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	if data.Name == "" {
		w.Write(errorcode.Error("名称不能为空"))
		return
	}
	sids, err := data.GetIds()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	isql := "update statusgroup set name =?,sids=? where id = ?"
	_, err = db.Mconn.Update(isql, data.Name, sids, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusGroupList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &departmentList{}
	s := "select id,name,sids from statusgroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var ids string
		one := &department{}

		err = rows.Scan(&one.Id, &one.Name, &ids)
		if err != nil {
			golog.Error()
			continue
		}

		idrows, err := db.Mconn.GetRowsIn("select name from status where id in (?)",
			(gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

		for idrows.Next() {
			var name string
			err = idrows.Scan(&name)
			if err != nil {
				golog.Error()
				continue
			}
			if name != "" {
				one.BugstatusList = append(one.BugstatusList, name)
			}

		}

		data.DepartmentList = append(data.DepartmentList, one)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func DeleteStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")

	ssql := "select count(id) from user where bugsid=?"
	var count int
	err := db.Mconn.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("有人再使用"))
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
	send, _ := json.Marshal(data)
	golog.Info(string(send))
	w.Write(send)
	return

}
