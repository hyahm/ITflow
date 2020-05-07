package handle

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"itflow/app/bugconfig"
	"itflow/app/public"
	"itflow/db"
	"itflow/model"
	"itflow/network/bug"
	"itflow/network/datalog"
	"itflow/network/response"
	"itflow/network/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
	//"strconv"
)

type statusList struct {
	StatusList []string `json:"statuslist"`
	Code       int      `json:"code"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {

	sl := &statusList{}
	for _, v := range bugconfig.CacheSidStatus {
		sl.StatusList = append(sl.StatusList, v)
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}

type mystatus struct {
	CheckStatus []string `json:"checkstatus"`
	Code        int      `json:"code"`
}

func ShowStatus(w http.ResponseWriter, r *http.Request) {

	sl := &mystatus{}

	nickname := xmux.GetData(r).Get("nickname").(string)
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

func GetPermStatus(w http.ResponseWriter, r *http.Request) {

	nickname := xmux.GetData(r).Get("nickname").(string)

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

func GetInfo(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := &user.UserInfo{}
	sl.NickName = xmux.GetData(r).Get("nickname").(string)
	row, err := db.Mconn.GetOne("select email,realname from user where nickname=?", sl.NickName)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&sl.Email, &sl.Realname)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}

func UpdateInfo(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := xmux.GetData(r).Data.(*user.UserInfo)
	nickname := xmux.GetData(r).Get("nickname").(string)

	uid := bugconfig.CacheNickNameUid[nickname]
	// 修改用户信息
	_, err := db.Mconn.Update("update user set email=?,realname=?,nickname=? where id=?", sl.Email, sl.Realname, sl.NickName, uid)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	oldrealname := bugconfig.CacheUidRealName[int64(uid)]
	oldnickname := bugconfig.CacheUidNickName[int64(uid)]
	// 更新缓存
	delete(bugconfig.CacheNickNameUid, oldnickname)
	delete(bugconfig.CacheRealNameUid, oldrealname)
	delete(bugconfig.CacheUidEmail, uid)
	bugconfig.CacheNickNameUid[sl.NickName] = int64(uid)
	bugconfig.CacheRealNameUid[sl.Realname] = int64(uid)
	bugconfig.CacheUidEmail[uid] = sl.Email

	bugconfig.CacheUidRealName[int64(uid)] = sl.Realname

	bugconfig.CacheUidNickName[int64(uid)] = sl.NickName

	err = insertlog("updateinfo", nickname+"修改了用户信息", r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}

func UpdateRoles(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	sl := &model.Table_roles{}

	ss, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = json.Unmarshal(ss, sl)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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
		golog.Debug("不存在此权限")
		w.Write(errorcode.Error("不存在此权限"))
		return
	}

	_, err = db.Mconn.Update("update user set rid=? where id=?", rid, sl.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	err = insertlog("updaterole", nickname+"修改了角色权限", r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	w.Write(errorcode.Success())
	return

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
	Code   int    `json:"code"`
}

func ChangeBugStatus(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	param := &getChangeStatus{}

	searchq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = json.Unmarshal(searchq, param)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var sid int64
	var ok bool
	if sid, ok = bugconfig.CacheStatusSid[param.Status]; !ok {
		golog.Errorf("找不到status id: %s", param.Status)
		w.Write(errorcode.Errorf("找不到status id: %s", param.Status))
		return
	}

	basesql := "update bugs set sid=? where id=?"

	_, err = db.Mconn.Update(basesql, sid, param.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "status",
		Action:   "change",
	}

	send, _ := json.Marshal(param)
	w.Write(send)
	return

}

func ChangeFilterStatus(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	param := &mystatus{}

	searchq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(searchq, param)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
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
	nickname := xmux.GetData(r).Get("nickname").(string)
	_, err = db.Mconn.Update(basesql, showstatus, bugconfig.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更新缓存
	bugconfig.CacheUidFilter[bugconfig.CacheNickNameUid[nickname]] = showstatus

	send, _ := json.Marshal(param)
	w.Write(send)
	return
}

func GetMyBugs(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	al := &model.AllArticleList{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := bugconfig.CacheNickNameUid[nickname]
	row, err := db.Mconn.GetOne("select count(id) from bugs where uid=?", uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	searchparam := xmux.GetData(r).Data.(*bug.SearchParam)
	start, end := public.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

	alsql := "select id,createtime,importent,status,bugtitle,uid,level,pid,env,spusers from bugs where uid=? and dustbin=0 order by id desc limit ?,?"
	rows, err := db.Mconn.GetRows(alsql, uid, start, end)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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

func CloseBug(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	var uid int64
	row, err := db.Mconn.GetOne("select uid from bugs where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	if uid != bugconfig.CacheNickNameUid[nickname] && uid != bugconfig.SUPERID {
		golog.Debug("没有权限")
		w.Write(errorcode.Error("没有权限"))
		return
	}
	_, err = db.Mconn.Update("update bugs set dustbin=true where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "bug",
		Action:   "close",
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

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
	Code        int      `json:"code"`
}

func BugEdit(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	al := &editList{}

	id := r.FormValue("id")
	var pid int64
	var eid int64
	var uidlist string
	var iid int64
	var lid int64
	var vid int64
	alsql := "select iid,bugtitle,lid,pid,eid,spusers,vid,content from bugs where id=?"
	row, err := db.Mconn.GetOne(alsql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&iid, &al.Title, &lid, &pid, &eid, &uidlist, &vid, &al.Content)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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
