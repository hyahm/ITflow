package handle

import (
	"bug/bugconfig"
	"bug/buglog"
	"encoding/json"
	"galog"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	//"strings"
)

type projectList struct {
	ProjectList []string `json:"projectlist"`
	Code        int      `json:"statuscode"`
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		conn, _, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		pl := &projectList{}

		for _, v := range bugconfig.CachePidName {
			pl.ProjectList = append(pl.ProjectList, v)
		}
		send, _ := json.Marshal(pl)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func BugCreate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		conn, nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			galog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &getArticle{}
		if bugconfig.CacheDefault["status"] <= 0 {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}

		err = json.Unmarshal(content, data)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		// pid
		var pid int64
		var eid int64
		var uid int64
		var iid int64
		var vid int64
		var lid int64
		var ok bool
		if lid, ok = bugconfig.CacheLevelLid[data.Level]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		if pid, ok = bugconfig.CacheProjectPid[data.Projectname]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		//
		if eid, ok = bugconfig.CacheEnvNameEid[data.Envname]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		if uid, ok = bugconfig.CacheNickNameUid[nickname]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		//
		if iid, ok = bugconfig.CacheImportantIid[data.Important]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		if vid, ok = bugconfig.CacheVersionNameVid[data.Version]; !ok {
			w.Write(errorcode.ErrorKeyNotFound())
			return
		}
		ul := make([]string, 0)
		for _, v := range data.Selectusers {
			if udd, ok := bugconfig.CacheRealNameUid[v]; ok {
				ul = append(ul, strconv.FormatInt(udd, 10))
			}
		}
		spusers := strings.Join(ul, ",")
		//spusers, nicknamelist, args := formatUserlistToData(data.Selectusers, data.Id)
		errorcode.UpdateTime = time.Now().Unix()
		// add
		var bugid int64

		//
		if data.Id == -1 {
			// 插入bug

			insertsql := "insert into bugs(uid,bugtitle,sid,content,iid,createtime,lid,pid,eid,spusers,vid) values(?,?,?,?,?,?,?,?,?,?,?)"

			bugid, err = conn.InsertWithID(insertsql,
				uid, data.Title, bugconfig.CacheDefault["status"], html.EscapeString(data.Content),
				iid, errorcode.UpdateTime, lid,
				pid, eid, spusers, vid)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}

			il := buglog.AddLog{
				Conn:     conn,
				Ip:       strings.Split(r.RemoteAddr, ":")[0],
				Classify: "bug",
			}
			err = il.Add(nickname, bugid, data.Title)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}

		} else {
			// update
			// 更新

			insertsql := "update bugs set bugtitle=?,content=?,iid=?,updatetime=?,lid=?,pid=?,eid=?,spusers=?,vid=? where id=?"

			_, err = conn.Update(insertsql, data.Title, html.EscapeString(data.Content), iid,
				time.Now().Unix(), lid, pid, eid, spusers, vid, data.Id)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}

			//// 插入日志
			il := buglog.AddLog{
				Conn:     conn,
				Ip:       strings.Split(r.RemoteAddr, ":")[0],
				Classify: "bug",
			}
			err = il.Update(nickname, bugid, nickname)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return

	}
	w.WriteHeader(http.StatusNotFound)
}
