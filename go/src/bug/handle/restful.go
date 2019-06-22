package handle

import (
	"bug/bugconfig"
	"bug/buglog"
	"bug/model"
	"encoding/json"
	"fmt"
	"gadb"
	"galog"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func RestList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		tl := &model.List_restful{}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()

		rows, err := conn.GetRows("select id,name,ownerid,auth,readuser,edituser,rid,eid from apiproject")
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			var oid int64
			var rid int64
			var eid int64
			tr := &model.Data_restful{}
			rows.Scan(&tr.Id, &tr.Name, &oid, &tr.Auth, &tr.Readuser, &tr.Edituser, &rid, &eid)
			tr.Owner = bugconfig.CacheUidRealName[oid]
			if tr.Readuser {
				tr.ReadName = bugconfig.CacheUidRealName[rid]
			} else {
				tr.ReadName = bugconfig.CacheGidGroup[rid]
			}
			if tr.Edituser {
				tr.EditName = bugconfig.CacheUidRealName[eid]
			} else {
				tr.EditName = bugconfig.CacheGidGroup[eid]
			}
			// 如果是创建者，直接是有权限的，添加进去
			if oid == bugconfig.CacheNickNameUid[nickname] {
				tl.List = append(tl.List, tr)
				continue
			}
			if tr.Auth {
				//如果是认证的
				if tr.Readuser {
					// 判断是否是可读的用户
					if rid == bugconfig.CacheNickNameUid[nickname] {
						tl.List = append(tl.List, tr)
						continue
					}

				} else {
					// 判断是都是可读的用户组
					var ids string
					err = conn.GetOne("select ids from usergroup where id=?", rid).Scan(&ids)
					if err != nil {
						galog.Error(err.Error())
						w.Write(errorcode.ErrorConnentMysql())
						return
					}
					var ingroup bool
					for _, v := range strings.Split(ids, ",") {
						if v == strconv.FormatInt(rid, 10) {
							ingroup = true
							break
						}
					}
					//如果在可读组里面就是可读的
					if ingroup {
						tl.List = append(tl.List, tr)
						continue
					}
				}
				// 如果是可编辑的权限也是有权限可见的
				if tr.Edituser {
					// 判断是否是可编辑的用户
					if eid == bugconfig.CacheNickNameUid[nickname] {
						tl.List = append(tl.List, tr)
					}
					continue
				} else {
					// 判断是都是可编辑的用户组
					var ids string
					err = conn.GetOne("select ids from usergroup where id=?", eid).Scan(&ids)
					if err != nil {
						galog.Error(err.Error())
						w.Write(errorcode.ErrorConnentMysql())
						return
					}
					var ingroup bool
					for _, v := range strings.Split(ids, ",") {
						if v == strconv.FormatInt(eid, 10) {
							ingroup = true
							break
						}
					}
					//如果在可读组里面就是可读的
					if ingroup {
						tl.List = append(tl.List, tr)
					}
				}
				// 没认证又不是创建者就是没权限
			}
		}
		send, _ := json.Marshal(tl)
		w.Write(send)
		return
	}
}

func RestUpdate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		tl := &model.Data_restful{}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(respbyte, tl)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		//验证权限修改
		// 拥有者才能修改权限
		var hasperm bool
		if bugconfig.CacheUidRealName[bugconfig.CacheNickNameUid[nickname]] == tl.Owner {
			hasperm = true
		}

		if !hasperm {
			w.Write(errorcode.ErrorNoPermission())
			return
		}

		var rid int64
		var eid int64
		if tl.Readuser {
			rid = bugconfig.CacheRealNameUid[tl.ReadName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if v == tl.ReadName {
					rid = k
				}
			}
		}

		if tl.Edituser {
			eid = bugconfig.CacheRealNameUid[tl.EditName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if v == tl.EditName {
					eid = k
				}
			}
		}

		_, err = conn.Update("update apiproject set name=?,auth=?,readuser=?,edituser=?,rid=?,eid=? where id=?",
			tl.Name,
			tl.Auth,
			tl.Readuser,
			tl.Edituser,
			rid, eid, tl.Id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "restproject",
		}
		err = il.Update(
			nickname, tl.Id, tl.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func RestAdd(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		uid := bugconfig.CacheNickNameUid[nickname]
		defer conn.Db.Close()
		dr := &model.Data_restful{}
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(bytedata, dr)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var rid int64
		var eid int64
		if dr.Readuser {
			rid = bugconfig.CacheRealNameUid[dr.ReadName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if v == dr.ReadName {
					rid = k
				}
			}
		}
		if dr.Edituser {
			eid = bugconfig.CacheRealNameUid[dr.EditName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if v == dr.EditName {
					eid = k
				}
			}
		}
		restsql := "insert into apiproject(name,ownerid,auth,readuser,edituser,rid,eid) values(?,?,?,?,?,?,?)"
		errorcode.Id, err = conn.Insert(restsql,
			dr.Name, uid, dr.Auth, dr.Readuser, dr.Edituser, rid, eid)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "restproject",
		}
		err = il.Add(
			nickname, errorcode.Id, dr.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func RestDel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		id := r.FormValue("id")
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		// 只有创建者才能删除
		eff, err := conn.Update("delete from apiproject where id=? and ownerid=?", id, bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 同时删除下面的所有接口
		if eff > 0 {
			_, err = conn.Update("delete from apilist where pid=? ", id)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "restproject",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func ApiList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		tl := &model.List_restful{}
		conn, nickname, err := logtokenmysql(r)
		pid := r.FormValue("pid")
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		//判断这个用户是否有权限访问
		hasperm, err := checkapiperm(conn, pid, bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if hasperm {
			rows, err := conn.GetRows("select id,name from apilist where pid=?", pid)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			for rows.Next() {
				tr := &model.Data_restful{}
				rows.Scan(&tr.Id, &tr.Name)
				tl.List = append(tl.List, tr)
			}
		} else {
			w.Write(errorcode.ErrorNoPermission())
			return
		}

		send, _ := json.Marshal(tl)
		w.Write(send)
		return
	}
}

func checkapiperm(conn *gadb.Db, pid string, uid int64) (bool, error) {
	var auth bool
	var readuser bool
	var edituser bool
	var oid int64
	var rid int64
	var eid int64
	err := conn.GetOne("select ownerid,auth,readuser,edituser,rid,eid from apiproject where id=?", pid).Scan(
		&oid, &auth, &readuser, &edituser, &rid, &eid)
	if err != nil {
		galog.Error(err.Error())
		return false, err
	}
	if uid == oid {
		return true, nil
	}
	if auth {
		//如果是认证的
		if readuser {
			// 判断是否是可读的用户
			if rid == uid {
				return true, nil
			}

		} else {
			// 判断是都是可读的用户组
			var ids string
			err = conn.GetOne("select ids from usergroup where id=?", rid).Scan(&ids)
			if err != nil {
				return false, err
			}
			for _, v := range strings.Split(ids, ",") {
				if v == strconv.FormatInt(rid, 10) {
					return true, nil
				}
			}
		}
		// 如果是可编辑的权限也是有权限可见的
		if edituser {
			// 判断是否是可编辑的用户
			if eid == uid {
				return true, nil
			}
		} else {
			// 判断是都是可编辑的用户组
			var ids string
			err = conn.GetOne("select ids from usergroup where id=?", eid).Scan(&ids)
			if err != nil {

				return false, nil
			}

			for _, v := range strings.Split(ids, ",") {
				if v == strconv.FormatInt(eid, 10) {
					return true, nil
				}
			}
		}
		// 没认证又不是创建者就是没权限
	}
	return false, nil
}

func checkeditperm(conn *gadb.Db, pid string, uid int64) (bool, error) {
	var auth bool
	var edituser bool
	var oid int64
	var eid int64
	err := conn.GetOne("select ownerid,auth,edituser,eid from apiproject where id=?", pid).Scan(
		&oid, &auth, &edituser, &eid)
	if err != nil {
		galog.Error(err.Error())
		return false, err
	}
	if uid == oid {
		return true, nil
	}
	if auth {
		// 如果是可编辑的权限也是有权限可见的
		if edituser {
			// 判断是否是可编辑的用户
			if eid == uid {
				return true, nil
			}
		} else {
			// 判断是都是可编辑的用户组
			var ids string
			err = conn.GetOne("select ids from usergroup where id=?", eid).Scan(&ids)
			if err != nil {

				return false, nil
			}

			for _, v := range strings.Split(ids, ",") {
				if v == strconv.FormatInt(eid, 10) {
					return true, nil
				}
			}
		}
		// 没认证又不是创建者就是没权限
	}
	return false, nil
}

func ApiUpdate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		tl := &model.Get_apilist{}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(respbyte, tl)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		//查出旧的
		var oldopts string
		err = conn.GetOne("select opts from apilist where id=?", tl.Id).Scan(&oldopts)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		hasperm, err := checkeditperm(conn, strconv.Itoa(tl.Pid), bugconfig.CacheNickNameUid[nickname])
		if hasperm {
			oidstr := make([]string, 0)
			for _, v := range tl.Opts {
				if v.Id > 0 {
					// 修改
					if tid, ok := bugconfig.CacheNameTid[v.Type]; ok {
						_, err = conn.Update("update options set info=?,name=?,tid=?,df=?,need=? where id=?",
							v.Info, v.Name, tid, v.Default, v.Need, v.Id)
						if err != nil {
							galog.Error(err.Error())
							w.Write(errorcode.ErrorConnentMysql())
							return
						}
						oidstr = append(oidstr, strconv.Itoa(v.Id))
					}
					tmpopts := make([]string, 0)
					for _, value := range strings.Split(oldopts, ",") {
						if value != strconv.Itoa(v.Id) {
							tmpopts = append(tmpopts)
						}
					}
					oldopts = strings.Join(tmpopts, ",")
				} else {
					// 添加
					if tid, ok := bugconfig.CacheNameTid[v.Type]; ok {
						tmpid, err := conn.Insert("insert into options(info,name,tid,df,need) values(?,?,?,?,?)",
							v.Info, v.Name, tid, v.Default, v.Need)
						if err != nil {
							galog.Error(err.Error())
							w.Write(errorcode.ErrorConnentMysql())
							return
						}
						oidstr = append(oidstr, strconv.FormatInt(tmpid, 10))
					}
				}
			}

			//删除多余的
			_, err = conn.Update(fmt.Sprintf("delete from options where id in (%s)", oldopts))
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			oids := strings.Join(oidstr, ",")
			ms := make([]string, 0)
			for _, v := range tl.Methods {
				ms = append(ms, v)
			}

			_, err = conn.Update("update apilist set url=?,information=?,opts=?,methods=?,result=?,name=?,hid=?,calltype=?,resp=? where id=?",
				tl.Url, html.EscapeString(tl.Information), oids, strings.Join(ms, ","), html.EscapeString(tl.Result), tl.Name, bugconfig.CacheHeaderHid[tl.Header], tl.CallType, html.EscapeString(tl.Resp), tl.Id)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}

		} else {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "api",
		}
		err = il.Update(
			nickname, tl.Id, tl.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func ApiAdd(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		al := &model.Get_apilist{}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		fmt.Println(string(respbyte))
		err = json.Unmarshal(respbyte, al)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		if al.Name == "" {
			galog.Error("name is empty")
			w.Write(errorcode.ErrorNull())
			return
		}
		if al.Url == "" {
			galog.Error("url is empty")
			w.Write(errorcode.ErrorNull())
			return
		}

		if len(al.Methods) == 0 {
			galog.Error("methoad is empty")
			w.Write(errorcode.ErrorNull())
			return
		}
		//先插入options
		oid := ""
		for i, v := range al.Opts {
			if tid, ok := bugconfig.CacheNameTid[v.Type]; ok {
				tmpid, err := conn.Insert("insert into options(info,name,tid,df,need) values(?,?,?,?,?)",
					v.Info, v.Name, tid, v.Default, v.Need)
				if err != nil {
					galog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
				if i == 0 {

					oid = strconv.FormatInt(tmpid, 10)
				} else {
					oid = oid + "," + strconv.FormatInt(tmpid, 10)
				}
			}

		}
		ms := ""
		for i, v := range al.Methods {
			if i == 0 {
				ms = v
			} else {
				ms = ms + "," + v
			}
		}
		var hid int64
		var ok bool
		if al.Header != "" {
			if hid, ok = bugconfig.CacheHeaderHid[al.Header]; !ok {
				galog.Error("key not found")
				w.Write(errorcode.ErrorKeyNotFound())
				return
			}
		}

		errorcode.Id, err = conn.Insert("insert into apilist(pid,url,information,opts,methods,result,name,uid,hid,calltype,resp) values(?,?,?,?,?,?,?,?,?,?,?)",
			al.Pid, al.Url, html.EscapeString(al.Information), oid, ms,
			html.EscapeString(al.Result), al.Name, bugconfig.CacheNickNameUid[nickname], hid, al.CallType, html.EscapeString(al.Resp))
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "api",
		}
		err = il.Add(
			nickname, errorcode.Id, al.Name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func ApiDel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	// 只有项目创建者和这个api的创建者才能删除这个接口文档
	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		id := r.FormValue("id")
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		var oids string
		var pid string
		var uid int64
		err = conn.GetOne("select pid,opts,uid from apilist where id=?", id).Scan(&pid, &oids, &uid)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		var oid int64
		err = conn.GetOne("select ownerid from apiproject where id=?", pid).Scan(&oid)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if uid != bugconfig.CacheNickNameUid[nickname] && oid != bugconfig.CacheNickNameUid[nickname] {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		ol := strings.Split(oids, ",")
		if ol[0] != "" {
			for _, v := range ol {
				_, err = conn.Update("delete from options where id=?", v)
				if err != nil {
					galog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
			}
		}

		_, err = conn.Update("delete from apilist where id=?", id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "api",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func ApiOne(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		sl := &model.Show_apilist{}
		id := r.FormValue("id")

		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl.Id, err = strconv.Atoi(id)
		if err != nil {
			w.Write(errorcode.ErrorParams())
			return
		}
		var oids string
		var ms string
		var hid int64
		err = conn.GetOne("select pid,url,information,opts,methods,result,name,hid,calltype,resp from apilist where id=?",
			id).Scan(
			&sl.Pid,
			&sl.Url,
			&sl.Information,
			&oids,
			&ms,
			&sl.Result,
			&sl.Name,
			&hid, &sl.CallType, &sl.Resp,
		)
		sl.Resp = html.UnescapeString(sl.Resp)
		sl.Result = html.UnescapeString(sl.Result)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 遍历请求头
		var ids string
		if hid > 0 {
			err = conn.GetOne("select hhids,remark from header where id=?", hid).Scan(&ids, &sl.Remark)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			hrows, err := conn.GetRows(fmt.Sprintf("select k,v from headerlist where id in (%v)", ids))
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			for hrows.Next() {
				two := &model.Table_headerlist{}
				hrows.Scan(&two.Key, &two.Value)
				sl.Header = append(sl.Header, two)
			}
		}
		// 判断权限
		hasperm, err := checkapiperm(conn, strconv.Itoa(sl.Pid), bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if !hasperm {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		// 遍历选项
		if len(oids) != 0 {

			orows, err := conn.GetRows(fmt.Sprintf("select id,info,name,tid,df,need from options where id in (%s)", oids))
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			for orows.Next() {
				o := &model.Table_opts{}
				var tid int64
				orows.Scan(&o.Id, &o.Info, &o.Name, &tid, &o.Default, &o.Need)
				o.Type = bugconfig.CacheTidName[tid]
				sl.Opts = append(sl.Opts, o)
			}

		}

		sl.Result = html.UnescapeString(sl.Result)
		sl.Information = html.UnescapeString(sl.Information)
		sl.Methods = strings.Split(ms, ",")
		send, _ := json.Marshal(sl)
		fmt.Println(string(send))
		w.Write(send)
		return
	}
}

func EditOne(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		sl := &model.One_apilist{}
		id := r.FormValue("id")

		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl.Id, err = strconv.Atoi(id)
		if err != nil {
			w.Write(errorcode.ErrorParams())
			return
		}
		var oids string
		var ms string
		var hid int64
		err = conn.GetOne("select pid,url,information,opts,methods,result,name,hid,calltype,resp from apilist where id=?",
			id).Scan(
			&sl.Pid,
			&sl.Url,
			&sl.Information,
			&oids,
			&ms,
			&sl.Result,
			&sl.Name,
			&hid, &sl.CallType, &sl.Resp,
		)
		sl.Resp = html.UnescapeString(sl.Resp)
		sl.Result = html.UnescapeString(sl.Result)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 遍历请求头
		var ids string
		if hid > 0 {
			err = conn.GetOne("select name,hhids,remark from header where id=?", hid).Scan(&sl.Header, &ids, &sl.Remark)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}
		// 判断权限
		hasperm, err := checkapiperm(conn, strconv.Itoa(sl.Pid), bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if !hasperm {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		// 遍历选项
		if len(oids) != 0 {

			orows, err := conn.GetRows(fmt.Sprintf("select id,info,name,tid,df,need from options where id in (%s)", oids))
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
			for orows.Next() {
				o := &model.Table_opts{}
				var tid int64
				orows.Scan(&o.Id, &o.Info, &o.Name, &tid, &o.Default, &o.Need)
				o.Type = bugconfig.CacheTidName[tid]
				sl.Opts = append(sl.Opts, o)
			}

		}

		sl.Result = html.UnescapeString(sl.Result)
		sl.Information = html.UnescapeString(sl.Information)
		sl.Methods = strings.Split(ms, ",")
		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
}

type header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type resp struct {
	Headers []*header `json:"header"`
	Resp    string    `json:"resp"`
	Url     string    `json:"url"`
	Method  string    `json:"method"`
}

func ApiResp(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		bb := &resp{}
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		err = json.Unmarshal(bytedata, bb)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		client := &http.Client{}
		//生成要访问的url

		//提交请求
		reqest, err := http.NewRequest(bb.Method, bb.Url, nil)

		for _, v := range bb.Headers {
			reqest.Header.Add(v.Key, v.Value)
		}

		if err != nil {
			galog.Error(err.Error())
			w.Write([]byte("请求失败, 请确认添加了url或者参数错误"))
			return
		}
		//处理返回结果
		response, err := client.Do(reqest)
		if err != nil {
			galog.Error(err.Error())
			w.Write([]byte("请求失败, 请确认添加了url或者参数错误"))
			return
		}
		defer response.Body.Close()
		send, err := ioutil.ReadAll(response.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write([]byte("获取数据失败"))
		}
		//fmt.Println(string(xx))
		//send, _ := json.Marshal(sl)
		//
		w.Write(send)
		return
	}
}
