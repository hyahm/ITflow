package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	network "itflow/model"
	"net/http"
	"strconv"
	"strings"

	"itflow/internal/bug"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func StatusList(w http.ResponseWriter, r *http.Request) {

	ls := &bug.ListStatus{}
	for k, v := range cache.CacheSidStatus {
		one := &bug.Status{}
		one.Id = k
		one.Name = v
		ls.StatusList = append(ls.StatusList, one)
	}

	send, _ := json.Marshal(ls)
	w.Write(send)
	return

}

func StatusAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	var err error
	s := xmux.GetData(r).Data.(*bug.Status)
	errorcode.Id, err = db.Mconn.Insert("insert into status(name) values(?)", s.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ConnectMysqlFail())
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: xmux.GetData(r).Get("nickname").(string),
		Classify: "status",
		Action:   "add",
	}

	// 更新缓存
	cache.CacheSidStatus[errorcode.Id] = s.Name
	cache.CacheStatusSid[s.Name] = errorcode.Id

	w.Write(errorcode.Success())
	return

}

func StatusRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	sid, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 如果bug有这个状态，就不能修改
	var bcount int
	err = db.Mconn.GetOne("select count(id) from bugs where sid=?", sid).Scan(&bcount)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if bcount > 0 {
		golog.Errorf("sid:%d 删除失败", sid)
		w.Write(errorcode.Error("还有bug"))
		return
	}

	//如果状态组存在也无法删除

	var hasgroup bool
	for _, ids := range cache.CacheSgidGroup {
		for _, v := range strings.Split(ids, ",") {
			if v == id {
				hasgroup = true
				break
			}
		}
		if hasgroup {
			w.Write(errorcode.Error("还有group"))
		}
	}

	_, err = db.Mconn.Update("delete from  status where id=?", sid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 默认值

	if cache.CacheDefault["status"] == int64(sid) {
		cache.CacheDefault["status"] = 0
		_, err = db.Mconn.Update("update defaultvalue set status=0 ")
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: xmux.GetData(r).Get("nickname").(string),
		Classify: "status",
		Action:   "delete",
	}

	// 更新缓存
	// 获取status的索引

	delete(cache.CacheStatusSid, cache.CacheSidStatus[int64(sid)])
	delete(cache.CacheSidStatus, int64(sid))

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	s := xmux.GetData(r).Data.(*bug.Status)
	_, err := db.Mconn.Update("update status set name=? where id=?", s.Name, s.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: xmux.GetData(r).Get("nickname").(string),
		Classify: "status",
		Action:   "update",
	}

	// 更新缓存

	delete(cache.CacheStatusSid, cache.CacheSidStatus[s.Id])
	cache.CacheSidStatus[s.Id] = s.Name
	cache.CacheStatusSid[s.Name] = s.Id

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func StatusGroupName(w http.ResponseWriter, r *http.Request) {

	sl := &network.List_StatusName{}
	for _, v := range cache.CacheSgidGroup {
		sl.StatusList = append(sl.StatusList, v)
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}
