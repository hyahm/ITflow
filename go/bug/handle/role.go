package handle

import (
	"encoding/json"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"itflow/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

func RoleList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.List_roles{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("rolegroup", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	s := "select id,name,rolelist from rolegroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var rids string
		one := &model.Data_roles{}
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

func RoleDel(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("rolegroup", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	ssql := "select count(id) from user where rid=?"
	var count int
	row, err := db.Mconn.GetOne(ssql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&count)
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
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "rolegroup",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 删除缓存
	delete(bugconfig.CacheRidGroup, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditRole(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Data_roles{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("rolegroup", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	bytedata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(bytedata, data)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	rl := make([]string, 0)
	for _, v := range data.RoleList {
		rl = append(rl, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
	}
	gsql := "update rolegroup set name=?,rolelist=?  where id=?"
	_, err = db.Mconn.Update(gsql, data.Name, strings.Join(rl, ","), data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "rolegroup",
	}
	err = il.Update(
		nickname, data.Id, data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	bugconfig.CacheRidGroup[data.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func AddRole(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)

		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Data_roles{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("rolegroup", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
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
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "rolegroup",
	}
	err = il.Add(
		nickname, errorcode.Id, data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	bugconfig.CacheRidGroup[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RoleGroupName(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Get_roles{}
	for _, v := range bugconfig.CacheRidGroup {
		data.Roles = append(data.Roles, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
