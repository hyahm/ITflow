package handle

import (
	"encoding/json"
	"fmt"
	"github.com/hyahm/golog"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type getDepartment struct {
	Id         int64    `json:"id"`
	StatusList []string `json:"checklist"`
	Department string   `json:"departmentname"`
}

func AddBugGroup(w http.ResponseWriter, r *http.Request) {
	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &getDepartment{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("statusgroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	list, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(list, data)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 重新排序
	ids := make([]string, 0)
	for _, v := range data.StatusList {
		ids = append(ids, fmt.Sprintf("%d", bugconfig.CacheStatusSid[v]))
	}

	ss := strings.Join(ids, ",")

	isql := "insert into statusgroup(name,sids) values(?,?)"
	errorcode.Id, err = bugconfig.Bug_Mysql.Insert(isql, data.Department, ss)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "statusgroup",
	}
	err = il.Add(
		nickname, errorcode.Id, data.Department)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 添加缓存
	bugconfig.CacheSgidGroup[errorcode.Id] = data.Department
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditBugGroup(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &getDepartment{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("statusgroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	list, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(list, data)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	ssids := make([]string, 0)
	for _, v := range data.StatusList {
		// 没找到就是key
		var sid int64
		var ok bool
		if sid, ok = bugconfig.CacheStatusSid[v]; !ok {
			w.Write(errorcode.Error("没有找到status"))
			return
		}
		ssids = append(ssids, strconv.FormatInt(sid, 10))
	}
	ss := strings.Join(ssids, ",")
	isql := "update statusgroup set name =?,sids=? where id = ?"
	_, err = bugconfig.Bug_Mysql.Update(isql, data.Department, ss, data.Id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "statusgroup",
	}
	err = il.Update(
		nickname, data.Id, data.Department)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	bugconfig.CacheSgidGroup[data.Id] = data.Department
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type department struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	BugstatusList []string `json:"bugstatuslist"`
}

type departmentList struct {
	DepartmentList []*department `json:"departmentlist"`
	Code           int           `json:"code"`
}

func BugGroupList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("statusgroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	data := &departmentList{}
	s := "select id,name,sids from statusgroup"
	rows, err := bugconfig.Bug_Mysql.GetRows(s)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var ids string
		one := &department{}
		rows.Scan(&one.Id, &one.Name, &ids)

		for _, v := range strings.Split(ids, ",") {
			id, err := strconv.Atoi(v)
			if err != nil {
				log.Println(err)
				continue
			}

			one.BugstatusList = append(one.BugstatusList, bugconfig.CacheSidStatus[int64(id)])
		}
		data.DepartmentList = append(data.DepartmentList, one)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func BugGroupDel(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
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
		permssion, err = asset.CheckPerm("statusgroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	ssql := "select count(id) from user where bugsid=?"
	var count int
	err = bugconfig.Bug_Mysql.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	if count > 0 {
		w.Write(errorcode.Error("没有找到group"))
		return
	}
	isql := "delete from  statusgroup where id = ?"
	_, err = bugconfig.Bug_Mysql.Update(isql, id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "statusgroup",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更新缓存
	delete(bugconfig.CacheSgidGroup, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupGet(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Send_groups{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("usergroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	gsql := "select id,name,ids from usergroup"
	rows, err := bugconfig.Bug_Mysql.GetRows(gsql)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		onegroup := &model.Get_groups{}
		var users string
		rows.Scan(&onegroup.Id, &onegroup.Name, &users)
		for _, v := range strings.Split(users, ",") {
			uid, _ := strconv.Atoi(v)
			onegroup.Users = append(onegroup.Users, bugconfig.CacheUidRealName[int64(uid)])
		}
		data.GroupList = append(data.GroupList, onegroup)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func GroupAdd(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("usergroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	data := &model.Get_groups{}
	respbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(respbyte, data)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	ids := make([]string, 0)
	for _, v := range data.Users {
		var uid int64
		var ok bool
		if uid, ok = bugconfig.CacheRealNameUid[v]; !ok {
			w.Write(errorcode.Errorf("没有此用户"))
			return
		}
		ids = append(ids, strconv.FormatInt(uid, 10))
	}
	gsql := "insert usergroup(name,ids,cuid) values(?,?,?)"
	errorcode.Id, err = bugconfig.Bug_Mysql.Insert(gsql, data.Name, strings.Join(ids, ","), bugconfig.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "usergroup",
	}
	err = il.Add(
		nickname, errorcode.Id, data.Name)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	bugconfig.CacheGidGroup[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupDel(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("usergroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 判断共享文件是否有在使用

	var hasshare bool
	sharerows, err := bugconfig.Bug_Mysql.GetRows("select readuser,rid,wid,writeuser from  sharefile")
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	for sharerows.Next() {
		var readuser bool
		var rid int64
		var wid int64
		var writeuser bool
		sharerows.Scan(&readuser, &rid, &wid, &writeuser)
		if !readuser {
			if rid == int64(id32) {
				hasshare = true
				return
			}
		}
		if !writeuser {
			if wid == int64(id32) {
				hasshare = true
				return
			}
		}
	}

	if hasshare {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	// 判断接口是否有在使用
	var hasrest bool
	restrows, err := bugconfig.Bug_Mysql.GetRows("select readuser,edituser,rid,eid from  apiproject")
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	for restrows.Next() {
		var readuser bool
		var rid int64
		var eid int64
		var edituser bool
		restrows.Scan(&readuser, &rid, &eid, &edituser)
		if !readuser {
			if rid == int64(id32) {
				hasshare = true
			}
		}
		if !edituser {
			if eid == int64(id32) {
				hasshare = true
				return
			}
		}
	}
	if hasrest {
		w.Write(errorcode.Error("没有权限"))
		return
	}

	gsql := "delete from usergroup where id=? and cuid=?"
	_, err = bugconfig.Bug_Mysql.Update(gsql, id, bugconfig.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "usergroup",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	delete(bugconfig.CacheGidGroup, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupUpdate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Get_groups{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("usergroup", nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	respbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(respbyte, data)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	gsql := "update usergroup set name=?,ids=? where id=? and cuid=?"
	ids := ""
	for i, v := range data.Users {
		if i == 0 {
			ids = strconv.FormatInt(bugconfig.CacheRealNameUid[v], 10)
		} else {
			ids = ids + "," + strconv.FormatInt(bugconfig.CacheRealNameUid[v], 10)
		}
	}
	_, err = bugconfig.Bug_Mysql.Update(gsql, data.Name, ids, data.Id, bugconfig.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "usergroup",
	}
	err = il.Update(
		nickname, data.Id, data.Name)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	bugconfig.CacheGidGroup[data.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
