package handle

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/model"
	"itflow/response"
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
	uid := xmux.GetInstance(r).Get("uid").(int64)
	status, err := model.GetStatusIDsByUid(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: status,
	}
	w.Write(res.Marshal())
}

func GetPermStatus(w http.ResponseWriter, r *http.Request) {
	// 获取可以改变的状态
	sl := &statusList{
		StatusList: make([]string, 0),
	}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	var err error
	var rows *sql.Rows
	if uid == cache.SUPERID {
		rows, err = db.Mconn.GetRows(`select name from status`)
		if err != nil {
			golog.Error(err)
			w.Write(sl.ErrorE(err))
			return
		}

	} else {
		var sids string
		err := db.Mconn.GetOne("select sids from statusgroup where id = (select sgid from jobs where id=(select jid from user where id=?))", uid).Scan(&sids)
		if err != nil {
			golog.Error(err)
			w.Write(sl.ErrorE(err))
			return
		}

		rows, err = db.Mconn.GetRowsIn(`select name from status where id in (?)`,
			strings.Split(sids, ","))
		if err != nil {
			golog.Error(err)
			w.Write(sl.ErrorE(err))
			return
		}

	}
	statusname := new(string)
	for rows.Next() {
		err = rows.Scan(statusname)
		if err != nil {
			golog.Info(err)
			continue
		}
		sl.StatusList = append(sl.StatusList, *statusname)
	}
	rows.Close()
	w.Write(sl.Marshal())

}

// func GetInfo(w http.ResponseWriter, r *http.Request) {

// 	errorcode := &response.Response{}

// 	sl := &user.UserInfo{}
// 	sl.NickName = xmux.GetInstance(r).Get("nickname").(string)
// 	err := db.Mconn.GetOne("select email,realname from user where nickname=?", sl.NickName).Scan(&sl.Email, &sl.Realname)
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.ErrorE(err))
// 		return
// 	}

// 	send, _ := json.Marshal(sl)
// 	w.Write(send)

// }

// func UpdateInfo(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &response.Response{}
// 	sl := xmux.GetInstance(r).Data.(*user.UserInfo)
// 	uid := xmux.GetInstance(r).Get("uid").(int64)
// 	// 修改用户信息
// 	_, err := db.Mconn.Update("update user set email=?,realname=?,nickname=? where id=?", sl.Email, sl.Realname, sl.NickName, uid)
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.ErrorE(err))
// 		return
// 	}

// 	send, _ := json.Marshal(sl)
// 	w.Write(send)
// 	return

// }

func ChangeBugStatus(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	param := xmux.GetInstance(r).Data.(*bug.ChangeStatus)

	// sid := param.Status.Id()
	// if sid == 0 {
	// 	golog.Errorf("找不到status id: %s", param.Status)
	// 	w.Write(errorcode.Errorf("找不到status id: %s", param.Status))
	// 	return
	// }

	basesql := "update bugs set sid=(select id from status where name=?),updatetime=? where id=?"

	_, err := db.Mconn.Update(basesql, param.Status, time.Now().Unix(), param.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(param)
	w.Write(send)
}

func ChangeShowStatus(w http.ResponseWriter, r *http.Request) {
	// 更新自己的显示状态
	user := xmux.GetInstance(r).Data.(*model.User)
	m := make(map[int64]struct{})
	for _, v := range user.ShowStatus {
		if v <= 0 {
			continue
		}
		m[v] = struct{}{}
	}
	user.ShowStatus = make([]int64, 0, len(m))
	for k := range m {
		user.ShowStatus = append(user.ShowStatus, k)
	}
	user.ID = xmux.GetInstance(r).Get("uid").(int64)
	err := user.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func GetMyBugs(w http.ResponseWriter, r *http.Request) {
	// uid := xmux.GetInstance(r).Get("uid").(int64)
	// mybug := xmux.GetInstance(r).Data.(*search.ReqMyBugFilter)
	// golog.Infof("%+v", *mybug)
	// // mybug.GetUsefulCondition(uid)
	// al := &model.AllArticleList{
	// 	Al:   make([]*model.ArticleList, 0),
	// 	Page: 1,
	// }
	// countsql := "select count(id) from bugs where dustbin=true and uid=? and sid in (select id from status where name in (?))"
	// err := db.Mconn.GetOneIn(countsql, uid, mybug.ShowsStatus).Scan(&al.Count)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(al.ErrorE(err))
	// 	return
	// }
	// page, start, end := xmux.GetLimit(al.Count, mybug.Page, mybug.Limit)
	// al.Page = page
	// // searchsql := "select id,createtime,iid,sid,title,lid,pid,eid,spusers from bugs join  on dustbin=true and uid=? "
	// searchsql := `select b.id,b.createtime,i.name,s.name,title,l.name,p.name,e.name,spusers,u.realname from bugs as b
	// join importants as i
	// join status as s
	// join level as l
	// join project as p
	// join environment as e
	// join user as u
	// on dustbin=true and b.iid = i.id and b.sid = s.id and b.lid = l.id and b.pid=p.id and b.eid = e.id and b.uid=u.id and uid=? limit ?,?`
	// rows, err := db.Mconn.GetRows(searchsql, uid, start, end)
	// // sch, err := mybug.GetUsefulCondition(uid,
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(al.ErrorE(err))
	// 	return
	// }

	// for rows.Next() {
	// 	bug := &model.ArticleList{
	// 		Handle: make([]string, 0),
	// 	}
	// 	var ids string
	// 	err = rows.Scan(&bug.ID,
	// 		&bug.Date, &bug.Importance, &bug.Status, &bug.Title, &bug.Level, &bug.Projectname,
	// 		&bug.Env, &ids, &bug.Author)
	// 	if err != nil {
	// 		golog.Info(err)
	// 		continue
	// 	}
	// 	realnames, err := db.Mconn.GetRows("select realname from user where id in (?)",
	// 		strings.Split(ids, ","))
	// 	if err != nil {
	// 		golog.Error(err)
	// 		w.Write(al.ErrorE(err))
	// 		return
	// 	}
	// 	for realnames.Next() {
	// 		var name string
	// 		err = realnames.Scan(&name)
	// 		if err != nil {
	// 			golog.Error(err)
	// 			continue
	// 		}
	// 		bug.Handle = append(bug.Handle, name)
	// 	}

	// 	al.Al = append(al.Al, bug)
	// }
	// rows.Close()
	// golog.Info(string(al.Marshal()))
	// w.Write(al.Marshal())
	// if err != nil {
	// 	if err == search.ErrorNoStatus {
	// 		al := &model.AllArticleList{
	// 			Al: make([]*model.ArticleList, 0),
	// 		}
	// 		w.Write(al.Marshal())
	// 		return
	// 	}
	// 	golog.Error(err)
	// 	al := &model.AllArticleList{
	// 		Al:   make([]*model.ArticleList, 0),
	// 		Code: 1,
	// 		Msg:  err.Error(),
	// 	}
	// 	w.Write(al.Marshal())
	// 	return
	// }
	// w.Write(sch.GetMyBugs())

}

func DeleteBug(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)
	bug := &model.Bug{}
	err := bug.Delete(uid, id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func ResumeBug(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	errorcode := &response.Response{}
	bug := &model.Bug{}
	err := bug.Resume(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	w.Write(errorcode.Success())
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

	thisUid := xmux.GetInstance(r).Get("uid").(int64)
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
}
