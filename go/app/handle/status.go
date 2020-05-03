package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"itflow/app/asset"
	"itflow/app/bugconfig"
	"itflow/app/model"
	"itflow/db"
	"itflow/model/datalog"
	"itflow/model/response"
	"net/http"
	"strconv"
	"strings"

	"itflow/model/bug"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func StatusList(w http.ResponseWriter, r *http.Request) {

	ls := &bug.ListStatus{}
	for k, v := range bugconfig.CacheSidStatus {
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
		w.Write(errorcode.ErrorE(err))
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
	bugconfig.CacheSidStatus[errorcode.Id] = s.Name
	bugconfig.CacheStatusSid[s.Name] = errorcode.Id

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusRemove(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")
	sid, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("status", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 如果bug有这个状态，就不能修改
	var bcount int
	row, err := db.Mconn.GetOne("select count(id) from bugs where sid=?", sid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&bcount)
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
	for _, ids := range bugconfig.CacheSgidGroup {
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

	if bugconfig.CacheDefault["status"] == int64(sid) {
		bugconfig.CacheDefault["status"] = 0
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
		Username: nickname,
		Classify: "status",
		Action:   "delete",
	}

	// 更新缓存
	// 获取status的索引

	delete(bugconfig.CacheStatusSid, bugconfig.CacheSidStatus[int64(sid)])
	delete(bugconfig.CacheSidStatus, int64(sid))

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("status", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	ss, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	s := &bug.Status{}
	fmt.Println(string(ss))
	err = json.Unmarshal(ss, s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	_, err = db.Mconn.Update("update status set name=? where id=?", s.Name, s.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "status",
		Action:   "update",
	}

	// 更新缓存

	delete(bugconfig.CacheStatusSid, bugconfig.CacheSidStatus[s.Id])
	bugconfig.CacheSidStatus[s.Id] = s.Name
	bugconfig.CacheStatusSid[s.Name] = s.Id

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func StatusGroupName(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	sl := &model.List_StatusName{}
	for _, v := range bugconfig.CacheSgidGroup {
		sl.StatusList = append(sl.StatusList, v)
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}
