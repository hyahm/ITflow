package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/status"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	if data.Name == "" {
		w.Write(errorcode.Error("名称不能为空"))
		return
	}
	// 重新排序
	ids := make([]string, 0)
	for _, v := range data.StatusList {
		ids = append(ids, fmt.Sprintf("%d", cache.CacheStatusSid[v]))
	}

	ss := strings.Join(ids, ",")
	var err error
	isql := "insert into statusgroup(name,sids) values(?,?)"
	errorcode.Id, err = db.Mconn.Insert(isql, data.Name, ss)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip: r.RemoteAddr,

		Classify: "statusgroup",
	}

	// 添加缓存
	cache.CacheSgidGroup[errorcode.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	if data.Name == "" {
		w.Write(errorcode.Error("名称不能为空"))
		return
	}

	isql := "update statusgroup set name =?,sids=? where id = ?"
	_, err := db.Mconn.Update(isql, data.Name, data.StatusList.ToStore(), data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "buggroup",
		Action:   "update",
	}

	cache.CacheSgidGroup[data.Id] = data.Name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusGroupList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &departmentList{}
	s := "select id,name,sids from statusgroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var ids cache.StoreStatusId
		one := &department{}
		rows.Scan(&one.Id, &one.Name, &ids)
		one.BugstatusList = ids.ToShow()
		// for _, v := range strings.Split(ids, ",") {
		// 	id, err := strconv.Atoi(v)
		// 	if err != nil {
		// 		log.Println(err)
		// 		continue
		// 	}

		// 	one.BugstatusList = append(one.BugstatusList, cache.CacheSidStatus[cache.StatusId(id)])
		// }
		data.DepartmentList = append(data.DepartmentList, one)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func DeleteStatusGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}

	ssql := "select count(id) from user where bugsid=?"
	var count int
	err = db.Mconn.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("有人再使用"))
		return
	}
	isql := "delete from  statusgroup where id = ?"
	_, err = db.Mconn.Update(isql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	golog.Info(nickname)
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "buggroup",
		Action:   "delete",
	}

	//更新缓存
	delete(cache.CacheSgidGroup, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GetStatusGroupName(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := &struct {
		Names []string `json:"names"`
		Code  int      `json:"code"`
	}{
		Names: make([]string, 0),
	}
	s := "select name from statusgroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		var name string
		rows.Scan(&name)
		data.Names = append(data.Names, name)

	}
	send, _ := json.Marshal(data)
	golog.Info(string(send))
	w.Write(send)
	return

}
