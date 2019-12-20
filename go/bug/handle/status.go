package handle

import (
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"itflow/db"
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
	Code       int       `json:"code"`
}

func StatusList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
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
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
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

func StatusAdd(w http.ResponseWriter, r *http.Request) {
	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
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
			golog.Error(err.Error())
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
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	s := &status{}
	err = json.Unmarshal(ss, s)

	errorcode.Id, err = db.Mconn.Insert("insert into status(name) values(?)", s.Name)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "status",
	}
	err = il.Add(
		nickname, errorcode.Id, s.Name)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
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
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")
	sid, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err.Error())
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
			golog.Error(err.Error())
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
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&bcount)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	if bcount > 0 {
		golog.Error("sid:%d 删除失败", sid)
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
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 默认值

	if bugconfig.CacheDefault["status"] == int64(sid) {
		bugconfig.CacheDefault["status"] = 0
		_, err = db.Mconn.Update("update defaultvalue set status=0 ")
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorE(err))
			return
		}
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "status",
	}
	err = il.Del(
		nickname, sid)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
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

func StatusUpdate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
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
			golog.Error(err.Error())
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
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}
	s := &status{}
	err = json.Unmarshal(ss, s)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	_, err = db.Mconn.Update("update status set name=? where id=?", s.Name, s.Id)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "status",
	}
	err = il.Update(
		nickname, s.Id, s.Name)
	if err != nil {
		golog.Error(err.Error())
		w.Write(errorcode.ErrorE(err))
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

func StatusGroupName(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err.Error())
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
