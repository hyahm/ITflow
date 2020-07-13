package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/env"
	"itflow/internal/response"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func EnvList(w http.ResponseWriter, r *http.Request) {

	el := &env.Envlist{
		Elist: make([]*env.Env, 0),
	}

	// 管理员
	for k, v := range cache.CacheEidEnv {
		pr := &env.Env{
			Id:      k,
			EnvName: v,
		}
		el.Elist = append(el.Elist, pr)
	}

	send, _ := json.Marshal(el)
	w.Write(send)
	return

}

func AddEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	envname := r.FormValue("name")

	getaritclesql := "insert into environment(envname) values(?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(getaritclesql, envname)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "env",
		Action:   "add",
	}

	// 添加缓存
	cache.CacheEidEnv[cache.EnvId(errorcode.Id)] = cache.Env(envname)
	cache.CacheEnvEid[cache.Env(envname)] = cache.EnvId(errorcode.Id)
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func UpdateEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	er := xmux.GetData(r).Data.(*env.Env)

	getaritclesql := "update environment set envname=? where id=?"

	_, err := db.Mconn.Update(getaritclesql, er.EnvName, er.Id)
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
		Classify: "bug",
		Action:   "update",
	}

	// 更新缓存
	delete(cache.CacheEnvEid, cache.CacheEidEnv[cache.EnvId(er.Id)])
	cache.CacheEidEnv[cache.EnvId(er.Id)] = er.EnvName
	cache.CacheEnvEid[er.EnvName] = cache.EnvId(er.Id)
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func DeleteEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	eid, err := strconv.Atoi(id)
	if err != nil {

		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var count int

	err = db.Mconn.GetOne("select count(id) from bugs where eid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.IsUse())
		return
	}
	getaritclesql := "delete from environment where id=?"

	_, err = db.Mconn.Update(getaritclesql, id)
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
		Classify: "env",
		Action:   "delete",
	}

	delete(cache.CacheEnvEid, cache.CacheEidEnv[cache.EnvId(eid)])
	delete(cache.CacheEidEnv, cache.EnvId(eid))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
