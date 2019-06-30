package handle

import (
	"bug/bugconfig"
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"net/http"
)

type DefaultValue struct {
	DefaultStatus string `json:"defaultstatus"`
	Important     string `json:"defaultimportant"`
	Level         string `json:"defaultlevel"`
	Code          int    `json:"statuscode"`
}

func DefaultStatus(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
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
		sl := &DefaultValue{}
		//如果是管理员的话,所有的都可以
		sl.DefaultStatus = bugconfig.CacheSidStatus[bugconfig.CacheDefault["status"]]
		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func DefaultSave(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
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
		sl := &DefaultValue{}
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		err = json.Unmarshal(bytedata, sl)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var sid int64
		var iid int64
		var lid int64
		var ok bool
		if sid, ok = bugconfig.CacheStatusSid[sl.DefaultStatus]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		if iid, ok = bugconfig.CacheImportantIid[sl.Important]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		if lid, ok = bugconfig.CacheLevelLid[sl.Level]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		//修改字段
		_, err = conn.Update("update defaultvalue set status=?, important=?,level=?", sid, iid, lid)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
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
	w.WriteHeader(http.StatusNotFound)
}

type defaultImportant struct {
	DefaultImportant string `json:"defaultimportant"`
	Code             int    `json:"statuscode"`
}

func DefaultImportant(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
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
		data := &defaultImportant{}
		data.DefaultImportant = bugconfig.CacheIidImportant[bugconfig.CacheDefault["important"]]
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type defaultLevel struct {
	DefaultLevel string `json:"defaultlevel"`
	Code         int    `json:"statuscode"`
}

func DefaultLevel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
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
		data := &defaultLevel{}
		data.DefaultLevel = bugconfig.CacheLidLevel[bugconfig.CacheDefault["level"]]
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
