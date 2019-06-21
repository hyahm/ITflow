package handle

import (
	"bug/bugconfig"
	"bug/buglog"
	"bug/model"
	"encoding/json"
	"fmt"
	"galog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func HeaderList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &model.List_headers{}

		gsql := "select id,name,hhids,remark from header"
		rows, err := conn.GetRows(gsql)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			// 查询
			one := &model.Data_header{}
			var hs string
			rows.Scan(&one.Id, &one.Name, &hs, &one.Remark)
			if hs != "" {
				hrow, err := conn.GetRows(fmt.Sprintf("select id,k,v from headerlist where id in (%v)", hs))
				if err != nil {
					galog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
				for hrow.Next() {
					hone := &model.Table_headerlist{}
					hrow.Scan(&hone.Id, &hone.Key, &hone.Value)
					one.Hhids = append(one.Hhids, hone)
				}

			}
			data.Headers = append(data.Headers, one)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func HeaderAdd(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] != bugconfig.SUPERID {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		data := &model.Data_header{}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		err = json.Unmarshal(respbyte, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		idstr := make([]string, 0)
		for _, v := range data.Hhids {
			id, err := conn.InsertWithID("insert into headerlist(k,v) values(?,?)", v.Key, v.Value)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			idstr = append(idstr, fmt.Sprintf("%v", id))
		}
		ids := strings.Join(idstr, ",")
		gsql := "insert into header(name,hhids,remark) values(?,?,?)"
		errorcode.Id, err = conn.Insert(gsql, data.Name, ids, data.Remark)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "header",
		}
		err = il.Add(
			nickname, errorcode.Id, data.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 添加缓存
		bugconfig.CacheHidHeader[errorcode.Id] = data.Name
		bugconfig.CacheHeaderHid[data.Name] = errorcode.Id
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func HeaderDel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodGet {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
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
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] != bugconfig.SUPERID {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		// 查看这个header 是否有文档在用
		var count int
		err = conn.GetOne(" select count(id) from apilist where hid=?", id).Scan(&count)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 如果在使用，返回错误
		if count > 0 {
			w.Write(errorcode.ErrorHasHeader())
			return
		}
		// 先要删除子header
		var hids string
		err = conn.GetOne("select hhids from header where id=?", id).Scan(&hids)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 不为空就删
		if hids != "" {
			_, err = conn.Update(fmt.Sprintf("delete from headerlist where id in (%v)", hids))
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}
		// 删除header
		_, err = conn.Update("delete from header where id=?", id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "header",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 删除缓存
		delete(bugconfig.CacheHeaderHid, bugconfig.CacheHidHeader[int64(id32)])
		delete(bugconfig.CacheHidHeader, int64(id32))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func HeaderUpdate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &model.Data_header{}
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] != bugconfig.SUPERID {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		err = json.Unmarshal(respbyte, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		// 原来的header
		var oldheadids string
		err = conn.GetOne("select hhids from header where id=?", data.Id).Scan(&oldheadids)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 需要删除的hhis
		delhhids := strings.Split(oldheadids, ",")
		// 先修改header list
		idstr := make([]string, 0)
		for _, v := range data.Hhids {
			// 如果id > 0 就修改，
			if v.Id > 0 {
				_, err = conn.Update("update headerlist set k=?,v=? where id=?", v.Key, v.Value, v.Id)
				if err != nil {
					galog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
				//update的，则不需要再old里面删除
				tmplist := make([]string, 0)
				for i, hv := range delhhids {
					if hv == fmt.Sprintf("%v", v.Id) {
						tmplist = append(tmplist, delhhids[0:i]...)
						tmplist = append(tmplist, delhhids[i+1:]...)
						delhhids = tmplist
					}
				}
				idstr = append(idstr, fmt.Sprintf("%d", v.Id))
				errorcode.HeaderList = append(errorcode.HeaderList, v)
			} else {
				hl := &model.Table_headerlist{
					Key:   v.Key,
					Value: v.Value,
				}
				//否则就添加,id也要返回
				hl.Id, err = conn.InsertWithID("insert into headerlist(k,v) values(?,?)", v.Key, v.Value)
				if err != nil {
					galog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
				bugconfig.CacheHidHeader[hl.Id] = data.Name
				errorcode.HeaderList = append(errorcode.HeaderList, hl)
			}

		}

		// 删除多余的
		if len(delhhids) > 0 {
			_, err = conn.Update(fmt.Sprintf("delete from headerlist where id in (%s)", strings.Join(delhhids, ",")))
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		// 修改header
		hids := strings.Join(idstr, ",")
		_, err = conn.Update("update header set name=?,hhids=?,remark=? where id=?", data.Name, hids, data.Remark, data.Id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "header",
		}
		err = il.Update(
			nickname, data.Id, data.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		delete(bugconfig.CacheHeaderHid, bugconfig.CacheHidHeader[data.Id])
		bugconfig.CacheHidHeader[data.Id] = data.Name
		bugconfig.CacheHeaderHid[data.Name] = data.Id
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type headerstruct struct {
	Headers []string `json:"headers"`
	Code    int      `json:"statuscode"`
}

func HeaderGet(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &headerstruct{}

		for _, v := range bugconfig.CacheHidHeader {
			data.Headers = append(data.Headers, v)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
