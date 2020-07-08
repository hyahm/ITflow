package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/internal/role"
	"itflow/internal/search"
	"itflow/internal/status"
	"itflow/internal/user"
	"itflow/model"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
	//"strconv"
)

type statusList struct {
	StatusList []cache.Status `json:"statuslist"`
	Code       int            `json:"code"`
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

	uid := xmux.GetData(r).Get("uid").(int64)

	sl := &statusList{}
	//如果是管理员的话,所有的都可以
	sl.StatusList = cache.CacheUidFilter[uid].ToShow()
	// if  == cache.SUPERID {
	// 	for _, v := range cache.CacheSidStatus {
	// 		sl.StatusList = append(sl.StatusList, v.ToString())
	// 	}
	// }
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

	// errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	// mybug.GetUsefulCondition(uid)
	countsql := "select count(id) from bugs where dustbin=0 and uid=? "
	searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs where dustbin=0 and uid=? "
	sch, err := mybug.GetUsefulCondition(uid, countsql, searchsql)
	if err != nil {
		if err == search.ErrorNoStatus {
			al := &model.AllArticleList{
				Al: make([]*model.ArticleList, 0),
			}
			w.Write(al.Marshal())
			return
		}
		golog.Error(err)
		al := &model.AllArticleList{
			Al:   make([]*model.ArticleList, 0),
			Code: 1,
			Msg:  err.Error(),
		}
		w.Write(al.Marshal())
		return
	}
	w.Write(sch.GetMyBugs())

}

func ResumeBug(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	golog.Info("0000")
	golog.Info(id)
	errorcode := &response.Response{}
	bug := &model.Bug{}
	err := bug.Resume(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info("0000")
	w.Write(errorcode.Success())
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