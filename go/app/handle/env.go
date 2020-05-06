package handle

import (
	"encoding/json"
	"io/ioutil"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/network/datalog"
	"itflow/network/response"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
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

	env := &envlist{}

	// 管理员

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
	bugconfig.CacheEidName[errorcode.Id] = envname
	bugconfig.CacheEnvNameEid[envname] = errorcode.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func UpdateEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	er := &envrow{}

	bpr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(bpr, er)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	getaritclesql := "update environment set envname=? where id=?"

	_, err = db.Mconn.Update(getaritclesql, er.EnvName, er.Id)
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
	delete(bugconfig.CacheEnvNameEid, bugconfig.CacheEidName[int64(er.Id)])
	bugconfig.CacheEidName[int64(er.Id)] = er.EnvName
	bugconfig.CacheEnvNameEid[er.EnvName] = int64(er.Id)
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

	row, err := db.Mconn.GetOne("select count(id) from bugs where eid=?", id)
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
		w.Write(errorcode.Error("存在此env"))
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

	delete(bugconfig.CacheEnvNameEid, bugconfig.CacheEidName[int64(eid)])
	delete(bugconfig.CacheEidName, int64(eid))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
