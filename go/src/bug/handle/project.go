package handle

import (
	"bug/asset"
	"bug/bugconfig"
	"bug/buglog"
	"encoding/json"
	"galog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type projectlist struct {
	Plist []*projectrow `json:"projectlist"`
	Code  int           `json:"statuscode"`
}

type projectrow struct {
	Id          int64  `json:"id"`
	ProjectName string `json:"projectname"`
}

func ProjectList(w http.ResponseWriter, r *http.Request) {
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
		projects := &projectlist{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("project", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		for k, v := range bugconfig.CachePidName {
			pr := &projectrow{
				Id:          k,
				ProjectName: v,
			}
			projects.Plist = append(projects.Plist, pr)
		}

		send, _ := json.Marshal(projects)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func AddProject(w http.ResponseWriter, r *http.Request) {
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
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("project", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		name := r.FormValue("name")

		getaritclesql := "insert into projectname(name) values(?)"

		errorcode.Id, err = conn.InsertWithID(getaritclesql, name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "project",
		}
		err = il.Add(
			nickname, errorcode.Id, name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 更新缓存
		bugconfig.CacheProjectPid[name] = errorcode.Id
		bugconfig.CachePidName[errorcode.Id] = name
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
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
		pr := &projectrow{}
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("project", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		bpr, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(bpr, pr)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		getaritclesql := "update projectname set name=? where id=?"

		_, err = conn.Update(getaritclesql, pr.ProjectName, pr.Id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "project",
		}
		err = il.Update(
			nickname, pr.Id, pr.ProjectName)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 更新缓存
		delete(bugconfig.CacheProjectPid, bugconfig.CachePidName[int64(pr.Id)])
		bugconfig.CachePidName[int64(pr.Id)] = pr.ProjectName
		bugconfig.CacheProjectPid[pr.ProjectName] = int64(pr.Id)

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
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
		var permssion bool
		// 管理员
		if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
			permssion = true
		} else {
			permssion, err = asset.CheckPerm("project", nickname, conn)
			if err != nil {
				galog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}
		}

		if !permssion {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		id := r.FormValue("id")
		pid, err := strconv.Atoi(id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		// 是否有bug使用
		var count int
		err = conn.GetOne("select count(id) from bugs where pid=?", id).Scan(&count)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		if count > 0 {
			w.Write(errorcode.ErrorHasBug())
			return
		}

		getaritclesql := "delete from projectname where id=?"

		_, err = conn.Insert(getaritclesql, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "project",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 更新缓存
		delete(bugconfig.CacheProjectPid, bugconfig.CachePidName[int64(pid)])
		delete(bugconfig.CachePidName, int64(pid))

		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
