package handle

import (
	"bug/asset"
	"bug/bugconfig"
	"bug/buglog"
	"bug/model"
	"encoding/json"
	"galog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func RoleList(w http.ResponseWriter, r *http.Request) {
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
		data := &model.List_roles{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("rolegroup", nickname, conn)
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
		s := "select id,name,rolelist from rolegroup"
		rows, err := conn.GetRows(s)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			var rids string
			one := &model.Data_roles{}
			rows.Scan(&one.Id, &one.Name, &rids)
			for _, v := range strings.Split(rids, ",") {
				id, _ := strconv.Atoi(v)
				one.RoleList = append(one.RoleList, bugconfig.CacheRidRole[int64(id)])
			}
			data.DataList = append(data.DataList, one)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func RoleDel(w http.ResponseWriter, r *http.Request) {
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
			permssion, err = asset.CheckPerm("rolegroup", nickname, conn)
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
		ssql := "select count(id) from user where rid=?"
		var count int
		err = conn.GetOne(ssql, id).Scan(&count)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if count > 0 {
			w.Write(errorcode.ErrorHasUser())
			return
		}
		isql := "delete from  rolegroup where id = ?"
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
			Classify: "rolegroup",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 删除缓存
		delete(bugconfig.CacheRidGroup, int64(id32))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func EditRole(w http.ResponseWriter, r *http.Request) {
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
		data := &model.Data_roles{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("rolegroup", nickname, conn)
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
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(bytedata, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		rl := make([]string, 0)
		for _, v := range data.RoleList {
			rl = append(rl, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
		}
		gsql := "update rolegroup set name=?,rolelist=?  where id=?"
		_, err = conn.Update(gsql, data.Name, strings.Join(rl, ","), data.Id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "rolegroup",
		}
		err = il.Update(
			nickname, data.Id, data.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		bugconfig.CacheRidGroup[data.Id] = data.Name
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func AddRole(w http.ResponseWriter, r *http.Request) {
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
		data := &model.Data_roles{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("rolegroup", nickname, conn)
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
		if data.Name == "" {
			galog.Error("name is empty")
			w.Write(errorcode.ErrorNull())
			return
		}
		ids := make([]string, 0)
		for _, v := range data.RoleList {
			ids = append(ids, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
		}
		gsql := "insert rolegroup(name,rolelist) values(?,?)"
		errorcode.Id, err = conn.InsertWithID(gsql, data.Name, strings.Join(ids, ","))
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "rolegroup",
		}
		err = il.Add(
			nickname, errorcode.Id, data.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		bugconfig.CacheRidGroup[errorcode.Id] = data.Name
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func RoleGroupName(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {

		conn, _, err := logtokenmysql(r)
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
		data := &model.Get_roles{}
		for _, v := range bugconfig.CacheRidGroup {
			data.Roles = append(data.Roles, v)
		}
		send, _ := json.Marshal(data)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
