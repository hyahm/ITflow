package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/perm"
	"itflow/internal/status"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/go-sql-driver/mysql"
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
	data := xmux.GetInstance(r).Data.(*status.StatusGroup)

	golog.Infof("%+v", *data)
	sids, err := data.GetIds()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	isql := "insert into statusgroup(name,sids) values(?,?)"
	errorcode.ID, err = db.Mconn.Insert(isql, data.Name, sids)
	if err != nil {
		golog.Error(err)
		if err.(*mysql.MySQLError).Number == 1062 {
			w.Write(errorcode.ErrorE(db.DuplicateErr))
			return
		}
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Update {
		w.Write(errorcode.Error("no perm"))
		return
	}
	data := xmux.GetInstance(r).Data.(*status.StatusGroup)

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
