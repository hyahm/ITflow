package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
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
	rows, err := db.Mconn.GetRows("select id,name from environment")
	if err != nil {
		golog.Error(err)
		w.Write(el.ErrorE(err))
		return
	}
	for rows.Next() {
		e := &env.Env{}
		err = rows.Scan(&e.Id, &e.EnvName)
		if err != nil {
			golog.Error(err)
			continue
		}
		el.Elist = append(el.Elist, e)
	}

	w.Write(el.Marshal())
	return

}

func AddEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	envname := r.FormValue("name")

	getaritclesql := "insert into environment(name) values(?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(getaritclesql, envname)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 添加缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func UpdateEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	er := xmux.GetData(r).Data.(*env.Env)

	getaritclesql := "update environment set name=? where id=?"

	_, err := db.Mconn.Update(getaritclesql, er.EnvName, er.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志

	// 更新缓存
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

	delete(cache.CacheEnvEid, cache.CacheEidEnv[cache.EnvId(eid)])
	delete(cache.CacheEidEnv, cache.EnvId(eid))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
