package handle

import (
	"encoding/json"
	"io/ioutil"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/network/datalog"
	"itflow/network/response"
	"itflow/network/rolegroup"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func RoleGroupList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := &rolegroup.List_roles{}

	s := "select id,name,rolelist from rolegroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var rids string
		one := &rolegroup.Data_roles{}
		rows.Scan(&one.Id, &one.Name, &rids)
		for _, v := range strings.Split(rids, ",") {
			id, _ := strconv.Atoi(v)
			one.RoleList = append(one.RoleList, bugconfig.CacheRidRole[int64(id)])
		}
		data.DataList = append(data.DataList, one)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func RoleGroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}

	ssql := "select count(id) from user where rid=?"
	var count int
	err = db.Mconn.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("没有用户"))
		return
	}
	isql := "delete from  rolegroup where id = ?"
	_, err = db.Mconn.Update(isql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "role",
		Action:   "roledelete",
	}

	// 删除缓存
	delete(bugconfig.CacheRidGroup, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditRoleGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*rolegroup.Data_roles)

	rl := make([]string, 0)
	for _, v := range data.RoleList {
		rl = append(rl, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
	}
	gsql := "update rolegroup set name=?,rolelist=?  where id=?"
	_, err := db.Mconn.Update(gsql, data.Name, strings.Join(rl, ","), data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "role",
		Action:   "update",
	}

	bugconfig.CacheRidGroup[data.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func AddRoleGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &rolegroup.Data_roles{}

	respbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(respbyte, data)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if data.Name == "" {
		golog.Error("name is empty")
		w.Write(errorcode.Error("name is empty"))
		return
	}
	ids := make([]string, 0)
	for _, v := range data.RoleList {
		ids = append(ids, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
	}
	gsql := "insert rolegroup(name,rolelist) values(?,?)"
	errorcode.Id, err = db.Mconn.Insert(gsql, data.Name, strings.Join(ids, ","))
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "role",
		Action:   "add",
	}

	bugconfig.CacheRidGroup[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RoleGroupName(w http.ResponseWriter, r *http.Request) {

	data := &rolegroup.Get_roles{}
	for _, v := range bugconfig.CacheRidGroup {
		data.Roles = append(data.Roles, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
