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

func PositionGet(w http.ResponseWriter, r *http.Request) {
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
		data := &model.List_jobs{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("position", nickname, conn)
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
		rows, err := conn.GetRows("select id,name,level,hypo from jobs")
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			var jid int64
			one := &model.Table_jobs{}
			rows.Scan(&one.Id, &one.Name, &one.Level, &jid)
			one.Hypo = bugconfig.CacheJidJobname[jid]
			data.Positions = append(data.Positions, one)
		}

		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func PositionAdd(w http.ResponseWriter, r *http.Request) {
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
		data := &model.Data_jobs{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("position", nickname, conn)
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
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		err = json.Unmarshal(respbyte, data)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		// 如果不存在管理层名，就参数错误
		var hid int64
		var ok bool
		if data.Hyponame != "" {
			if hid, ok = bugconfig.CacheJobnameJid[data.Hyponame]; !ok {
				w.Write(errorcode.ErrorParams())
				return
			}
		}

		errorcode.Id, err = conn.InsertWithID("insert into jobs(name,level,hypo) value(?,?,?)", data.Name, data.Level, hid)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "position",
		}
		err = il.Add(
			nickname, errorcode.Id, data.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		//更新缓存
		bugconfig.CacheJobnameJid[data.Name] = errorcode.Id
		bugconfig.CacheJidJobname[errorcode.Id] = data.Name
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func PositionDel(w http.ResponseWriter, r *http.Request) {
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
		id32, err := strconv.Atoi(id)
		if err != nil {
			w.Write(errorcode.ErrorParams())
			return
		}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("position", nickname, conn)
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
		// 如果这个职位被使用了，不允许被删除
		var count int
		err = conn.GetOne("select count(id) from user where jid=?", id).Scan(&count)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if count > 0 {
			w.Write(errorcode.ErrorHasPosition())
			return
		}
		// 是否被所属使用
		err = conn.GetOne("select count(id) from jobs where hypo=?", id).Scan(&count)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if count > 0 {
			w.Write(errorcode.ErrorHasHypo())
			return
		}
		gsql := "delete from jobs where id=?"

		_, err = conn.Update(gsql, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "position",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 删除缓存
		delete(bugconfig.CacheJobnameJid, bugconfig.CacheJidJobname[int64(id32)])
		delete(bugconfig.CacheJidJobname, int64(id32))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func PositionUpdate(w http.ResponseWriter, r *http.Request) {
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
		data := &model.Update_jobs{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("position", nickname, conn)
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
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		err = json.Unmarshal(respbyte, data)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var hid int64
		var ok bool
		// 不存在这个职位的hypo，直接返回参数错误
		if data.Hypo != "" {
			if hid, ok = bugconfig.CacheJobnameJid[data.Hypo]; !ok {
				w.Write(errorcode.ErrorParams())
				return
			}
		}

		_, err = conn.Update("update jobs set name=?,level=?,hypo=? where id=?", data.Name, data.Level, hid, data.Id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "position",
		}
		err = il.Update(
			nickname, data.Id, data.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 更新缓存
		delete(bugconfig.CacheJobnameJid, bugconfig.CacheJidJobname[data.Id])
		bugconfig.CacheJidJobname[data.Id] = data.Name
		bugconfig.CacheJobnameJid[data.Name] = data.Id

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type hypos struct {
	Hypos []string `json:"hypos"`
	Code  int      `json:"statuscode"`
}

func GetHypos(w http.ResponseWriter, r *http.Request) {
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
		data := &hypos{}
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] != bugconfig.SUPERID {
			w.Write(errorcode.ErrorNoPermission())
			return
		}

		rows, err := conn.GetRows("select name  from jobs where level=1")
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			var name string
			rows.Scan(&name)
			data.Hypos = append(data.Hypos, name)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type positions struct {
	Positions []string `json:"positions"`
	Code      int      `json:"statuscode"`
}

func GetPositions(w http.ResponseWriter, r *http.Request) {
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
		data := &positions{}

		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			for _, v := range bugconfig.CacheJidJobname {
				data.Positions = append(data.Positions, v)
			}
		} else {
			var name string
			err = conn.GetOne("select name from jobs where hypo=?", bugconfig.CacheUidJid[bugconfig.CacheNickNameUid[nickname]]).Scan(&name)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			data.Positions = append(data.Positions, name)
		}

		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
