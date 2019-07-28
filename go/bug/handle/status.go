package handle

import (
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type status struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type listStatus struct {
	StatusList []*status `json:"statuslist"`
	Code       int       `json:"statuscode"`
}

func StatusList(w http.ResponseWriter, r *http.Request) {
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
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("status", nickname, conn)
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
		ls := &listStatus{}
		for k, v := range bugconfig.CacheSidStatus {
			one := &status{}
			one.Id = k
			one.Name = v
			ls.StatusList = append(ls.StatusList, one)
		}

		send, _ := json.Marshal(ls)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func StatusAdd(w http.ResponseWriter, r *http.Request) {
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
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("status", nickname, conn)
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
		ss, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		s := &status{}
		err = json.Unmarshal(ss, s)

		errorcode.Id, err = conn.InsertWithID("insert into status(name) values(?)", s.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "status",
		}
		err = il.Add(
			nickname, errorcode.Id, s.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 更新缓存
		bugconfig.CacheSidStatus[errorcode.Id] = s.Name
		bugconfig.CacheStatusSid[s.Name] = errorcode.Id

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func StatusRemove(w http.ResponseWriter, r *http.Request) {
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

		id := r.FormValue("id")
		sid, err := strconv.Atoi(id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("status", nickname, conn)
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
		// 如果bug有这个状态，就不能修改
		var bcount int
		err = conn.GetOne("select count(id) from bugs where sid=?", sid).Scan(&bcount)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if bcount > 0 {
			golog.Error("sid:%d 删除失败", sid)
			w.Write(errorcode.ErrorHasBug())
			return
		}

		//如果状态组存在也无法删除

		var hasgroup bool
		for _, ids := range bugconfig.CacheSgidGroup {
			for _, v := range strings.Split(ids, ",") {
				if v == id {
					hasgroup = true
					break
				}
			}
			if hasgroup {
				w.Write(errorcode.ErrorHasGroup())
			}
		}

		_, err = conn.Update("delete from  status where id=?", sid)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 默认值

		if bugconfig.CacheDefault["status"] == int64(sid) {
			bugconfig.CacheDefault["status"] = 0
			_, err = conn.Update("update defaultvalue set status=0 ")
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "status",
		}
		err = il.Del(
			nickname, sid)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 更新缓存
		// 获取status的索引

		delete(bugconfig.CacheStatusSid, bugconfig.CacheSidStatus[int64(sid)])
		delete(bugconfig.CacheSidStatus, int64(sid))

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {
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

		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("status", nickname, conn)
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
		ss, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		s := &status{}
		err = json.Unmarshal(ss, s)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}

		_, err = conn.Update("update status set name=? where id=?", s.Name, s.Id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "status",
		}
		err = il.Update(
			nickname, s.Id, s.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 更新缓存

		delete(bugconfig.CacheStatusSid, bugconfig.CacheSidStatus[s.Id])
		bugconfig.CacheSidStatus[s.Id] = s.Name
		bugconfig.CacheStatusSid[s.Name] = s.Id

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func StatusGroupName(w http.ResponseWriter, r *http.Request) {
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

		sl := &model.List_StatusName{}
		for _, v := range bugconfig.CacheSgidGroup {
			sl.StatusList = append(sl.StatusList, v)
		}

		send, _ := json.Marshal(sl)
		w.Write(send)
		return
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
