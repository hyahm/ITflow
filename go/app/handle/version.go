package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"itflow/app/asset"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/model/datalog"
	"itflow/model/response"
	"net/http"
	"strconv"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type addVersion struct {
	Projectname  string `json:'projectname'`
	Platform     string `json:'platform'`
	Version      string `json:'version'`
	Runenv       string `json:'runenv'`
	Iphoneurl    string `json:'iphoneurl'`
	Notiphoneurl string `json:'notiphoneurl'`
}

func AddVersion(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	version_add := &addVersion{}
	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(s, version_add)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("version", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	uid := bugconfig.CacheNickNameUid[nickname]
	add_version_sql := "insert into version(name,urlone,urltwo,createtime,createuid) values(?,?,?,?,?)"

	vid, err := db.Mconn.Insert(add_version_sql, version_add.Version, version_add.Iphoneurl, version_add.Notiphoneurl, time.Now().Unix(), uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "version",
		Action:   "add",
		Msg:      fmt.Sprintf("add version id: %s", version_add.Version),
	}

	// 增加缓存
	bugconfig.CacheVidName[vid] = version_add.Version
	bugconfig.CacheVersionNameVid[version_add.Version] = vid
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type versionInfo struct {
	Id           int    `json:"id"`
	Projectname  string `json:"projectname"`
	Version      string `json:"version"`
	Runenv       string `json:"runenv"`
	Iphoneurl    string `json:"iphoneurl"`
	Notiphoneurl string `json:"notiphoneurl"`
	Date         int    `json:"date"`
}

type versionInfoList struct {
	VersionList []*versionInfo `json:"versionlist"`
	Code        int            `json:"code"`
}

type pageLimit struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func VersionList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	al := &versionInfoList{}

	m, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	pl := &pageLimit{}
	err = json.Unmarshal(m, pl)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("version", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	get_version_sql := "select id,name,urlone,urltwo,createtime from version order by id desc"

	rows, err := db.Mconn.GetRows(get_version_sql)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		rl := &versionInfo{}
		rows.Scan(&rl.Id, &rl.Version, &rl.Iphoneurl, &rl.Notiphoneurl, &rl.Date)
		al.VersionList = append(al.VersionList, rl)
	}

	send, _ := json.Marshal(al)
	w.Write(send)
	return

}

func VersionRemove(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	id := r.FormValue("id")
	var bugcount int
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("version", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	row, err := db.Mconn.GetOne("select count(id) from bugs where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&bugcount)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if bugcount != 0 {
		golog.Errorf("vid:%s has bugs", id)
		w.Write(errorcode.Errorf("vid:%s has bugs", id))
		return
	}
	deletevl := "delete from version where id=?"
	errorcode.Id, err = db.Mconn.Update(deletevl, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	vid, err := strconv.Atoi(id)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "version",
		Action:   "delete",
		Msg:      fmt.Sprintf("delete version id: %s", id),
	}

	delete(bugconfig.CacheVersionNameVid, bugconfig.CacheEidName[int64(vid)])
	delete(bugconfig.CacheVidName, int64(vid))

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type updateversion struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Iphone   string `json:"iphone"`
	NoIphone string `json:"noiphone"`
}

func VersionUpdate(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &response.Response{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	data := &updateversion{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("version", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	if !permssion {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	getdata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(getdata, data)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	uid := bugconfig.CacheNickNameUid[nickname]
	versionsql := "update version set name=?,urlone=?,urltwo=?,createuid=? where id=?"
	_, err = db.Mconn.Update(versionsql, data.Name, data.Iphone, data.NoIphone, uid, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "version",
		Action:   "update",
		Msg:      fmt.Sprintf("update version id %v to %v", data.Id, data.Name),
	}

	delete(bugconfig.CacheVersionNameVid, data.Name)
	bugconfig.CacheVidName[int64(data.Id)] = data.Name
	bugconfig.CacheVersionNameVid[data.Name] = int64(data.Id)

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
