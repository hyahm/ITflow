package handle

import (
	"encoding/json"
	"github.com/hyahm/golog"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"net/http"
	"strconv"
	"strings"
	"time"
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

	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		version_add := &addVersion{}
		s, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(s, version_add)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("version", nickname)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		uid := bugconfig.CacheNickNameUid[nickname]
		add_version_sql := "insert into version(name,urlone,urltwo,createtime,createuid) values(?,?,?,?,?)"

		vid, err := bugconfig.Bug_Mysql.Insert(add_version_sql, version_add.Version, version_add.Iphoneurl, version_add.Notiphoneurl, time.Now().Unix(), uid)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "version",
		}
		err = il.Add(
			nickname, vid, version_add.Version)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加缓存
		bugconfig.CacheVidName[vid] = version_add.Version
		bugconfig.CacheVersionNameVid[version_add.Version] = vid
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
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
	Code        int            `json:"statuscode"`
}

type pageLimit struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func VersionList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		al := &versionInfoList{}

		m, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		pl := &pageLimit{}
		err = json.Unmarshal(m, pl)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("version", nickname)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		get_version_sql := "select id,name,urlone,urltwo,createtime from version order by id desc"

		rows, err := bugconfig.Bug_Mysql.GetRows(get_version_sql)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
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
	w.WriteHeader(http.StatusNotFound)
}

func VersionRemove(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodGet {
		nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
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
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		err = bugconfig.Bug_Mysql.GetOne("select count(id) from bugs where id=?", id).Scan(&bugcount)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		if bugcount != 0 {
			golog.Error("vid:%s has bugs", id)
			w.Write(errorcode.ErrorHasBug())
			return
		}
		deletevl := "delete from version where id=?"
		errorcode.Id, err = bugconfig.Bug_Mysql.Update(deletevl, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		vid, err := strconv.Atoi(id)
		if err != nil {
			w.Write(errorcode.ErrorParams())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "version",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		delete(bugconfig.CacheVersionNameVid, bugconfig.CacheEidName[int64(vid)])
		delete(bugconfig.CacheVidName, int64(vid))

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type updateversion struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Iphone   string `json:"iphone"`
	NoIphone string `json:"noiphone"`
}

func VersionUpdate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == http.MethodPost {
		nickname, err := logtokenmysql(r)
		errorcode := &errorstruct{}
		if err != nil {
			golog.Error(err.Error())
			if err == NotFoundToken {
				w.Write(errorcode.ErrorNotFoundToken())
				return
			}
			w.Write(errorcode.ErrorConnentMysql())
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
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		getdata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(getdata, data)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		uid := bugconfig.CacheNickNameUid[nickname]
		versionsql := "update version set name=?,urlone=?,urltwo=?,createuid=? where id=?"
		_, err = bugconfig.Bug_Mysql.Update(versionsql, data.Name, data.Iphone, data.NoIphone, uid, data.Id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		// 增加日志
		il := buglog.AddLog{
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "version",
		}
		err = il.Update(
			nickname, data.Id, data.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		delete(bugconfig.CacheVersionNameVid, data.Name)
		bugconfig.CacheVidName[int64(data.Id)] = data.Name
		bugconfig.CacheVersionNameVid[data.Name] = int64(data.Id)

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
