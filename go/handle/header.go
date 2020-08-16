package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/response"
	"itflow/model"
	network "itflow/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func HeaderList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &network.List_headers{}

	gsql := "select id,name,hhids,remark from header"
	rows, err := db.Mconn.GetRows(gsql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		// 查询
		one := &network.Data_header{}
		var hs string
		rows.Scan(&one.Id, &one.Name, &hs, &one.Remark)
		if hs != "" {
			hrow, err := db.Mconn.GetRows(fmt.Sprintf("select id,k,v from headerlist where id in (%v)", hs))
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}
			for hrow.Next() {
				hone := &network.Table_headerlist{}
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

func HeaderAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	// 管理员
	if cache.CacheNickNameUid[nickname] != cache.SUPERID {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	data := xmux.GetData(r).Data.(*model.Data_header)

	idstr := make([]string, 0)
	for _, v := range data.Hhids {
		id, err := db.Mconn.Insert("insert into headerlist(k,v) values(?,?)", v.Key, v.Value)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		idstr = append(idstr, fmt.Sprintf("%v", id))
	}
	ids := strings.Join(idstr, ",")
	gsql := "insert into header(name,hhids,remark) values(?,?,?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(gsql, data.Name, ids, data.Remark)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志

	// 添加缓存
	cache.CacheHidHeader[errorcode.Id] = data.Name
	cache.CacheHeaderHid[data.Name] = errorcode.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func HeaderDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 管理员
	if cache.CacheNickNameUid[nickname] != cache.SUPERID {
		w.Write(errorcode.Error("没有权限"))
		return
	}
	// 查看这个header 是否有文档在用
	var count int
	err = db.Mconn.GetOne(" select count(id) from apilist where hid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 如果在使用，返回错误
	if count > 0 {
		w.Write(errorcode.Error("没有请求头"))
		return
	}
	// 先要删除子header
	var hids string
	err = db.Mconn.GetOne("select hhids from header where id=?", id).Scan(&hids)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 不为空就删
	if hids != "" {
		_, err = db.Mconn.Update(fmt.Sprintf("delete from headerlist where id in (%v)", hids))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}
	// 删除header
	_, err = db.Mconn.Update("delete from header where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志

	// 删除缓存
	delete(cache.CacheHeaderHid, cache.CacheHidHeader[int64(id32)])
	delete(cache.CacheHidHeader, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func HeaderUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	data := xmux.GetData(r).Data.(*model.Data_header)
	// 管理员
	if cache.CacheNickNameUid[nickname] != cache.SUPERID {
		w.Write(errorcode.Error("没有权限"))
		return
	}

	// 原来的header
	var oldheadids string
	err := db.Mconn.GetOne("select hhids from header where id=?", data.Id).Scan(&oldheadids)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 需要删除的hhis
	delhhids := strings.Split(oldheadids, ",")
	// 先修改header list
	idstr := make([]string, 0)
	for _, v := range data.Hhids {
		// 如果id > 0 就修改，
		if v.Id > 0 {
			_, err = db.Mconn.Update("update headerlist set k=?,v=? where id=?", v.Key, v.Value, v.Id)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
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
			hl := &network.Table_headerlist{
				Key:   v.Key,
				Value: v.Value,
			}
			//否则就添加,id也要返回
			hl.Id, err = db.Mconn.Insert("insert into headerlist(k,v) values(?,?)", v.Key, v.Value)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}
			cache.CacheHidHeader[hl.Id] = data.Name
			errorcode.HeaderList = append(errorcode.HeaderList, hl)
		}

	}

	// 删除多余的
	if len(delhhids) > 0 {
		_, err = db.Mconn.Update(fmt.Sprintf("delete from headerlist where id in (%s)", strings.Join(delhhids, ",")))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	// 修改header
	hids := strings.Join(idstr, ",")
	_, err = db.Mconn.Update("update header set name=?,hhids=?,remark=? where id=?", data.Name, hids, data.Remark, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志

	delete(cache.CacheHeaderHid, cache.CacheHidHeader[data.Id])
	cache.CacheHidHeader[data.Id] = data.Name
	cache.CacheHeaderHid[data.Name] = data.Id
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type headerstruct struct {
	Headers []string `json:"headers"`
	Code    int      `json:"code"`
}

func HeaderGet(w http.ResponseWriter, r *http.Request) {

	data := &headerstruct{}

	for _, v := range cache.CacheHidHeader {
		data.Headers = append(data.Headers, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
