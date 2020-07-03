package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/role"
	"itflow/internal/status"
	"itflow/internal/user"
	"itflow/model"
	"itflow/pkg/pager"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
	//"strconv"
)

type statusList struct {
	StatusList []string `json:"statuslist"`
	Code       int      `json:"code"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	// 获取状态名
	w.Write(status.GetNames())
	return

}

func ShowStatus(w http.ResponseWriter, r *http.Request) {

	// sl := xmux.GetData(r).Data.(*status.Status)
	sl := &status.Status{
		CheckStatus: make([]cache.Status, 0),
	}
	uid := xmux.GetData(r).Get("uid").(int64)
	sl.CheckStatus = cache.CacheUidFilter[uid].ToShow()

	// 遍历出每一个status id
	// for _, v := range strings.Split(cache.CacheUidFilter[uid], ",") {
	// 	sid, _ := strconv.Atoi(v)
	// 	//判断这个id是否存在
	// 	if value := cache.CacheSidStatus[cache.StatusId(sid)]; value.ToString() != "" {
	// 		sl.CheckStatus = append(sl.CheckStatus, value)
	// 	}
	// }
	send, _ := json.Marshal(sl)
	w.Write(send)
	return
}

func GetPermStatus(w http.ResponseWriter, r *http.Request) {

	nickname := xmux.GetData(r).Get("nickname").(string)

	sl := &statusList{}
	//如果是管理员的话,所有的都可以
	if cache.CacheNickNameUid[nickname] == cache.SUPERID {
		for _, v := range cache.CacheSidStatus {
			sl.StatusList = append(sl.StatusList, v.ToString())
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
	err := db.Mconn.GetOne("select email,realname from user where nickname=?", sl.NickName).Scan(&sl.Email, &sl.Realname)
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

	uid := cache.CacheNickNameUid[nickname]
	// 修改用户信息
	_, err := db.Mconn.Update("update user set email=?,realname=?,nickname=? where id=?", sl.Email, sl.Realname, sl.NickName, uid)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	oldrealname := cache.CacheUidRealName[int64(uid)]
	oldnickname := cache.CacheUidNickName[int64(uid)]
	// 更新缓存
	delete(cache.CacheNickNameUid, oldnickname)
	delete(cache.CacheRealNameUid, oldrealname)
	delete(cache.CacheUidEmail, uid)
	cache.CacheNickNameUid[sl.NickName] = int64(uid)
	cache.CacheRealNameUid[sl.Realname] = int64(uid)
	cache.CacheUidEmail[uid] = sl.Email

	cache.CacheUidRealName[int64(uid)] = sl.Realname

	cache.CacheUidNickName[int64(uid)] = sl.NickName

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

	sl := xmux.GetData(r).Data.(*role.Role)

	var rid int64 // 这个是修改后的rid
	for k, v := range cache.CacheRidGroup {
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

	_, err := db.Mconn.Update("update user set rid=? where id=?", rid, sl.Id)
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

func ChangeBugStatus(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	param := xmux.GetData(r).Data.(*bug.ChangeStatus)

	sid := param.Status.Id()
	if sid == 0 {
		golog.Errorf("找不到status id: %s", param.Status)
		w.Write(errorcode.Errorf("找不到status id: %s", param.Status))
		return
	}

	basesql := "update bugs set sid=? where id=?"

	_, err := db.Mconn.Update(basesql, sid, param.Id)
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

	param := xmux.GetData(r).Data.(*status.Status)

	showstatus := param.CheckStatus.ToStore()
	//

	uid := xmux.GetData(r).Get("uid").(int64)
	user := &model.User{}
	err := user.UpdateShowStatus(showstatus, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更新缓存
	cache.CacheUidFilter[uid] = showstatus
	send, _ := json.Marshal(param)
	w.Write(send)
	return
}

func GetMyBugs(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	al := &model.AllArticleList{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := cache.CacheNickNameUid[nickname]
	err := db.Mconn.GetOne("select count(id) from bugs where uid=?", uid).Scan(&al.Count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	searchparam := xmux.GetData(r).Data.(*bug.SearchParam)
	start, end := pager.GetPagingLimitAndPage(al.Count, searchparam.Page, searchparam.Limit)

	alsql := "select id,createtime,importent,status,title,uid,level,pid,env,spusers from bugs where uid=? and dustbin=0 order by id desc limit ?,?"
	rows, err := db.Mconn.GetRows(alsql, uid, start, end)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		bl := &model.ArticleList{}
		var statusid cache.StatusId
		var uid int64
		var pid int64
		var eid int64
		rows.Scan(&bl.ID, &bl.Date, &bl.Importance, &statusid, &bl.Title, &uid, &bl.Level, &pid, &eid, &bl.Handle)

		bl.Status = cache.CacheSidStatus[statusid]
		bl.Author = cache.CacheUidRealName[uid]
		bl.Projectname = cache.CachePidName[pid]
		bl.Env = cache.CacheEidName[eid]

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
	err := db.Mconn.GetOne("select uid from bugs where id=?", id).Scan(&uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	nickname := xmux.GetData(r).Get("nickname").(string)
	if uid != cache.CacheNickNameUid[nickname] && uid != cache.SUPERID {
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

func BugEdit(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	w.Write(bug.RespEditBugData(id))
	return

}
