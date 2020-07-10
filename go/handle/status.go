package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/status"
	network "itflow/model"
	"net/http"
	"strconv"
	"strings"

	"itflow/internal/bug"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func StatusList(w http.ResponseWriter, r *http.Request) {

	w.Write(status.StatusList())
	return

}

func StatusAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	var err error
	s := xmux.GetData(r).Data.(*bug.ReqStatus)
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
	cache.CacheSidStatus[cache.StatusId(errorcode.Id)] = s.Name
	cache.CacheStatusSid[s.Name] = cache.StatusId(errorcode.Id)

	w.Write(errorcode.Success())
	return

}

func StatusRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	sid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if cache.DefaultSid.ToInt64() == sid {
		w.Write(errorcode.Error("this status is set to default, can not remove. you can change to other status and delete"))
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
		w.Write(errorcode.IsUse())
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

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: xmux.GetData(r).Get("nickname").(string),
		Classify: "status",
		Action:   "delete",
	}

	// 更新缓存
	// 获取status的索引

	delete(cache.CacheStatusSid, cache.CacheSidStatus[cache.StatusId(sid)])
	delete(cache.CacheSidStatus, cache.StatusId(sid))

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	s := xmux.GetData(r).Data.(*bug.ReqStatus)
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
