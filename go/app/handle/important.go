package handle

import (
	"encoding/json"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/model"
	network "itflow/model"
	"itflow/network/datalog"
	"itflow/network/response"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ImportantGet(w http.ResponseWriter, r *http.Request) {

	data := &network.List_importants{}

	for k, v := range bugconfig.CacheIidImportant {
		one := &network.Table_importants{}
		one.Id = k
		one.Name = v
		data.ImportantList = append(data.ImportantList, one)
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func ImportantAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Data_importants)

	var err error
	errorcode.Id, err = db.Mconn.Insert("insert into importants(name) value(?)", data.Name)
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
		Classify: "important",
		Action:   "add",
	}

	//更新缓存
	bugconfig.CacheImportantIid[data.Name] = errorcode.Id
	bugconfig.CacheIidImportant[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ImportantDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 判断是否有bug在使用
	var count int
	err = db.Mconn.GetOne("select count(id) from bugs where iid=?", id32).Scan(&count)
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
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "important",
		Action:   "delete",
	}

	// 删除缓存
	delete(bugconfig.CacheImportantIid, bugconfig.CacheIidImportant[int64(id32)])
	delete(bugconfig.CacheIidImportant, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ImportantUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Update_importants)
	gsql := "update importants set name=? where id=?"

	_, err := db.Mconn.Update(gsql, data.Name, data.Id)
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
		Classify: "important",
		Action:   "update",
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

	data := &importantslist{}
	for _, v := range bugconfig.CacheIidImportant {
		data.Importants = append(data.Importants, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
