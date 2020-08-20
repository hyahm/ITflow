package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/response"
	"itflow/internal/status"
	"itflow/internal/usergroup"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	isql := "insert into statusgroup(name,sids) values(?,?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(isql, data.Name, data.StatusList.ToStore())
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 添加缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*status.StatusGroup)

	isql := "update statusgroup set name =?,sids=? where id = ?"
	_, err := db.Mconn.Update(isql, data.Name, data.StatusList.ToStore(), data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

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
		data.DepartmentList = append(data.DepartmentList, one)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func BugGroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")

	ssql := "select count(id) from user where bugsid=?"
	var count int
	err := db.Mconn.GetOne(ssql, id).Scan(&count)
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

	//更新缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupNamesGet(w http.ResponseWriter, r *http.Request) {

	data := &usergroup.RespUserGroupName{
		UserGroupNames: make([]string, 0),
	}
	// for _, v := range cache.CacheUGidUserGroup {
	// 	data.UserGroupNames = append(data.UserGroupNames, v.Name)
	// }
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func UserGroupGet(w http.ResponseWriter, r *http.Request) {
	// 可以获取所有用户组
	errorcode := &response.Response{}

	data := &usergroup.RespUserGroupList{
		UserGroupList: make([]*usergroup.RespUserGroup, 0),
	}
	uid := xmux.GetData(r).Get("uid").(int64)
	if cache.SUPERID == uid {
		gsql := "select id,name,ifnull(ids,'') from usergroup"
		rows, err := db.Mconn.GetRows(gsql)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		realname := new(string)
		for rows.Next() {
			onegroup := &usergroup.RespUserGroup{
				Users: make([]string, 0),
			}
			var users string
			rows.Scan(&onegroup.Id, &onegroup.Name, &users)
			namerows, err := db.Mconn.GetRows(fmt.Sprintf("select realname from user where id in ('%s')", users))
			if err != nil {
				golog.Error(err)
				continue
			}

			for namerows.Next() {
				err = namerows.Scan(realname)
				if err != nil {
					golog.Error(err)
					continue
				}
				onegroup.Users = append(onegroup.Users, *realname)
			}

			data.UserGroupList = append(data.UserGroupList, onegroup)
		}
	} else {
		gsql := "select id,name,ids from usergroup where uid=?"
		rows, err := db.Mconn.GetRows(gsql, uid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for rows.Next() {
			onegroup := &usergroup.RespUserGroup{}
			var users string
			rows.Scan(&onegroup.Id, &onegroup.Name, &users)
			// for _, v := range strings.Split(users, ",") {
			// uid, err := strconv.ParseInt(v, 10, 64)
			// if err != nil {
			// 	continue
			// }
			// onegroup.Users = append(onegroup.Users, cache.CacheUidRealName[uid])
			// }
			data.UserGroupList = append(data.UserGroupList, onegroup)
		}
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func GroupAdd(w http.ResponseWriter, r *http.Request) {
	golog.Info("add..................")
	errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	data := xmux.GetData(r).Data.(*usergroup.RespUserGroup)

	ids := make([]string, 0)
	// for _, v := range data.Users {
	// 	var uid int64
	// 	var ok bool
	// 	if uid, ok = cache.CacheRealNameUid[v]; !ok {
	// 		w.Write(errorcode.Errorf("没有此用户"))
	// 		return
	// 	}
	// 	ids = append(ids, strconv.FormatInt(uid, 10))
	// }
	// if _, ok := cache.CacheUserGroupUGid[data.Name]; ok {
	// 	errorcode.Code = 1
	// 	w.Write(errorcode.Errorf("%s 重复", data.Name))
	// 	return
	// }
	gsql := "insert usergroup(name,ids,uid) values(?,?,?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(gsql, data.Name, strings.Join(ids, ","), uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	ug := &cache.UG{}
	ug.Ugid = errorcode.Id
	ug.Name = data.Name
	ug.Uids = strings.Join(ids, ",")

	// if ug, ok := cache.CacheUGidUserGroup[errorcode.Id]; ok {
	// 	delete(cache.CacheUserGroupUGid, ug.Name)
	// }

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	id := r.FormValue("id")
	// id64, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

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

	gsql := "delete from usergroup where id=? and uid=?"
	_, err := db.Mconn.Update(gsql, id, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// if ug, ok := cache.CacheUGidUserGroup[id64]; ok {
	// 	delete(cache.CacheUserGroupUGid, ug.Name)
	// }

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GroupUpdate(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	data := xmux.GetData(r).Data.(*usergroup.RespUpdateUserGroup)
	// insert into usergroup(name, ids, uid) values('', (select id from user where realname in ('')), 1)
	// ids := make([]string, 0)
	// ulen := len(data.Users)
	// if ulen == 0 {

	// } else if ulen == 1 {
	// 	var id string
	// err := db.Mconn.GetOne("select id from user where realname=?", data.Users[0]).Scan(&id)
	// 	if err != nil {
	// 		golog.Error(err)
	// 		w.Write(errorcode.ErrorE(err))
	// 		return
	// 	}
	// 	ids = append(ids, id)
	// } else {
	db.Mconn.OpenDebug()
	idrows, err := db.Mconn.GetRowsIn("select id from user where realname in (?)", data.Users)
	golog.Info(db.Mconn.GetSql())
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	ids := make([]string, 0)
	for idrows.Next() {
		var id string
		err = idrows.Scan(&id)
		if err != nil {
			golog.Error(err)
			continue
		}
		ids = append(ids, id)
	}
	// }

	golog.Info(ids)
	updatesql := "update usergroup set name=?, ids=?, uid=? where id=?"
	_, err = db.Mconn.Update(updatesql, data.Name, strings.Join(ids, ","), uid, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// for _, v := range data.Users {
	// 	userid, ok := cache.CacheRealNameUid[v]
	// 	if !ok {
	// 		continue
	// 	}
	// 	ids = append(ids, strconv.FormatInt(userid, 10))

	// }
	// usergroup := &model.UserGroups{
	// 	Id:   data.Id,
	// 	Ids:  strings.Join(ids, ","),
	// 	Name: data.Name,
	// }
	// err := usergroup.Update()
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	// ug := &cache.UG{
	// 	Ugid: data.Id,
	// 	Name: data.Name,
	// 	Uids: strings.Join(ids, ","),
	// }
	// if thisUg, ok := cache.CacheUGidUserGroup[data.Id]; ok {
	// 	delete(cache.CacheUserGroupUGid, thisUg.Name)
	// }

	w.Write(errorcode.Success())
	return

}
