package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/status"
	"itflow/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	// // 重新排序
	// ids := make([]string, 0)
	// for _, v := range data.StatusList {
	// 	ids = append(ids, fmt.Sprintf("%d", cache.CacheStatusSid[v]))
	// }

	// ss := strings.Join(ids, ",")

	isql := "insert into statusgroup(name,sids) values(?,?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(isql, data.Name, data.StatusList.ToStore())
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

func EditBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	// ssids := make([]string, 0)
	// for _, v := range data.StatusList {
	// 	// 没找到就是key
	// 	var sid int64
	// 	var ok bool
	// 	if sid, ok = cache.CacheStatusSid[v]; !ok {
	// 		w.Write(errorcode.Error("没有找到status"))
	// 		return
	// 	}
	// 	ssids = append(ssids, strconv.FormatInt(sid, 10))
	// }
	// ss := strings.Join(ssids, ",")
	isql := "update statusgroup set name =?,sids=? where id = ?"
	_, err := db.Mconn.Update(isql, data.Name, data.StatusList.ToStore(), data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info(db.Mconn.GetSql())
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

type department struct {
	Id            int64            `json:"id"`
	Name          string           `json:"name"`
	BugstatusList cache.StatusList `json:"bugstatuslist"`
}

type departmentList struct {
	DepartmentList []*department `json:"departmentlist"`
	Code           int           `json:"code"`
}

func BugGroupList(w http.ResponseWriter, r *http.Request) {

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

		// 	one.BugstatusList = append(one.BugstatusList, cache.CacheSidStatus[int64(id)])
		// }
		data.DepartmentList = append(data.DepartmentList, one)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func BugGroupDel(w http.ResponseWriter, r *http.Request) {

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
		w.Write(errorcode.Error("没有找到group"))
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

func GroupGet(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &model.Send_groups{
		GroupList: make([]*model.Get_groups, 0),
	}
	uid := xmux.GetData(r).Get("uid").(int64)
	if cache.SUPERID == uid {
		gsql := "select id,name,ids from usergroup"
		rows, err := db.Mconn.GetRows(gsql)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for rows.Next() {
			onegroup := &model.Get_groups{}
			var users string
			rows.Scan(&onegroup.Id, &onegroup.Name, &users)
			for _, v := range strings.Split(users, ",") {
				uid, _ := strconv.Atoi(v)
				onegroup.Users = append(onegroup.Users, cache.CacheUidRealName[int64(uid)])
			}
			data.GroupList = append(data.GroupList, onegroup)
		}
	} else {
		gsql := "select id,name,ids from usergroup where cuid=?"
		rows, err := db.Mconn.GetRows(gsql, uid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for rows.Next() {
			onegroup := &model.Get_groups{}
			var users string
			rows.Scan(&onegroup.Id, &onegroup.Name, &users)
			for _, v := range strings.Split(users, ",") {
				uid, _ := strconv.Atoi(v)
				onegroup.Users = append(onegroup.Users, cache.CacheUidRealName[int64(uid)])
			}
			data.GroupList = append(data.GroupList, onegroup)
		}
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func GroupAdd(w http.ResponseWriter, r *http.Request) {
	golog.Info("add..................")
	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	data := xmux.GetData(r).Data.(*model.Get_groups)

	ids := make([]string, 0)
	for _, v := range data.Users {
		var uid int64
		var ok bool
		if uid, ok = cache.CacheRealNameUid[v]; !ok {
			w.Write(errorcode.Errorf("没有此用户"))
			return
		}
		ids = append(ids, strconv.FormatInt(uid, 10))
	}
	gsql := "insert usergroup(name,ids,cuid) values(?,?,?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(gsql, data.Name, strings.Join(ids, ","), cache.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "usergroup",
		Action:   "add",
	}
	ug := &cache.UG{}
	ug.Gid = errorcode.Id
	ug.Name = data.Name
	ug.Uids = strings.Join(ids, ",")
	cache.CacheGidGroup[errorcode.Id] = ug
	cache.CacheGroupGid[ug.Name] = ug
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)
	id := r.FormValue("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 判断共享文件是否有在使用

	// var hasshare bool
	// sharerows, err := db.Mconn.GetRows("select readuser,rid,wid,writeuser from  sharefile")
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// for sharerows.Next() {
	// 	var readuser bool
	// 	var rid int64
	// 	var wid int64
	// 	var writeuser bool
	// 	sharerows.Scan(&readuser, &rid, &wid, &writeuser)
	// 	if !readuser {
	// 		if rid == int64(id32) {
	// 			hasshare = true
	// 			return
	// 		}
	// 	}
	// 	if !writeuser {
	// 		if wid == int64(id32) {
	// 			hasshare = true
	// 			return
	// 		}
	// 	}
	// }

	// if hasshare {
	// 	w.Write(errorcode.Error("没有权限"))
	// 	return
	// }
	// // 判断接口是否有在使用
	// var hasrest bool
	// restrows, err := db.Mconn.GetRows("select readuser,edituser,rid,eid from  apiproject")
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// for restrows.Next() {
	// 	var readuser bool
	// 	var rid int64
	// 	var eid int64
	// 	var edituser bool
	// 	restrows.Scan(&readuser, &rid, &eid, &edituser)
	// 	if !readuser {
	// 		if rid == int64(id32) {
	// 			hasshare = true
	// 		}
	// 	}
	// 	if !edituser {
	// 		if eid == int64(id32) {
	// 			hasshare = true
	// 			return
	// 		}
	// 	}
	// }
	// if hasrest {
	// 	w.Write(errorcode.Error("没有权限"))
	// 	return
	// }

	gsql := "delete from usergroup where id=? and cuid=? or cuid=?"
	_, err = db.Mconn.Update(gsql, id, uid, cache.SUPERID)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "usergroup",
		Action:   "delete",
	}

	delete(cache.CacheGroupGid, cache.CacheGidGroup[id64].Name)
	delete(cache.CacheGidGroup, id64)
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Get_groups)
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)

	gsql := "update usergroup set name=?,ids=? where id=? and cuid=? or cuid=?"
	ids := ""
	for i, v := range data.Users {
		if i == 0 {
			ids = strconv.FormatInt(cache.CacheRealNameUid[v], 10)
		} else {
			ids = ids + "," + strconv.FormatInt(cache.CacheRealNameUid[v], 10)
		}
	}
	_, err := db.Mconn.Update(gsql, data.Name, ids, data.Id, cache.SUPERID, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "usergroup",
		Action:   "update",
	}
	ug := &cache.UG{
		Gid:  data.Id,
		Name: data.Name,
		Uids: ids,
	}
	delete(cache.CacheGroupGid, cache.CacheGidGroup[data.Id].Name)
	cache.CacheGidGroup[data.Id] = ug
	cache.CacheGroupGid[data.Name] = ug
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
