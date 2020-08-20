package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/response"
	"itflow/internal/role"
	"itflow/internal/search"
	"itflow/internal/status"
	"itflow/internal/user"
	"itflow/model"
	"net/http"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
	//"strconv"
)

type statusList struct {
	StatusList []string `json:"statuslist"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
}

func (sl *statusList) Marshal() []byte {
	send, err := json.Marshal(sl)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (sl *statusList) Error(msg string) []byte {
	sl.Code = 1
	sl.Msg = msg
	return sl.Marshal()
}

func (sl *statusList) ErrorE(err error) []byte {
	return sl.Error(err.Error())
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	// 获取状态名
	w.Write(status.GetNames())
	return

}

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	// 获取显示的状态名
	// sl := xmux.GetData(r).Data.(*status.Status)
	sl := &status.Status{
		CheckStatus: make([]string, 0),
	}
	uid := xmux.GetData(r).Get("uid").(int64)
	rows, err := db.Mconn.GetRows(`select s.name from user as u join status as s on u.id=? and s.id in (u.showstatus)`, uid)
	if err != nil {
		golog.Error(err)
		w.Write(sl.ErrorE(err))
		return
	}
	for rows.Next() {
		statusname := new(string)
		err = rows.Scan(statusname)
		if err != nil {
			golog.Error(err)
			continue
		}
		sl.CheckStatus = append(sl.CheckStatus, *statusname)
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return
}

func GetPermStatus(w http.ResponseWriter, r *http.Request) {
	// 获取可以改变的状态
	sl := &statusList{
		StatusList: make([]string, 0),
	}
	uid := xmux.GetData(r).Get("uid").(int64)
	rows, err := db.Mconn.GetRows(`select name from status where id in (select s.sids from user as u join statusgroup  as s on u.id=? and u.bugsid=s.id )`, uid)
	if err != nil {
		golog.Error(err)
		w.Write(sl.ErrorE(err))
		return
	}

	for rows.Next() {
		statusname := new(string)
		err = rows.Scan(statusname)
		if err != nil {
			golog.Error(err)
			continue
		}
		sl.StatusList = append(sl.StatusList, *statusname)
	}

	w.Write(sl.Marshal())
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
	uid := xmux.GetData(r).Get("uid").(int64)
	// 修改用户信息
	_, err := db.Mconn.Update("update user set email=?,realname=?,nickname=? where id=?", sl.Email, sl.Realname, sl.NickName, uid)
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
	var rid int64
	err := model.CheckRoleNameInGroup(sl.Name, &rid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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

func ChangeBugStatus(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	param := xmux.GetData(r).Data.(*bug.ChangeStatus)

	sid := param.Status.Id()
	if sid == 0 {
		golog.Errorf("找不到status id: %s", param.Status)
		w.Write(errorcode.Errorf("找不到status id: %s", param.Status))
		return
	}

	basesql := "update bugs set sid=?,updatetime=? where id=?"

	_, err := db.Mconn.Update(basesql, sid, time.Now().Unix(), param.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(param)
	w.Write(send)
	return

}

func ChangeFilterStatus(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	param := xmux.GetData(r).Data.(*status.Status)

	rows, err := db.Mconn.GetRows(fmt.Sprintf("select id from status where name in ('%s')", strings.Join(param.CheckStatus, ",")))
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	sids := make([]string, 0)
	for rows.Next() {
		sid := new(string)
		err = rows.Scan(sid)
		sids = append(sids, *sid)
	}

	uid := xmux.GetData(r).Get("uid").(int64)
	user := &model.User{}
	err = user.UpdateShowStatus(strings.Join(sids, ","), uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	w.Write(errorcode.Success())
	return
}

func GetMyBugs(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	mybug := xmux.GetData(r).Data.(*search.ReqMyBugFilter)
	// mybug.GetUsefulCondition(uid)
	countsql := "select count(id) from bugs where dustbin=true and uid=? "
	// searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs join  on dustbin=true and uid=? "
	searchsql := `select b.id,b.createtime,i.name,s.name,title,l.name,p.name,e.name,spusers,u.realname from bugs as b 
	join importants as i 
	join statusgroup as s 
	join level as l 
	join project as p 
	join environment as e  
	join user as u 
	on dustbin=true and b.iid = i.id and b.sid = s.id and b.lid = l.id and b.pid=p.id and b.eid = e.id and b.uid=u.id and uid=?`
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

	thisUid := xmux.GetData(r).Get("uid").(int64)
	if uid != thisUid && uid != cache.SUPERID {
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

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func BugEdit(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	w.Write(bug.RespEditBugData(id))
	return

}
