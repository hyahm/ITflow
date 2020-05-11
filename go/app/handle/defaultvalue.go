package handle

import (
	"encoding/json"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/network/response"
	"net/http"

	"itflow/network/defaults"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	sl := &defaults.DefaultValue{}
	//如果是管理员的话,所有的都可以
	sl.DefaultStatus = bugconfig.CacheSidStatus[bugconfig.CacheDefault["status"]]
	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetData(r).Data.(*defaults.DefaultValue)

	var sid int64
	var iid int64
	var lid int64
	var ok bool
	if sid, ok = bugconfig.CacheStatusSid[sl.DefaultStatus]; !ok {
		w.Write(errorcode.Error("没有找到status "))
		return
	}
	if iid, ok = bugconfig.CacheImportantIid[sl.Important]; !ok {
		w.Write(errorcode.Error("没有找到important "))
		return
	}
	if lid, ok = bugconfig.CacheLevelLid[sl.Level]; !ok {
		w.Write(errorcode.Error("没有找到level "))
		return
	}
	//修改字段
	_, err := db.Mconn.Update("update defaultvalue set status=?, important=?,level=?", sid, iid, lid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 更新缓存
	bugconfig.CacheDefault["status"] = sid
	bugconfig.CacheDefault["important"] = iid
	bugconfig.CacheDefault["level"] = lid
	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}

type defaultImportant struct {
	DefaultImportant string `json:"defaultimportant"`
	Code             int    `json:"code"`
}

func DefaultImportant(w http.ResponseWriter, r *http.Request) {

	data := &defaultImportant{}
	data.DefaultImportant = bugconfig.CacheIidImportant[bugconfig.CacheDefault["important"]]
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

type defaultLevel struct {
	DefaultLevel string `json:"defaultlevel"`
	Code         int    `json:"code"`
}

func DefaultLevel(w http.ResponseWriter, r *http.Request) {

	data := &defaultLevel{}
	data.DefaultLevel = bugconfig.CacheLidLevel[bugconfig.CacheDefault["level"]]
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
