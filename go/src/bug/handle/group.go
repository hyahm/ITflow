package handle

import (
	"bug/asset"
	"bug/bugconfig"
	"bug/buglog"
	"bug/model"
	"encoding/json"
	"fmt"
	"galog"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type getDepartment struct {
	Id         int64    `json:"id"`
	StatusList []string `json:"checklist"`
	Department string   `json:"departmentname"`
}

func AddBugGroup(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &getDepartment{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("statusgroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		list, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(list, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		// 重新排序
		ids := make([]string, 0)
		for _, v := range data.StatusList {
			ids = append(ids, fmt.Sprintf("%d", bugconfig.CacheStatusSid[v]))
		}

		ss := strings.Join(ids, ",")

		isql := "insert into statusgroup(name,sids) values(?,?)"
		errorcode.Id, err = conn.InsertWithID(isql, data.Department, ss)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "statusgroup",
		}
		err = il.Add(
			nickname, errorcode.Id, data.Department)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 添加缓存
		bugconfig.CacheSgidGroup[errorcode.Id] = data.Department
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func EditBugGroup(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {

		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &getDepartment{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("statusgroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		list, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		fmt.Println(string(list))
		err = json.Unmarshal(list, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		ssids := make([]string, 0)
		for _, v := range data.StatusList {
			// 没找到就是key
			var sid int64
			var ok bool
			if sid, ok = bugconfig.CacheStatusSid[v]; !ok {
				w.Write(errorcode.ErrorKeyNotFound())
				return
			}
			ssids = append(ssids, strconv.FormatInt(sid, 10))
		}
		ss := strings.Join(ssids, ",")
		isql := "update statusgroup set name =?,sids=? where id = ?"
		_, err = conn.Update(isql, data.Department, ss, data.Id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "statusgroup",
		}
		err = il.Update(
			nickname, data.Id, data.Department)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		bugconfig.CacheSgidGroup[data.Id] = data.Department
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type department struct {
	Id            int64    `json:"id"`
	Name          string   `json:"name"`
	BugstatusList []string `json:"bugstatuslist"`
}

type departmentList struct {
	DepartmentList []*department `json:"departmentlist"`
	Code           int           `json:"statuscode"`
}

func BugGroupList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {

		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("statusgroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		data := &departmentList{}
		s := "select id,name,sids from statusgroup"
		rows, err := conn.GetRows(s)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			var ids string
			one := &department{}
			rows.Scan(&one.Id, &one.Name, &ids)

			for _, v := range strings.Split(ids, ",") {
				id, err := strconv.Atoi(v)
				if err != nil {
					log.Println(err)
					continue
				}

				one.BugstatusList = append(one.BugstatusList, bugconfig.CacheSidStatus[int64(id)])
			}
			data.DepartmentList = append(data.DepartmentList, one)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func BugGroupDel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {

		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		id := r.FormValue("id")
		id32, err := strconv.Atoi(id)
		if err != nil {
			w.Write(errorcode.ErrorParams())
			return
		}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("statusgroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		ssql := "select count(id) from user where bugsid=?"
		var count int
		err = conn.GetOne(ssql, id).Scan(&count)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if count > 0 {
			w.Write(errorcode.ErrorHasGroup())
			return
		}
		isql := "delete from  statusgroup where id = ?"
		_, err = conn.Update(isql, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "statusgroup",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		//更新缓存
		delete(bugconfig.CacheSgidGroup, int64(id32))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GroupGet(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &model.Send_groups{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("usergroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		gsql := "select id,name,ids from usergroup"
		rows, err := conn.GetRows(gsql)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			onegroup := &model.Get_groups{}
			var users string
			rows.Scan(&onegroup.Id, &onegroup.Name, &users)
			for _, v := range strings.Split(users, ",") {
				uid, _ := strconv.Atoi(v)
				onegroup.Users = append(onegroup.Users, bugconfig.CacheUidRealName[int64(uid)])
			}
			data.GroupList = append(data.GroupList, onegroup)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GroupAdd(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("usergroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		data := &model.Get_groups{}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(respbyte, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		ids := make([]string, 0)
		for _, v := range data.Users {
			var uid int64
			var ok bool
			if uid, ok = bugconfig.CacheRealNameUid[v]; !ok {
				w.Write(errorcode.ErrorKeyNotFound())
				return
			}
			ids = append(ids, strconv.FormatInt(uid, 10))
		}
		gsql := "insert usergroup(name,ids,cuid) values(?,?,?)"
		errorcode.Id, err = conn.Insert(gsql, data.Name, strings.Join(ids, ","), bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "usergroup",
		}
		err = il.Add(
			nickname, errorcode.Id, data.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		bugconfig.CacheGidGroup[errorcode.Id] = data.Name
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GroupDel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodGet {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("usergroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		id := r.FormValue("id")
		id32, err := strconv.Atoi(id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		// 判断共享文件是否有在使用

		var hasshare bool
		sharerows, err := conn.GetRows("select readuser,rid,wid,writeuser from  sharefile")
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for sharerows.Next() {
			var readuser bool
			var rid int64
			var wid int64
			var writeuser bool
			sharerows.Scan(&readuser, &rid, &wid, &writeuser)
			if !readuser {
				if rid == int64(id32) {
					hasshare = true
					return
				}
			}
			if !writeuser {
				if wid == int64(id32) {
					hasshare = true
					return
				}
			}
		}

		if hasshare {
			w.Write(errorcode.ErrorHasGroup())
			return
		}
		// 判断接口是否有在使用
		var hasrest bool
		restrows, err := conn.GetRows("select readuser,edituser,rid,eid from  apiproject")
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for restrows.Next() {
			var readuser bool
			var rid int64
			var eid int64
			var edituser bool
			restrows.Scan(&readuser, &rid, &eid, &edituser)
			if !readuser {
				if rid == int64(id32) {
					hasshare = true
				}
			}
			if !edituser {
				if eid == int64(id32) {
					hasshare = true
					return
				}
			}
		}
		if hasrest {
			w.Write(errorcode.ErrorHasGroup())
			return
		}

		gsql := "delete from usergroup where id=? and cuid=?"
		_, err = conn.Update(gsql, id, bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "usergroup",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		delete(bugconfig.CacheGidGroup, int64(id32))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GroupUpdate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &model.Get_groups{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("usergroup", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(respbyte, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		gsql := "update usergroup set name=?,ids=? where id=? and cuid=?"
		ids := ""
		for i, v := range data.Users {
			if i == 0 {
				ids = strconv.FormatInt(bugconfig.CacheRealNameUid[v], 10)
			} else {
				ids = ids + "," + strconv.FormatInt(bugconfig.CacheRealNameUid[v], 10)
			}
		}
		_, err = conn.Update(gsql, data.Name, ids, data.Id, bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "usergroup",
		}
		err = il.Update(
			nickname, data.Id, data.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		bugconfig.CacheGidGroup[data.Id] = data.Name
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
