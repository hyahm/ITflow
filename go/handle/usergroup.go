package handle

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/response"
	"itflow/internal/usergroup"
	"net/http"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	// data := xmux.GetInstance(r).Data.(*status.StatusGroup)

	// isql := "insert into statusgroup(name,sids) values(?,?)"
	// var err error
	// errorcode.Id, err = db.Mconn.Insert(isql, data.Name, data.StatusList.ToStore())
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	// 添加缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	// data := xmux.GetInstance(r).Data.(*status.StatusGroup)

	// isql := "update statusgroup set name =?,sids=? where id = ?"
	// _, err := db.Mconn.Update(isql, data.Name, data.StatusList.ToStore(), data.Id)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type department struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	BugstatusList []string `json:"bugstatuslist"`
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
		var ids string
		one := &department{}
		rows.Scan(&one.Id, &one.Name, &ids)
		idrows, err := db.Mconn.GetRowsIn("select name from status where id in (?)", strings.Split(ids, ","))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

		for idrows.Next() {
			var name string
			err = idrows.Scan(&name)
			if err != nil {
				golog.Error()
				continue
			}
			if name != "" {
				one.BugstatusList = append(one.BugstatusList, name)
			}

		}
		idrows.Close()
		data.DepartmentList = append(data.DepartmentList, one)
	}
	rows.Close()
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
	rows, err := db.Mconn.GetRows("select name from usergroup")
	if err != nil {
		golog.Error(err)
		w.Write(data.ErrorE(err))
		return
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Info(err)
			continue
		}
		data.UserGroupNames = append(data.UserGroupNames, name)
	}
	rows.Close()
	w.Write(data.Marshal())
	return

}

func UserGroupGet(w http.ResponseWriter, r *http.Request) {
	// 可以获取所有用户组

	data := &usergroup.RespUserGroupList{
		UserGroupList: make([]*usergroup.RespUserGroup, 0),
	}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	var rows *sql.Rows
	var err error
	if cache.SUPERID == uid {
		gsql := "select id,name,ifnull(ids,'') from usergroup"
		rows, err = db.Mconn.GetRows(gsql)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}
	} else {
		gsql := "select id,name,ids from usergroup where uid=?"
		rows, err = db.Mconn.GetRows(gsql, uid)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}
	}
	realname := new(string)
	for rows.Next() {
		onegroup := &usergroup.RespUserGroup{
			Users: make([]string, 0),
		}
		var users string
		rows.Scan(&onegroup.Id, &onegroup.Name, &users)
		namerows, err := db.Mconn.GetRowsIn("select realname from user where id in (?)",
			strings.Split(users, ","))
		if err != nil {
			golog.Info(err)
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
		namerows.Close()
		data.UserGroupList = append(data.UserGroupList, onegroup)
	}
	rows.Close()
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func GroupAdd(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	data := xmux.GetInstance(r).Data.(*usergroup.RespUserGroup)
	rows, err := db.Mconn.GetRowsIn("select id from user where realname in (?)", data.Users)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	ids := make([]string, 0)
	id := new(string)
	for rows.Next() {
		err = rows.Scan(id)
		if err != nil {
			golog.Info(err)
			continue
		}
		ids = append(ids, *id)
	}
	rows.Close()
	gsql := "insert usergroup(name,ids,uid) values(?,?,?)"
	errorcode.Id, err = db.Mconn.Insert(gsql, data.Name, strings.Join(ids, ","), uid)
	if err != nil {
		golog.Error(err)
		if err.(*mysql.MySQLError).Number == 1062 {
			w.Write(errorcode.ErrorE(db.DuplicateErr))
			return
		}
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write(errorcode.Success())
	return

}

func GroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	uid := xmux.GetInstance(r).Get("uid").(int64)
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
	uid := xmux.GetInstance(r).Get("uid").(int64)
	data := xmux.GetInstance(r).Data.(*usergroup.RespUpdateUserGroup)
	golog.Infof("%+v", *data)
	ids, err := data.GetIds()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info(ids)
	updatesql := "update usergroup set name=?, ids=?, uid=? where id=?"
	_, err = db.Mconn.Update(updatesql, data.Name, ids, uid, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write(errorcode.Success())
	return

}
