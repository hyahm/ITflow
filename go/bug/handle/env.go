package handle

import (
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type envlist struct {
	Elist []*envrow `json:"envlist"`
	Code  int       `json:"statuscode"`
}

type envrow struct {
	Id      int64  `json:"id"`
	EnvName string `json:"envname"`
}

func EnvList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		env := &envlist{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("env", nickname, conn)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
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
	w.WriteHeader(http.StatusNotFound)
}

func AddEnv(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("env", nickname, conn)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		envname := r.FormValue("name")

		getaritclesql := "insert into environment(envname) values(?)"

		errorcode.Id, err = conn.InsertWithID(getaritclesql, envname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "env",
		}
		err = il.Add(
			nickname, errorcode.Id, envname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 添加缓存
		bugconfig.CacheEidName[errorcode.Id] = envname
		bugconfig.CacheEnvNameEid[envname] = errorcode.Id
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateEnv(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}

		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		er := &envrow{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("env", nickname, conn)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		bpr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(bpr, er)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		getaritclesql := "update environment set envname=? where id=?"

		_, err = conn.Update(getaritclesql, er.EnvName, er.Id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "env",
		}
		err = il.Update(
			nickname, er.Id, er.EnvName)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
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
	w.WriteHeader(http.StatusNotFound)
}

func DeleteEnv(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorNotFoundToken())
			return
		}
		defer conn.Db.Close()
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("env", nickname, conn)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		id := r.FormValue("id")
		eid, err := strconv.Atoi(id)
		if err != nil {

			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var count int

		err = conn.GetOne("select count(id) from bugs where eid=?", id).Scan(&count)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if count > 0 {
			w.Write(errorcode.ErrorHasEnv())
			return
		}
		getaritclesql := "delete from environment where id=?"

		_, err = conn.Update(getaritclesql, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "env",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		delete(bugconfig.CacheEnvNameEid, bugconfig.CacheEidName[int64(eid)])
		delete(bugconfig.CacheEidName, int64(eid))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
