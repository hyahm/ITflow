package handle

import (
	"encoding/json"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"itflow/db"
	"itflow/model/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

func ImportantGet(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.List_importants{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("important", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	for k, v := range bugconfig.CacheIidImportant {
		one := &model.Table_importants{}
		one.Id = k
		one.Name = v
		data.ImportantList = append(data.ImportantList, one)
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func ImportantAdd(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Data_importants{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("important", nickname)
		if err != nil {
			golog.Error(err)
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
	errorcode.Id, err = db.Mconn.Insert("insert into importants(name) value(?)", data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "important",
	}
	err = il.Add(
		nickname, errorcode.Id, data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	//更新缓存
	bugconfig.CacheImportantIid[data.Name] = errorcode.Id
	bugconfig.CacheIidImportant[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ImportantDel(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("important", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	// 判断是否有bug在使用
	var count int
	row, err := db.Mconn.GetOne("select count(id) from bugs where iid=?", id32)
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
		w.Write(errorcode.Error("不存在bug"))
		return
	}
	// 是否设定为了默认值
	if bugconfig.CacheDefault["important"] == int64(id32) {
		w.Write(errorcode.Error("没有设定默认值"))
		return
	}
	gsql := "delete from importants where id=?"

	_, err = db.Mconn.Update(gsql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "important",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 删除缓存
	delete(bugconfig.CacheImportantIid, bugconfig.CacheIidImportant[int64(id32)])
	delete(bugconfig.CacheIidImportant, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ImportantUpdate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &model.Update_importants{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("important", nickname)
		if err != nil {
			golog.Error(err)
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

	gsql := "update importants set name=? where id=?"

	_, err = db.Mconn.Update(gsql, data.Name, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "important",
	}
	err = il.Update(
		nickname, data.Id, data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 删除strings key
	delete(bugconfig.CacheImportantIid, bugconfig.CacheIidImportant[data.Id])
	bugconfig.CacheIidImportant[data.Id] = data.Name
	bugconfig.CacheImportantIid[data.Name] = data.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type importantslist struct {
	Importants []string `json:"importants"`
	Code       int      `json:"code"`
}

func GetImportants(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &importantslist{}
	for _, v := range bugconfig.CacheIidImportant {
		data.Importants = append(data.Importants, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
