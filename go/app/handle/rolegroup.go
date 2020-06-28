package handle

import (
	"encoding/json"
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
	data := &rolegroup.RoleGroupList{}

	s := "select id,name,rolelist from rolegroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var rids string
		one := &rolegroup.RoleGroup{}
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

func GetRoleGroupName(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	s := "select name from rolegroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	resp := struct {
		Code     int      `json:"code"`
		RoleList []string `json:"rolelist"`
	}{
		RoleList: make([]string, 0),
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		resp.RoleList = append(resp.RoleList, name)

	}
	send, _ := json.Marshal(resp)
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

	data := xmux.GetData(r).Data.(*rolegroup.RoleGroup)

	if err := data.Update(); err != nil {
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

	data := xmux.GetData(r).Data.(*rolegroup.RoleGroup)

	if data.Name == "" {
		golog.Error("name is empty")
		w.Write(errorcode.Error("name is empty"))
		return
	}
	// 插入数据
	if err := data.Insert(); err != nil {
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

	bugconfig.CacheRidGroup[data.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RoleGroupName(w http.ResponseWriter, r *http.Request) {

	data := &rolegroup.Roles{}
	for _, v := range bugconfig.CacheRidGroup {
		data.Roles = append(data.Roles, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func GetRoleGroup(w http.ResponseWriter, r *http.Request) {

	rl := &rolegroup.Roles{}
	for _, v := range bugconfig.CacheRidGroup {
		rl.Roles = append(rl.Roles, v)
	}
	send, _ := json.Marshal(rl)
	w.Write(send)
	return

}
