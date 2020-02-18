package handle

import (
	"encoding/json"
	"io/ioutil"
	"itflow/bug/asset"
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type projectlist struct {
	Plist []*projectrow `json:"projectlist"`
	Code  int           `json:"code"`
}

type projectrow struct {
	Id          int64  `json:"id"`
	ProjectName string `json:"projectname"`
}

func ProjectList(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	projects := &projectlist{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("project", nickname)
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

func AddProject(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
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
		permssion, err = asset.CheckPerm("project", nickname)
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
	name := r.FormValue("name")

	getaritclesql := "insert into projectname(name) values(?)"

	errorcode.Id, err = db.Mconn.Insert(getaritclesql, name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "project",
	}
	err = il.Add(
		nickname, errorcode.Id, name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 更新缓存
	bugconfig.CacheProjectPid[name] = errorcode.Id
	bugconfig.CachePidName[errorcode.Id] = name
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	pr := &projectrow{}
	var permssion bool
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("project", nickname)
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
	bpr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(bpr, pr)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	getaritclesql := "update projectname set name=? where id=?"

	_, err = db.Mconn.Update(getaritclesql, pr.ProjectName, pr.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "project",
	}
	err = il.Update(
		nickname, pr.Id, pr.ProjectName)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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

func DeleteProject(w http.ResponseWriter, r *http.Request) {

	nickname, err := logtokenmysql(r)
	errorcode := &errorstruct{}
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
		permssion, err = asset.CheckPerm("project", nickname)
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
	id := r.FormValue("id")
	pid, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 是否有bug使用
	var count int
	row, err := db.Mconn.GetOne("select count(id) from bugs where pid=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("没有bug"))
		return
	}

	getaritclesql := "delete from projectname where id=?"

	_, err = db.Mconn.Insert(getaritclesql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	il := buglog.AddLog{
		Ip:       strings.Split(r.RemoteAddr, ":")[0],
		Classify: "project",
	}
	err = il.Del(
		nickname, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 更新缓存
	delete(bugconfig.CacheProjectPid, bugconfig.CachePidName[int64(pid)])
	delete(bugconfig.CachePidName, int64(pid))

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
