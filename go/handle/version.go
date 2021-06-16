package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/perm"
	"itflow/internal/response"
	"itflow/internal/version"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func AddVersion(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Insert {
		w.Write(errorcode.Error("no perm"))
		return
	}
	version_add := xmux.GetInstance(r).Data.(*version.RespVersion)

	uid := xmux.GetInstance(r).Get("uid").(int64)
	add_version_sql := "insert into version(pid,name,urlone,urltwo,createtime,createuid) values((select id from project where name=?),?,?,?,?,?)"
	var err error
	errorcode.UpdateTime = time.Now().Unix()
	errorcode.Id, err = db.Mconn.Insert(add_version_sql, version_add.Project,
		version_add.Name, version_add.Url, version_add.BakUrl, errorcode.UpdateTime, uid)
	if err != nil {
		golog.Error(err)
		if err.(*mysql.MySQLError).Number == 1062 {
			w.Write(errorcode.ErrorE(db.DuplicateErr))
			return
		}
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func VersionList(w http.ResponseWriter, r *http.Request) {

	al := &version.VersionList{
		VersionList: make([]*version.RespVersion, 0),
	}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Select {
		w.Write(al.Error("no perm"))
		return
	}
	get_version_sql := `select v.id, ifnull(p.name,''), v.name, v.urlone, v.urltwo, v.createtime 
	from version as v  
	left join 
	project as p  
	on v.pid =p.id 
	order by v.id desc;`

	rows, err := db.Mconn.GetRows(get_version_sql)
	if err != nil {
		golog.Error(err)
		w.Write(al.ErrorE(err))
		return
	}

	for rows.Next() {
		rl := &version.RespVersion{}
		err = rows.Scan(&rl.Id, &rl.Project, &rl.Name, &rl.Url, &rl.BakUrl, &rl.Date)
		if err != nil {
			golog.Info(err)
			continue
		}
		al.VersionList = append(al.VersionList, rl)
	}
	rows.Close()
	send, _ := json.Marshal(al)
	w.Write(send)
	return

}

func VersionRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Delete {
		w.Write(errorcode.Error("no perm"))
		return
	}
	id := r.FormValue("id")
	var bugcount int

	err := db.Mconn.GetOne("select count(id) from bugs where vid=?", id).Scan(&bugcount)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if bugcount > 0 {
		golog.Errorf("vid:%s has bugs", id)
		w.Write(errorcode.IsUse())
		return
	}
	deletevl := "delete from version where id=?"
	errorcode.Id, err = db.Mconn.Update(deletevl, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// vid, err := strconv.Atoi(id)
	// if err != nil {
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// 增加日志

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

	errorcode := &response.Response{}
	perm := xmux.GetInstance(r).Get("perm").(perm.OptionPerm)
	if !perm.Update {
		w.Write(errorcode.Error("no perm"))
		return
	}
	data := xmux.GetInstance(r).Data.(*version.RespVersion)

	uid := xmux.GetInstance(r).Get("uid").(int64)
	versionsql := "update version set pid=(select id from project where name=?),name=?,urlone=?,urltwo=?,createuid=? where id=?"
	_, err := db.Mconn.Update(versionsql, data.Project, data.Name, data.Url, data.BakUrl, uid, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}
