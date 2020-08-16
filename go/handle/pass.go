package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/assist"
	"itflow/internal/bug"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func PassBug(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	ub := xmux.GetData(r).Data.(*bug.PassBug)
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)
	// 获取参数

	// 判断用户是否能处理这个project
	pid, ok := cache.CacheProjectPid[ub.ProjectName]
	if !ok {
		golog.Error("not found project")
		w.Write(errorcode.Error("not found project"))
		return
	}
	var ugid int64
	err := db.Mconn.GetOne("select ugid from project where id=?", pid).Scan(&ugid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 判断这个bug是不是自己的任务，只有自己的任务才可以转交
	var havePerm bool
	for _, v := range strings.Split(cache.CacheUGidUserGroup[ugid].Uids, ",") {
		permuid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			continue
		}
		if permuid == uid {
			havePerm = true
			break
		}
	}
	if !havePerm {
		golog.Error("you have not permssion")
		w.Write(errorcode.Error("you have not permssion"))
		return
	}

	var splist string
	var hasperm bool
	err = db.Mconn.GetOne("select spusers from bugs where id=? ", ub.Id).Scan(&splist)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	myuid := strconv.FormatInt(cache.CacheNickNameUid[nickname], 10)
	for _, v := range strings.Split(splist, ",") {
		if myuid == v {
			hasperm = true
			break
		}
	}
	if !hasperm {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 判断状态是否存在
	sid := ub.Status.Id()
	if sid == 0 {
		w.Write(errorcode.Error("没有status"))
		return
	}
	idstr := make([]string, 0)
	mails := make([]string, 0)
	for _, v := range ub.SelectUsers {
		var thisUid int64
		if thisUid, ok = cache.CacheRealNameUid[v]; !ok {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		idstr = append(idstr, strconv.FormatInt(thisUid, 10))
		mails = append(mails, cache.CacheUidEmail[thisUid])
	}

	ul := strings.Join(idstr, ",")

	remarksql := "insert into informations(uid,bid,info,time) values(?,?,?,?)"
	_, err = db.Mconn.Insert(remarksql, cache.CacheNickNameUid[nickname], ub.Id, ub.Remark, time.Now().Unix())
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//更改bug

	_, err = db.Mconn.Update("update bugs set sid=?,spusers=?,updatetime=? where id=?", sid, ul, time.Now().Unix(), ub.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if cache.CacheEmail.Enable {
		go cache.CacheEmail.SendMail("转让bug", fmt.Sprintf("由%s 转交给你", cache.CacheUidRealName[uid]), mails...)
	}

	go datalog.InsertLog("bug", fmt.Sprintf("bug id: %v", ub.Id),
		r.RemoteAddr, nickname, "pass")

	send, _ := json.Marshal(ub)
	w.Write(send)
	return

}

func TaskList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	al := &model.AllArticleList{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	uid := cache.CacheNickNameUid[nickname]

	getaritclesql := "select id,createtime,importent,status,title,uid,level,pid,spusers from bugs where id in (select bid from userandbug where uid=?)  order by id desc "

	rows, err := db.Mconn.GetRows(getaritclesql, uid)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		sendlist := &model.ArticleList{}
		var statusid cache.StatusId
		var spusers string
		var uid int64
		var pid cache.ProjectId
		rows.Scan(&sendlist.ID, &sendlist.Date, &sendlist.Importance, &statusid, &sendlist.Title, &uid, &sendlist.Level, &pid, &spusers)
		sendlist.Handle = assist.FormatUserlistToShow(spusers)
		sendlist.Status = cache.CacheSidStatus[statusid]

		sendlist.Author = cache.CacheUidRealName[uid]
		sendlist.Projectname = cache.CachePidProject[pid]

		al.Al = append(al.Al, sendlist)
	}
	send, _ := json.Marshal(al)
	w.Write(send)
	return

}
