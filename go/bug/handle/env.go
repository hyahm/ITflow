package handle

import (
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/db"
	"net/http"
	"strconv"
	"strings"
)

type envlist struct {
	Elist []*envrow `json:"envlist"`
	Code  int       `json:"code"`
}

type envrow struct {
	Id      int64  `json:"id"`
	EnvName string `json:"envname"`
}

func EnvList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	env := &envlist{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("env", nickname)
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
	for k, v := range bugconfig.CacheEidName {
		pr := &envrow{
			Id:      k,
			EnvName: v,
		}
		env.Elist = append(env.Elist, pr)
	}

	send, _ := json.Marshal(env)
	w.Write(send)
	return

}

func AddEnv(w http.ResponseWriter, r *http.Request) {

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
		permssion, err = asset.CheckPerm("env", nickname)
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
	envname := r.FormValue("name")

	getaritclesql := "insert into environment(envname) values(?)"

	errorcode.Id, err = db.Mconn.Insert(getaritclesql, envname)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "env",
	}
	err = il.Add(
		nickname, errorcode.Id, envname)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 添加缓存
	bugconfig.CacheEidName[errorcode.Id] = envname
	bugconfig.CacheEnvNameEid[envname] = errorcode.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func UpdateEnv(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}

	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	er := &envrow{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("env", nickname)
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
	bpr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(bpr, er)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	getaritclesql := "update environment set envname=? where id=?"

	_, err = db.Mconn.Update(getaritclesql, er.EnvName, er.Id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "env",
	}
	err = il.Update(
		nickname, er.Id, er.EnvName)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 更新缓存
	delete(bugconfig.CacheEnvNameEid, bugconfig.CacheEidName[int64(er.Id)])
	bugconfig.CacheEidName[int64(er.Id)] = er.EnvName
	bugconfig.CacheEnvNameEid[er.EnvName] = int64(er.Id)
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func DeleteEnv(w http.ResponseWriter, r *http.Request) {

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
		permssion, err = asset.CheckPerm("env", nickname)
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
	eid, err := strconv.Atoi(id)
	if err != nil {

		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	var count int

	row, err := db.Mconn.GetOne("select count(id) from bugs where eid=?", id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&count)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	if count > 0 {
		w.Write(errorcode.Error("存在此env"))
		return
	}
	getaritclesql := "delete from environment where id=?"

	_, err = db.Mconn.Update(getaritclesql, id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "env",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	delete(bugconfig.CacheEnvNameEid, bugconfig.CacheEidName[int64(eid)])
	delete(bugconfig.CacheEidName, int64(eid))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
