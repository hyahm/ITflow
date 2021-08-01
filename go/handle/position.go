package handle

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func PositionGet(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &model.Jobs{
		Positions: make([]*model.Job, 0),
	}

	rows, err := db.Mconn.GetRows(`select j.id,j.name,level,hypo,IFNULL(s.name,''), IFNULL(r.name,'') from jobs as j 
	left join statusgroup as s  on j.bugsid = s.id 
	left join rolegroup as r on j.rid=r.id`)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	x := make(map[int64]string)
	for rows.Next() {
		one := &model.Job{}
		rows.Scan(&one.Id, &one.Name, &one.Level, &one.Hid, &one.StatusGroup, &one.RoleGroup)
		x[one.Id] = one.Name
		data.Positions = append(data.Positions, one)
	}
	rows.Close()
	for i := range data.Positions {
		if data.Positions[i].Hid > 0 {
			data.Positions[i].HypoName = x[data.Positions[i].Hid]
		}
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func PositionAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetInstance(r).Data.(*model.Job)
	golog.Infof("%#v", *data)
	// 如果不存在管理层名，就参数错误
	var hid int64
	var err error
	golog.Info("start add position")
	if data.HypoName != "" {
		err = db.Mconn.GetOne("select id from jobs where name=?", data.HypoName).Scan(&hid)
		if err != nil {
			if err != sql.ErrNoRows {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}

		}
	}

	errorcode.ID, err = db.Mconn.Insert(`insert into jobs(name,level,hypo,rid,bugsid) 
	value(?,?,?,(select id from rolegroup where name=?),(select id from statusgroup where name=?))`,
		data.Name, data.Level, hid, data.RoleGroup, data.StatusGroup)
	if err != nil {
		golog.Error(err)
		if err.(*mysql.MySQLError).Number == 1062 {
			w.Write(errorcode.ErrorE(db.DuplicateErr))
			return
		}
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, err := json.Marshal(errorcode)
	if err != nil {
		golog.Error(err)
	}
	w.Write(send)
	return

}

func PositionDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")

	// 如果这个职位被使用了，不允许被删除
	var count int
	err := db.Mconn.GetOne("select count(id) from user where jid=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		golog.Error("有用户使用了此职位")
		w.Write(errorcode.Error("有用户使用了此职位"))
		return
	}
	// 是否被所属使用
	err = db.Mconn.GetOne("select count(id) from jobs where hypo=?", id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		golog.Error("此管理者还管理者其他员工")
		w.Write(errorcode.Error("此管理者还管理者其他员工"))
		return
	}
	gsql := "delete from jobs where id=?"

	_, err = db.Mconn.Update(gsql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func PositionUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := xmux.GetInstance(r).Data.(*model.Job)

	var hid int64
	golog.Infof("%+v", *data)
	// 不存在这个职位的hypo，直接返回参数错误
	if data.HypoName != "" {
		err := db.Mconn.GetOne("select id from jobs where name=?", data.HypoName).Scan(&hid)
		if err != nil {
			if err == sql.ErrNoRows {
				w.Write(errorcode.Error("没有找到职位"))
				return
			}
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	_, err := db.Mconn.Update(`update jobs set name=?,level=?,hypo=?, 
	bugsid=(select id from statusgroup where name=?),rid=(select id from rolegroup where name=?) where id=?`,
		data.Name, data.Level, hid, data.StatusGroup, data.RoleGroup, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type Hypo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RespHypos struct {
	Hypos []*Hypo `json:"hypos"`
	Code  int     `json:"code"`
}

func GetHypos(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	// nickname := xmux.GetInstance(r).Get("nickname").(string)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	data := &RespHypos{
		Hypos: make([]*Hypo, 0),
	}
	// 管理员
	if uid != cache.SUPERID {
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	jid := r.FormValue("id")
	rows, err := db.Mconn.GetRows("select id,name from jobs where level=1 and id<>?", jid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		hn := &Hypo{}
		err = rows.Scan(&hn.Id, &hn.Name)
		if err != nil {
			golog.Info(err)
			continue
		}
		data.Hypos = append(data.Hypos, hn)
	}
	rows.Close()
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

type positions struct {
	Positions []string `json:"positions"`
	Code      int      `json:"code"`
}

func GetPositions(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &positions{}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	var rows *sql.Rows
	var err error
	if uid == cache.SUPERID {
		rows, err = db.Mconn.GetRows("select name from jobs")
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	} else {
		rows, err = db.Mconn.GetRows("select name from jobs where hypo=(select jid from user where id=?)", uid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Info(err)
			continue
		}
		data.Positions = append(data.Positions, name)
	}
	rows.Close()
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
