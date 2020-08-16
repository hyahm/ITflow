package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/response"
	"itflow/model"
	network "itflow/model"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func PositionGet(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &network.List_jobs{
		Positions: make([]*network.Table_jobs, 0),
	}

	rows, err := db.Mconn.GetRows("select id,name,level,hypo from jobs")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var jid int64
		one := &network.Table_jobs{}
		rows.Scan(&one.Id, &one.Name, &one.Level, &jid)
		one.Hypo = cache.CacheJidJobname[jid]
		data.Positions = append(data.Positions, one)
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func PositionAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Data_jobs)

	// 如果不存在管理层名，就参数错误
	var hid int64
	var ok bool
	var err error
	if data.Hyponame != "" {
		if hid, ok = cache.CacheJobnameJid[data.Hyponame]; !ok {
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	errorcode.Id, err = db.Mconn.Insert("insert into jobs(name,level,hypo) value(?,?,?)", data.Name, data.Level, hid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	//更新缓存
	cache.CacheJobnameJid[data.Name] = errorcode.Id
	cache.CacheJidJobname[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func PositionDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 如果这个职位被使用了，不允许被删除
	var count int
	err = db.Mconn.GetOne("select count(id) from user where jid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 是否被所属使用
	err = db.Mconn.GetOne("select count(id) from jobs where hypo=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("没有职位"))
		return
	}
	gsql := "delete from jobs where id=?"

	_, err = db.Mconn.Update(gsql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 删除缓存
	delete(cache.CacheJobnameJid, cache.CacheJidJobname[int64(id32)])
	delete(cache.CacheJidJobname, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func PositionUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Update_jobs)

	var hid int64
	var ok bool
	// 不存在这个职位的hypo，直接返回参数错误
	if data.Hypo != "" {
		if hid, ok = cache.CacheJobnameJid[data.Hypo]; !ok {
			w.Write(errorcode.Error("没有职位"))
			return
		}
	}

	_, err := db.Mconn.Update("update jobs set name=?,level=?,hypo=? where id=?", data.Name, data.Level, hid, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 更新缓存
	delete(cache.CacheJobnameJid, cache.CacheJidJobname[data.Id])
	cache.CacheJidJobname[data.Id] = data.Name
	cache.CacheJobnameJid[data.Name] = data.Id

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type hypos struct {
	Hypos []string `json:"hypos"`
	Code  int      `json:"code"`
}

func GetHypos(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	data := &hypos{
		Hypos: make([]string, 0),
	}
	// 管理员
	if cache.CacheNickNameUid[nickname] != cache.SUPERID {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	rows, err := db.Mconn.GetRows("select name  from jobs where level=1")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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

type positions struct {
	Positions []string `json:"positions"`
	Code      int      `json:"code"`
}

func GetPositions(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &positions{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	if cache.CacheNickNameUid[nickname] == cache.SUPERID {
		for _, v := range cache.CacheJidJobname {
			data.Positions = append(data.Positions, v)
		}
	} else {
		var name string
		err := db.Mconn.GetOne("select name from jobs where hypo=?", cache.CacheUidJid[cache.CacheNickNameUid[nickname]]).Scan(&name)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

		data.Positions = append(data.Positions, name)
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
