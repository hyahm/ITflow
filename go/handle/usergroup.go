package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/usergroup"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddBugGroup(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	// data := xmux.GetInstance(r).Data.(*status.StatusGroup)

	// isql := "insert into statusgroup(name,sids) values(?,?)"
	// var err error
	// errorcode.ID, err = db.Mconn.Insert(isql, data.Name, data.StatusList.ToStore())
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	// 添加缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)

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

// bug状态组
// func BugGroupList(w http.ResponseWriter, r *http.Request) {

// 	errorcode := &response.Response{}
// 	DepartmentList := make([]model.StatusGroup, 0)
// 	s := "select id,name,sids from statusgroup"
// 	rows, err := db.Mconn.GetRows(s)
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.ErrorE(err))
// 		return
// 	}
// 	for rows.Next() {
// 		var ids string
// 		one := &department{}
// 		rows.Scan(&one.Id, &one.Name, &ids)
// 		idrows, err := db.Mconn.GetRowsIn("select name from status where id in (?)", strings.Split(ids, ","))
// 		if err != nil {
// 			golog.Error(err)
// 			w.Write(errorcode.ErrorE(err))
// 			return
// 		}

// 		for idrows.Next() {
// 			var name string
// 			err = idrows.Scan(&name)
// 			if err != nil {
// 				golog.Error()
// 				continue
// 			}
// 			if name != "" {
// 				one.BugstatusList = append(one.BugstatusList, name)
// 			}

// 		}
// 		idrows.Close()
// 		DepartmentList = append(DepartmentList, one)
// 	}
// 	rows.Close()
// 	send, _ := json.Marshal(data)
// 	w.Write(send)
// 	return

// }

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
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Info(err)
			continue
		}
		data.UserGroupNames = append(data.UserGroupNames, name)
	}

	w.Write(data.Marshal())
}

func UserGroupGet(w http.ResponseWriter, r *http.Request) {
	// 可以获取所有用户组
	uid := xmux.GetInstance(r).Get("uid").(int64)
	usergroups, err := model.GetUserGroupList(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: usergroups,
	}
	w.Write(res.Marshal())

}

func GroupAdd(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	// uid := xmux.GetInstance(r).Get("uid").(int64)
	// data := xmux.GetInstance(r).Data.(*model.UserGroup)
	// rows, err := db.Mconn.GetRowsIn("select id from user where realname in (?)", data.Users)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// ids := make([]string, 0)
	// id := new(string)
	// for rows.Next() {
	// 	err = rows.Scan(id)
	// 	if err != nil {
	// 		golog.Info(err)
	// 		continue
	// 	}
	// 	ids = append(ids, *id)
	// }
	// rows.Close()
	// gsql := "insert usergroup(name,ids,uid) values(?,?,?)"
	// errorcode.ID, err = db.Mconn.Insert(gsql, data.Name, strings.Join(ids, ","), uid)
	// if err != nil {
	// 	golog.Error(err)
	// 	if err.(*mysql.MySQLError).Number == 1062 {
	// 		w.Write(errorcode.ErrorE(db.DuplicateErr))
	// 		return
	// 	}
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }

	w.Write(errorcode.Success())
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
}

func GroupUpdate(w http.ResponseWriter, r *http.Request) {
	// uid := xmux.GetInstance(r).Get("uid").(int64)

	ug := xmux.GetInstance(r).Data.(*model.UserGroup)
	err := ug.Update()
	if err != nil {
		golog.Error(err)
		response.ErrorE(err)
		return
	}

	w.Write(response.Success())

}
