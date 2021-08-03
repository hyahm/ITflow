package handle

import (
	"encoding/json"
	"fmt"
	"html"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/project"
	"net/http"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
	//"strings"
)

type projectList struct {
	ProjectList []string `json:"projectlist"`
	Code        int      `json:"code"`
}

func GetProject(w http.ResponseWriter, r *http.Request) {

	pl := &projectList{}

	// for _, v := range cache.CachePidProject {
	// 	pl.ProjectList = append(pl.ProjectList, v.ToString())
	// }
	send, _ := json.Marshal(pl)
	w.Write(send)
	return

}

func GetMyProject(w http.ResponseWriter, r *http.Request) {
	myproject := &project.MyProject{
		Name: make([]string, 0),
	}
	uid := xmux.GetInstance(r).Get("uid").(int64)

	w.Write(myproject.Get(uid))
	return

}

// 添加或编辑
func BugCreate(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	data := xmux.GetInstance(r).Data.(*bug.EditBug)
	var ids = make([]string, 0)
	rows, err := db.Mconn.GetRowsIn("select id from user where realname in (?)",
		data.Selectusers)
	if err != nil {
		golog.Error(err)
		w.Write(data.ErrorE(err))
		return
	}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			golog.Info(err)
			continue
		}
		ids = append(ids, id)
	}
	rows.Close()
	if data.Id <= 0 {
		// 插入bug
		// err = bug.CreateBug()
		// if err != nil {
		// 	w.Write(errorcode.ErrorE(err))
		// 	return
		// }
		golog.Info(cache.DefaultCreateSid)
		if cache.DefaultCreateSid == 0 {
			w.Write(data.Error("必须先设置一个默认的创建状态（系统设置->默认值->bug创建时的状态）"))
			return
		}
		createsql := `insert into bugs(tid,uid,title,sid,content,iid,createtime,vid,spusers,lid,eid,pid) 
		values(?,?,?,?,?,(select ifnull(min(id),0) from importants where name=?),?,
		(select ifnull(min(id),0) from version where name=?),?,
		(select ifnull(min(id),0) from level where name=?),
		(select ifnull(min(id),0) from environment where name=?),
		(select ifnull(min(id),0) from project where name=?))`
		golog.Info(data.Typ)
		data.Id, err = db.Mconn.Insert(createsql, data.Typ, uid, data.Title,
			cache.DefaultCreateSid, html.EscapeString(data.Content), data.Important,
			time.Now().Unix(), data.Version, strings.Join(ids, ","), data.Level, data.Envname, data.Projectname)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}

		// 根据处理人的id找出邮箱地址
		idrows, err := db.Mconn.GetRowsIn("select email from user where id in (?)", ids)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}
		toEmails := make([]string, 0)
		for idrows.Next() {
			var et string
			idrows.Scan(&et)
			toEmails = append(toEmails, et)
		}
		bugUrl := fmt.Sprintf("%s/showbug/%d", r.Referer(), data.Id)
		content := fmt.Sprintf(`<html><body><h1>%s<h1>bug地址:<a href="%s">%s</a></body></html>`, data.Title, bugUrl, bugUrl)
		cache.CacheEmail.SendMail("您有一个新的bug需要处理",
			content,
			toEmails...)
	} else {
		// update
		permsql := "select uid from bugs where id=?"
		var bugUid int64
		err := db.Mconn.GetOne(permsql, data.Id).Scan(&bugUid)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}
		if bugUid != uid && uid != cache.SUPERID {
			w.Write(data.Error("没有权限"))
			return
		}
		updatesql := `update bugs set tid=?, title=?,content=?,
			iid=(select ifnull(min(id),0) from importants where name=?),
			updatetime=?,
			vid=(select ifnull(min(id),0) from version where name=?),
			spusers=?,
			lid=(select ifnull(min(id),0) from level where name=?),
			eid=(select ifnull(min(id),0) from environment where name=?),
			pid=(select ifnull(min(id),0) from project where name=?) 
		where id=?`
		_, err = db.Mconn.Update(updatesql, data.Typ, data.Title,
			html.EscapeString(data.Content), data.Important,
			time.Now().Unix(), data.Version, strings.Join(ids, ","),
			data.Level, data.Envname, data.Projectname, data.Id)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}

		// err = bug.EditBug()
		// if err != nil {
		// 	w.Write(errorcode.ErrorE(err))
		// 	return
		// }
		// go datalog.InsertLog("bug", nickname+fmt.Sprintf(" update bug id: %d", data.Id), r.RemoteAddr, nickname, "update")

		// 根据处理人的id找出邮箱地址
		idrows, err := db.Mconn.GetRowsIn("select email from user where id in (?)", ids)
		if err != nil {
			golog.Error(err)
			w.Write(data.ErrorE(err))
			return
		}
		toEmails := make([]string, 0)
		for idrows.Next() {
			var et string
			idrows.Scan(&et)
			toEmails = append(toEmails, et)
		}
		bugUrl := fmt.Sprintf("%s/showbug/%d", r.Referer(), data.Id)
		content := fmt.Sprintf(`<html><body><h1>%s<h1>bug地址:<a href="%s">%s</a></body></html>`, data.Title, bugUrl, bugUrl)
		golog.Info(content)
		cache.CacheEmail.SendMail("您有一个的bug需要处理",
			content,
			toEmails...)
	}

	w.Write(data.Marshal())
	return

}
