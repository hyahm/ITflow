package handle

import (
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"itflow/bug/public"
	"encoding/json"
	"fmt"
	"github.com/hyahm/golog"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	//"strconv"
)

type statusList struct {
	StatusList []string `json:"statuslist"`
	Code       int      `json:"statuscode"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}

		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl := &statusList{}
		for _, v := range bugconfig.CacheSidStatus {
			sl.StatusList = append(sl.StatusList, v)
		}

		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type mystatus struct {
	CheckStatus []string `json:"checkstatus"`
	Code        int      `json:"statuscode"`
}

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl := &mystatus{}
		// 遍历出每一个status id
		for _, v := range strings.Split(bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]], ",") {
			sid, _ := strconv.Atoi(v)
			//判断这个id是否存在
			if value, ok := bugconfig.CacheSidStatus[int64(sid)]; ok {
				sl.CheckStatus = append(sl.CheckStatus, value)
			}
		}
		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GetPermStatus(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl := &statusList{}
		//如果是管理员的话,所有的都可以
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			for _, v := range bugconfig.CacheSidStatus {
				sl.StatusList = append(sl.StatusList, v)
			}
		}
		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type userinfo struct {
	Code     int    `json:"statuscode"`
	Nickname string `json:"nickname"`
	Realname string `json:"realname"`
	Email    string `json:"email"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, name, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl := &userinfo{}

		err = conn.GetOne("select email,realname from user where nickname=?", name).Scan(&sl.Email, &sl.Realname)

		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		sl.Nickname = name
		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateInfo(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, name, err := logtokenmysql(r)
		errorcode := &errorstruct{}

		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl := &userinfo{}

		ss, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(ss, sl)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		uid := bugconfig.CacheNickNameUid[name]
		// 修改用户信息
		_, err = conn.Update("update user set email=?,realname=?,nickname=? where id=?", sl.Email, sl.Realname, sl.Nickname, uid)

		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		oldrealname := bugconfig.CacheUidRealName[int64(uid)]
		oldnickname := bugconfig.CacheUidNickName[int64(uid)]
		// 更新缓存
		delete(bugconfig.CacheNickNameUid, oldnickname)
		delete(bugconfig.CacheRealNameUid, oldrealname)
		delete(bugconfig.CacheUidEmail, uid)
		bugconfig.CacheNickNameUid[sl.Nickname] = int64(uid)
		bugconfig.CacheRealNameUid[sl.Realname] = int64(uid)
		bugconfig.CacheUidEmail[uid] = sl.Email

		bugconfig.CacheUidRealName[int64(uid)] = sl.Realname

		bugconfig.CacheUidNickName[int64(uid)] = sl.Nickname

		err = insertlog(conn, "updateinfo", name+"修改了用户信息", r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		send, _ := json.Marshal(sl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateRoles(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		sl := &model.Table_roles{}

		ss, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}

		err = json.Unmarshal(ss, sl)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		var rid int64 // 这个是修改后的rid
		for k, v := range bugconfig.CacheRidGroup {
			if v == sl.Name {
				rid = k
				break
			}
		}
		if rid == 0 {
			w.Write(errorcode.ErrorGetData())
			return
		}

		_, err = conn.Update("update user set rid=? where id=?", rid, sl.Id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		err = insertlog(conn, "updaterole", nickname+"修改了角色权限", r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type logstruct struct {
	Id       int    `json:"id"`
	Exectime int64  `json:"exectime"`
	Classify string `json:"classify"`
	Content  string `json:"content"`
	Ip       string `json:"ip"`
}

type loglist struct {
	LogList []*logstruct `json:"loglist"`
	Code    int          `json:"statuscode"`
}

func LogList(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		alllog := &loglist{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("log", nickname, conn)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		dsql := "select id,exectime,classify,content,ip from log order by id desc"
		rows, err := conn.GetRows(dsql)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			log := &logstruct{}
			rows.Scan(&log.Id, &log.Exectime, &log.Classify, &log.Content, &log.Ip)
			alllog.LogList = append(alllog.LogList, log)
		}
		send, _ := json.Marshal(alllog)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func SearchLog(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		alllog := &model.Search_log{}
		listlog := &model.List_log{}
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}

		err = json.Unmarshal(bytedata, alllog)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("log", nickname, conn)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		basesql := "select id,exectime,classify,content,ip from log "
		endsql := ""
		// 如果搜索了类别
		if alllog.Classify != "" {
			//判断是否在类别数组中
			var realclassify bool
			for _, v := range bugconfig.CLASSIFY {
				if v == alllog.Classify {
					realclassify = true
					break
				}
			}
			if !realclassify {
				w.Write(errorcode.ErrorKeyNotFound())
				return
			}
			endsql = fmt.Sprintf("where classify='%v' ", alllog.Classify)
		}
		// 如果有时间选择，并且不为0
		if alllog.StartTime != 0 {
			if len(endsql) == 0 {
				endsql = fmt.Sprintf("where exectime between %d and %d ", alllog.StartTime, alllog.EndTime)
			} else {
				endsql += fmt.Sprintf(" and exectime between %d and %d ", alllog.StartTime, alllog.EndTime)
			}
		}
		//获取总行数
		fmt.Println("endsql:", endsql)
		err = conn.GetOne("select count(id) from log " + endsql).Scan(&listlog.Count)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		start, end := public.GetPagingLimitAndPage(listlog.Count, alllog.Page, alllog.Limit)
		listlog.Page = start / alllog.Limit

		fmt.Println("------", start, "------", end)
		rows, err := conn.GetRows(basesql+endsql+" limit ?,?", start, end)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		for rows.Next() {
			one := &model.Table_log{}
			//basesql := "select id,exectime,classify,content,ip from log "
			rows.Scan(&one.Id, &one.Exectime, &one.Classify, &one.Content, &one.Ip)
			listlog.LogList = append(listlog.LogList, one)
		}

		send, _ := json.Marshal(listlog)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func LogClassify(w http.ResponseWriter, r *http.Request) {

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		send, _ := json.Marshal(bugconfig.CLASSIFY)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type getBugSearchParam struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Level   string   `json:"level"`
	Project string   `json:"project"`
	Title   string   `json:"title"`
	Status  []string `json:"status"`
}

type getAllBugSearchParam struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Level   string   `json:"level"`
	Project string   `json:"project"`
	Title   string   `json:"title"`
	Status  string   `json:"status"`
	Total   int      `json:"total"`
	Handle  []string `json:"handle"`
}

type getChangeStatus struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Code   int    `json:"statuscode"`
}

func ChangeBugStatus(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		param := &getChangeStatus{}

		searchq, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}

		err = json.Unmarshal(searchq, param)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var sid int64
		var ok bool
		if sid, ok = bugconfig.CacheStatusSid[param.Status]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}

		basesql := "update bugs set sid=? where id=?"

		_, err = conn.Update(basesql, sid, param.Id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "status",
		}
		err = il.Update(
			param.Id, nickname, param.Status)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		send, _ := json.Marshal(param)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func ChangeFilterStatus(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		param := &mystatus{}

		searchq, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(searchq, param)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		fmt.Printf("%+v \n", param.CheckStatus)
		ids := make([]string, 0)
		for _, v := range param.CheckStatus {
			// 如果存在这个状态才添加进来
			if value, ok := bugconfig.CacheStatusSid[v]; ok {
				ids = append(ids, strconv.FormatInt(value, 10))
			}

		}
		showstatus := strings.Join(ids, ",")
		//
		basesql := "update user set showstatus=? where id=?"

		_, err = conn.Update(basesql, showstatus, bugconfig.CacheNickNameUid[nickname])
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		//更新缓存
		bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]] = showstatus

		send, _ := json.Marshal(param)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GetAllBugs(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		al := &model.AllArticleList{}

		searchq, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			al.Code = 7
			send, _ := json.Marshal(al)
			w.Write(send)
			return
		}
		searchparam := &getBugSearchParam{}
		err = json.Unmarshal(searchq, searchparam)
		if err != nil {
			golog.Error(err.Error())
			al.Code = 5
			send, _ := json.Marshal(al)
			w.Write(send)
			return
		}

		err = conn.GetOne("select count(id) from bugs").Scan(&al.Count)
		if err != nil {
			golog.Error(err.Error())
			al.Code = 5
			send, _ := json.Marshal(al)
			w.Write(send)
			return
		}
		start, end := public.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

		alsql := "select id,createtime,importent,status,bugtitle,uid,level,pid,env,spusers from bugs where dustbin=0 order by id desc limit ?,?"
		rows, err := conn.GetRows(alsql, start, end)
		if err != nil {
			golog.Error(err.Error())
			al.Code = 1
			send, _ := json.Marshal(al)
			w.Write(send)
			return
		}
		for rows.Next() {
			bl := &model.ArticleList{}
			var statusid int64
			var uid int64
			var pid int64
			var eid int64
			var spusers string
			rows.Scan(&bl.ID, &bl.Date, &bl.Importance, &statusid, &bl.Title, &uid, &bl.Level, &pid, &eid, &spusers)

			bl.Handle = formatUserlistToRealname(spusers)
			bl.Status = bugconfig.CacheSidStatus[statusid]
			bl.Author = bugconfig.CacheUidRealName[uid]
			bl.Projectname = bugconfig.CachePidName[pid]
			bl.Env = bugconfig.CacheEidName[eid]

			al.Al = append(al.Al, bl)

		}
		send, _ := json.Marshal(al)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func GetMyBugs(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, name, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		al := &model.AllArticleList{}

		searchq, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		searchparam := &getBugSearchParam{}
		err = json.Unmarshal(searchq, searchparam)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		uid := bugconfig.CacheNickNameUid[name]
		err = conn.GetOne("select count(id) from bugs where uid=?", uid).Scan(&al.Count)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		start, end := public.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

		alsql := "select id,createtime,importent,status,bugtitle,uid,level,pid,env,spusers from bugs where uid=? and dustbin=0 order by id desc limit ?,?"
		rows, err := conn.GetRows(alsql, uid, start, end)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			bl := &model.ArticleList{}
			var statusid int64
			var uid int64
			var pid int64
			var eid int64
			rows.Scan(&bl.ID, &bl.Date, &bl.Importance, &statusid, &bl.Title, &uid, &bl.Level, &pid, &eid, &bl.Handle)

			bl.Status = bugconfig.CacheSidStatus[statusid]
			bl.Author = bugconfig.CacheUidRealName[uid]
			bl.Projectname = bugconfig.CachePidName[pid]
			bl.Env = bugconfig.CacheEidName[eid]

			al.Al = append(al.Al, bl)

		}
		send, _ := json.Marshal(al)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func CloseBug(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodGet {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		id := r.FormValue("id")
		var uid int64
		err = conn.GetOne("select uid from bugs where id=?", id).Scan(&uid)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if uid != bugconfig.CacheNickNameUid[nickname] && uid != bugconfig.SUPERID {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		_, err = conn.Update("update bugs set dustbin=true where id=?", id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "bug",
		}
		err = il.Del(id, nickname)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type editList struct {
	Importance  string   `json:"importance"`
	Title       string   `json:"title"`
	Level       string   `json:"level"`
	Version     string   `json:"version"`
	Projectname string   `json:"projectname"`
	Env         string   `json:"env"`
	Handle      []string `json:"handle"`
	Content     string   `json:"content"`
	Code        int      `json:"statuscode"`
}

func BugEdit(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodGet {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		al := &editList{}

		id := r.FormValue("id")
		var pid int64
		var eid int64
		var uidlist string
		var iid int64
		var lid int64
		var vid int64
		alsql := "select iid,bugtitle,lid,pid,eid,spusers,vid,content from bugs where id=?"
		err = conn.GetOne(alsql, id).Scan(&iid, &al.Title, &lid, &pid, &eid, &uidlist, &vid, &al.Content)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		al.Importance = bugconfig.CacheIidImportant[iid]
		al.Level = bugconfig.CacheLidLevel[lid]
		al.Version = bugconfig.CacheVidName[vid]
		al.Content = html.UnescapeString(al.Content)
		al.Handle = formatUserlistToShow(uidlist)
		al.Projectname = bugconfig.CachePidName[pid]
		al.Env = bugconfig.CacheEidName[eid]

		send, _ := json.Marshal(al)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
