package handle

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/model/datalog"
	"itflow/model/response"
	"net/http"
	"strconv"
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

	_, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	pl := &projectList{}

	for _, v := range bugconfig.CachePidName {
		pl.ProjectList = append(pl.ProjectList, v)
	}
	send, _ := json.Marshal(pl)
	w.Write(send)
	return

}

func BugCreate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &getArticle{}
	if bugconfig.CacheDefault["status"] <= 0 {
		w.Write(errorcode.ErrorE(err))
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	err = json.Unmarshal(content, data)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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
		w.Write(errorcode.Error("没有找到level key"))
		return
	}
	if pid, ok = bugconfig.CacheProjectPid[data.Projectname]; !ok {
		w.Write(errorcode.Error("没有找到project key"))
		return
	}
	//
	if eid, ok = bugconfig.CacheEnvNameEid[data.Envname]; !ok {
		w.Write(errorcode.Error("没有找到env key"))
		return
	}
	if uid, ok = bugconfig.CacheNickNameUid[nickname]; !ok {
		w.Write(errorcode.Error("没有找到nickname key"))
		return
	}
	//
	if iid, ok = bugconfig.CacheImportantIid[data.Important]; !ok {
		w.Write(errorcode.Error("没有找到important key"))
		return
	}
	if vid, ok = bugconfig.CacheVersionNameVid[data.Version]; !ok {
		w.Write(errorcode.Error("没有找到version key"))
		return
	}
	ul := make([]string, 0)
	for _, v := range data.Selectusers {
		start := strings.Index(v, "(")
		end := strings.LastIndex(v, ")")
		v = v[start+1 : end]
		if udd, ok := bugconfig.CacheRealNameUid[v]; ok {
			ul = append(ul, strconv.FormatInt(udd, 10))
		}
	}
	spusers := strings.Join(ul, ",")
	//spusers, nicknamelist, args := formatUserlistToData(data.Selectusers, data.Id)
	errorcode.UpdateTime = time.Now().Unix()
	// add

	//
	if data.Id == -1 {
		// 插入bug

		insertsql := "insert into bugs(uid,bugtitle,sid,content,iid,createtime,lid,pid,eid,spusers,vid) values(?,?,?,?,?,?,?,?,?,?,?)"

		_, err = db.Mconn.Insert(insertsql,
			uid, data.Title, bugconfig.CacheDefault["status"], html.EscapeString(data.Content),
			iid, errorcode.UpdateTime, lid,
			pid, eid, spusers, vid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		xmux.GetData(r).End = &datalog.AddLog{
			Ip:       r.RemoteAddr,
			Username: nickname,
			Classify: "bug",
			Action:   "create",
		}

	} else {
		// update
		// 更新
		insertsql := "update bugs set bugtitle=?,content=?,iid=?,updatetime=?,lid=?,pid=?,eid=?,spusers=?,vid=? where id=?"

		_, err = db.Mconn.Update(insertsql, data.Title, html.EscapeString(data.Content), iid,
			time.Now().Unix(), lid, pid, eid, spusers, vid, data.Id)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

		//// 插入日志
		xmux.GetData(r).End = &datalog.AddLog{
			Ip:       r.RemoteAddr,
			Username: nickname,
			Classify: "bug",
			Action:   "update",
		}

	}

	w.Write(errorcode.Success())
	return

}
