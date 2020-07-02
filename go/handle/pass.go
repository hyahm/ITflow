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
	// 获取参数

	// 判断这个bug是不是自己的任务，只有自己的任务才可以转交
	var splist string
	var hasperm bool
	err := db.Mconn.GetOne("select spusers from bugs where id=? ", ub.Id).Scan(&splist)
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
	var sid int64
	var ok bool
	if sid, ok = cache.CacheStatusSid[ub.Status]; !ok {
		w.Write(errorcode.Error("没有status"))
		return
	}
	idstr := make([]string, 0)
	for _, v := range ub.SelectUsers {
		var uid int64
		if uid, ok = cache.CacheRealNameUid[v]; !ok {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		idstr = append(idstr, strconv.FormatInt(uid, 10))
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

	_, err = db.Mconn.Update("update bugs set sid=?,spusers=? where id=?", sid, ul, ub.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "bug",
		Action:   "pass",
		Msg:      fmt.Sprintf("bug id: %v", ub.Id),
	}

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
		var statusid int64
		var spusers string
		var uid int64
		var pid int64
		rows.Scan(&sendlist.ID, &sendlist.Date, &sendlist.Importance, &statusid, &sendlist.Title, &uid, &sendlist.Level, &pid, &spusers)
		sendlist.Handle = assist.FormatUserlistToShow(spusers)
		sendlist.Status = cache.CacheSidStatus[statusid]

		sendlist.Author = cache.CacheUidRealName[uid]
		sendlist.Projectname = cache.CachePidName[pid]

		al.Al = append(al.Al, sendlist)
	}
	send, _ := json.Marshal(al)
	w.Write(send)
	return

}
