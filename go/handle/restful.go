package handle

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/restful"
	"itflow/model"
	network "itflow/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func RestList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	tl := &network.List_restful{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	rows, err := db.Mconn.GetRows("select id,name,ownerid,auth,readuser,edituser,rid,eid from apiproject")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var oid int64
		var rid int64
		var eid int64
		tr := &network.Data_restful{}
		rows.Scan(&tr.Id, &tr.Name, &oid, &tr.Auth, &tr.Readuser, &tr.Edituser, &rid, &eid)
		tr.Owner = cache.CacheUidRealName[oid]
		if tr.Readuser {
			tr.ReadName = cache.CacheUidRealName[rid]
		} else {
			tr.ReadName = cache.CacheUGidUserGroup[rid].Name
		}
		if tr.Edituser {
			tr.EditName = cache.CacheUidRealName[eid]
		} else {
			tr.EditName = cache.CacheUGidUserGroup[eid].Name
		}
		// 如果是创建者，直接是有权限的，添加进去
		if oid == cache.CacheNickNameUid[nickname] {
			tl.List = append(tl.List, tr)
			continue
		}
		if tr.Auth {
			//如果是认证的
			if tr.Readuser {
				// 判断是否是可读的用户
				if rid == cache.CacheNickNameUid[nickname] {
					tl.List = append(tl.List, tr)
					continue
				}

			} else {
				// 判断是都是可读的用户组
				var ids string
				err := db.Mconn.GetOne("select ids from usergroup where id=?", rid).Scan(&ids)
				if err != nil {
					golog.Error(err)
					w.Write(errorcode.ErrorE(err))
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
				if eid == cache.CacheNickNameUid[nickname] {
					tl.List = append(tl.List, tr)
				}
				continue
			} else {
				// 判断是都是可编辑的用户组
				var ids string
				err := db.Mconn.GetOne("select ids from usergroup where id=?", eid).Scan(&ids)
				if err != nil {
					golog.Error(err)
					w.Write(errorcode.ErrorE(err))
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

func RestUpdate(w http.ResponseWriter, r *http.Request) {
	nickname := xmux.GetData(r).Get("nickname").(string)
	errorcode := &response.Response{}
	tl := xmux.GetData(r).Data.(*model.Data_restful)

	//验证权限修改
	// 拥有者才能修改权限
	var hasperm bool
	if cache.CacheUidRealName[cache.CacheNickNameUid[nickname]] == tl.Owner {
		hasperm = true
	}

	if !hasperm {
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	var rid int64
	var eid int64
	if tl.Readuser {
		rid = cache.CacheRealNameUid[tl.ReadName]
	} else {
		for k, v := range cache.CacheUGidUserGroup {
			if v.Name == tl.ReadName {
				rid = k
			}
		}
	}

	if tl.Edituser {
		eid = cache.CacheRealNameUid[tl.EditName]
	} else {
		for k, v := range cache.CacheUGidUserGroup {
			if v.Name == tl.EditName {
				eid = k
			}
		}
	}

	_, err := db.Mconn.Update("update apiproject set name=?,auth=?,readuser=?,edituser=?,rid=?,eid=? where id=?",
		tl.Name,
		tl.Auth,
		tl.Readuser,
		tl.Edituser,
		rid, eid, tl.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "rest",
		Action:   "update",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RestAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := cache.CacheNickNameUid[nickname]

	dr := xmux.GetData(r).Data.(*model.Data_restful)

	var rid int64
	var eid int64
	if dr.Readuser {
		rid = cache.CacheRealNameUid[dr.ReadName]
	} else {
		for k, v := range cache.CacheUGidUserGroup {
			if v.Name == dr.ReadName {
				rid = k
			}
		}
	}
	if dr.Edituser {
		eid = cache.CacheRealNameUid[dr.EditName]
	} else {
		for k, v := range cache.CacheUGidUserGroup {
			if v.Name == dr.EditName {
				eid = k
			}
		}
	}
	restsql := "insert into apiproject(name,ownerid,auth,readuser,edituser,rid,eid) values(?,?,?,?,?,?,?)"
	var err error
	errorcode.Id, err = db.Mconn.Insert(restsql,
		dr.Name, uid, dr.Auth, dr.Readuser, dr.Edituser, rid, eid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "rest",
		Action:   "add",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func RestDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	id := r.FormValue("id")
	nickname := xmux.GetData(r).Get("nickname").(string)
	// 只有创建者才能删除
	eff, err := db.Mconn.Update("delete from apiproject where id=? and ownerid=?", id, cache.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 同时删除下面的所有接口
	if eff > 0 {
		_, err = db.Mconn.Update("delete from apilist where pid=? ", id)
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
		Classify: "rest",
		Action:   "delete",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ApiList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	tl := &network.List_restful{}

	pid := r.FormValue("pid")

	nickname := xmux.GetData(r).Get("nickname").(string)
	//判断这个用户是否有权限访问
	hasperm, err := checkapiperm(pid, cache.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if hasperm {
		rows, err := db.Mconn.GetRows("select id,name from apilist where pid=?", pid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for rows.Next() {
			tr := &network.Data_restful{}
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

func checkapiperm(pid string, uid int64) (bool, error) {
	var auth bool
	var readuser bool
	var edituser bool
	var oid int64
	var rid int64
	var eid int64
	err := db.Mconn.GetOne("select ownerid,auth,readuser,edituser,rid,eid from apiproject where id=?", pid).Scan(&oid, &auth, &readuser, &edituser, &rid, &eid)
	if err != nil {
		golog.Error(err)
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
			err = db.Mconn.GetOne("select ids from usergroup where id=?", rid).Scan(&ids)
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
			err := db.Mconn.GetOne("select ids from usergroup where id=?", eid).Scan(&ids)
			if err != nil {
				return false, err
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

func checkeditperm(pid string, uid int64) (bool, error) {
	var auth bool
	var edituser bool
	var oid int64
	var eid int64
	err := db.Mconn.GetOne("select ownerid,auth,edituser,eid from apiproject where id=?", pid).Scan(
		&oid, &auth, &edituser, &eid)
	if err != nil {
		golog.Error(err)
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
			err := db.Mconn.GetOne("select ids from usergroup where id=?", eid).Scan(&ids)
			if err != nil {
				return false, err
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

	errorcode := &response.Response{}
	tl := xmux.GetData(r).Data.(*model.Get_apilist)
	nickname := xmux.GetData(r).Get("nickname").(string)

	//查出旧的
	var oldopts string
	err := db.Mconn.GetOne("select opts from apilist where id=?", tl.Id).Scan(&oldopts)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	hasperm, err := checkeditperm(strconv.Itoa(tl.Pid), cache.CacheNickNameUid[nickname])
	if hasperm {
		oidstr := make([]string, 0)
		for _, v := range tl.Opts {
			if v.Id > 0 {
				// 修改
				if tid, ok := cache.CacheNameTid[v.Type]; ok {
					_, err = db.Mconn.Update("update options set info=?,name=?,tid=?,df=?,need=? where id=?",
						v.Info, v.Name, tid, v.Default, v.Need, v.Id)
					if err != nil {
						golog.Error(err)
						w.Write(errorcode.ErrorE(err))
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
				if tid, ok := cache.CacheNameTid[v.Type]; ok {
					tmpid, err := db.Mconn.Insert("insert into options(info,name,tid,df,need) values(?,?,?,?,?)",
						v.Info, v.Name, tid, v.Default, v.Need)
					if err != nil {
						golog.Error(err)
						w.Write(errorcode.ErrorE(err))
						return
					}
					oidstr = append(oidstr, strconv.FormatInt(tmpid, 10))
				}
			}
		}

		//删除多余的
		_, err = db.Mconn.Update(fmt.Sprintf("delete from options where id in (%s)", oldopts))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		oids := strings.Join(oidstr, ",")
		ms := make([]string, 0)
		for _, v := range tl.Methods {
			ms = append(ms, v)
		}

		_, err = db.Mconn.Update("update apilist set url=?,information=?,opts=?,methods=?,result=?,name=?,hid=?,calltype=?,resp=? where id=?",
			tl.Url, html.EscapeString(tl.Information), oids, strings.Join(ms, ","), html.EscapeString(tl.Result), tl.Name, cache.CacheHeaderHid[tl.Header], tl.CallType, html.EscapeString(tl.Resp), tl.Id)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

	} else {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "api",
		Action:   "update",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ApiAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	al := xmux.GetData(r).Data.(*model.Get_apilist)

	if al.Name == "" {
		golog.Error("name is empty")
		w.Write(errorcode.Error("name is empty"))
		return
	}
	if al.Url == "" {
		golog.Error("url is empty")
		w.Write(errorcode.Error("url is empty"))
		return
	}

	if len(al.Methods) == 0 {
		golog.Error("methoad is empty")
		w.Write(errorcode.Error("methoad is empty"))
		return
	}
	//先插入options
	oid := ""
	for i, v := range al.Opts {
		if tid, ok := cache.CacheNameTid[v.Type]; ok {
			tmpid, err := db.Mconn.Insert("insert into options(info,name,tid,df,need) values(?,?,?,?,?)",
				v.Info, v.Name, tid, v.Default, v.Need)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
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
		if hid, ok = cache.CacheHeaderHid[al.Header]; !ok {
			golog.Error("key not found")
			w.Write(errorcode.Error("key not found"))
			return
		}
	}
	var err error
	errorcode.Id, err = db.Mconn.Insert("insert into apilist(pid,url,information,opts,methods,result,name,uid,hid,calltype,resp) values(?,?,?,?,?,?,?,?,?,?,?)",
		al.Pid, al.Url, html.EscapeString(al.Information), oid, ms,
		html.EscapeString(al.Result), al.Name, cache.CacheNickNameUid[nickname], hid, al.CallType, html.EscapeString(al.Resp))
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "api",
		Action:   "add",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ApiDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	id := r.FormValue("id")

	var oids string
	var pid string
	var uid int64
	err := db.Mconn.GetOne("select pid,opts,uid from apilist where id=?", id).Scan(&pid, &oids, &uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	var oid int64
	err = db.Mconn.GetOne("select ownerid from apiproject where id=?", pid).Scan(&oid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	nickname := xmux.GetData(r).Get("nickname").(string)
	if uid != cache.CacheNickNameUid[nickname] && oid != cache.CacheNickNameUid[nickname] {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	ol := strings.Split(oids, ",")
	if ol[0] != "" {
		for _, v := range ol {
			_, err = db.Mconn.Update("delete from options where id=?", v)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}
		}
	}

	_, err = db.Mconn.Update("delete from apilist where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "api",
		Action:   "delete",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ApiOne(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	sl := &network.Show_apilist{}
	id := r.FormValue("id")
	nickname := xmux.GetData(r).Get("nickname").(string)
	var err error
	sl.Id, err = strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	var oids string
	var ms string
	var hid int64
	err = db.Mconn.GetOne("select pid,url,information,opts,methods,result,name,hid,calltype,resp from apilist where id=?",
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
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	sl.Resp = html.UnescapeString(sl.Resp)
	sl.Result = html.UnescapeString(sl.Result)

	// 遍历请求头
	var ids string
	if hid > 0 {
		err := db.Mconn.GetOne("select hhids,remark from header where id=?", hid).Scan(&ids, &sl.Remark)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

		hrows, err := db.Mconn.GetRows(fmt.Sprintf("select k,v from headerlist where id in (%v)", ids))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for hrows.Next() {
			two := &network.Table_headerlist{}
			hrows.Scan(&two.Key, &two.Value)
			sl.Header = append(sl.Header, two)
		}
	}
	// 判断权限
	hasperm, err := checkapiperm(strconv.Itoa(sl.Pid), cache.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if !hasperm {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 遍历选项
	if len(oids) != 0 {

		orows, err := db.Mconn.GetRows(fmt.Sprintf("select id,info,name,tid,df,need from options where id in (%s)", oids))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for orows.Next() {
			o := &network.Table_opts{}
			var tid int64
			orows.Scan(&o.Id, &o.Info, &o.Name, &tid, &o.Default, &o.Need)
			o.Type = cache.CacheTidName[tid]
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

func EditOne(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	sl := &network.One_apilist{}
	id := r.FormValue("id")
	var err error
	sl.Id, err = strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	var oids string
	var ms string
	var hid int64
	err = db.Mconn.GetOne("select pid,url,information,opts,methods,result,name,hid,calltype,resp from apilist where id=?",
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
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	sl.Resp = html.UnescapeString(sl.Resp)
	sl.Result = html.UnescapeString(sl.Result)

	// 遍历请求头
	var ids string
	if hid > 0 {
		err := db.Mconn.GetOne("select name,hhids,remark from header where id=?", hid).Scan(&sl.Header, &ids, &sl.Remark)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

	}
	// 判断权限
	nickname := xmux.GetData(r).Get("nickname").(string)
	hasperm, err := checkapiperm(strconv.Itoa(sl.Pid), cache.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if !hasperm {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 遍历选项
	if len(oids) != 0 {

		orows, err := db.Mconn.GetRows(fmt.Sprintf("select id,info,name,tid,df,need from options where id in (%s)", oids))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for orows.Next() {
			o := &network.Table_opts{}
			var tid int64
			orows.Scan(&o.Id, &o.Info, &o.Name, &tid, &o.Default, &o.Need)
			o.Type = cache.CacheTidName[tid]
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

func ApiResp(w http.ResponseWriter, r *http.Request) {

	bb := xmux.GetData(r).Data.(*restful.Resp)

	client := &http.Client{}
	//生成要访问的url

	//提交请求
	reqest, err := http.NewRequest(bb.Method, bb.Url, nil)

	for _, v := range bb.Headers {
		reqest.Header.Add(v.Key, v.Value)
	}

	if err != nil {
		golog.Error(err)
		w.Write([]byte("请求失败, 请确认添加了url或者参数错误"))
		return
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		golog.Error(err)
		w.Write([]byte("请求失败, 请确认添加了url或者参数错误"))
		return
	}
	defer response.Body.Close()
	send, err := ioutil.ReadAll(response.Body)
	if err != nil {
		golog.Error(err)
		w.Write([]byte("获取数据失败"))
	}
	//fmt.Println(string(xx))
	//send, _ := json.Marshal(sl)
	//
	w.Write(send)
	return

}
