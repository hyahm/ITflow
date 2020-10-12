package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/env"
	"itflow/internal/perm"
	"itflow/internal/response"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func EnvList(w http.ResponseWriter, r *http.Request) {

	el := &env.Envlist{
		Elist: make([]*env.Env, 0),
	}
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Select {
		w.Write(el.Error("no perm"))
		return
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
			golog.Info(err)
			continue
		}
		el.Elist = append(el.Elist, e)
	}
	rows.Close()
	w.Write(el.Marshal())
	return

}

func AddEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Insert {
		w.Write(errorcode.Error("no perm"))
		return
	}
	envname := r.FormValue("name")

	getaritclesql := "insert into environment(name) values(?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(getaritclesql, envname)
	if err != nil {
		golog.Error(err)
		if err.(*mysql.MySQLError).Number == 1062 {
			w.Write(errorcode.ErrorE(db.DuplicateErr))
			return
		}
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
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Update {
		w.Write(errorcode.Error("no perm"))
		return
	}
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
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Delete {
		w.Write(errorcode.Error("no perm"))
		return
	}
	id := r.FormValue("id")

	var count int

	err := db.Mconn.GetOne("select count(id) from bugs where eid=?", id).Scan(&count)
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

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
