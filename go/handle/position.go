package handle

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/response"
	"itflow/model"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func PositionGet(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := &model.List_jobs{
		Positions: make([]*model.Table_jobs, 0),
	}

	rows, err := db.Mconn.GetRows("select id,name,level,hypo from jobs")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	x := make(map[int64]string)
	for rows.Next() {
		one := &model.Table_jobs{}
		rows.Scan(&one.Id, &one.Name, &one.Level, &one.Hid)
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

	data := xmux.GetData(r).Data.(*model.Data_jobs)

	// 如果不存在管理层名，就参数错误
	var hid int64
	var err error
	golog.Info("start add position")
	if data.Hyponame != "" {
		err = db.Mconn.GetOne("select id from jobs where name=?", data.Hyponame).Scan(&hid)
		if err != nil {
			if err != sql.ErrNoRows {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}

		}
	}

	errorcode.Id, err = db.Mconn.Insert("insert into jobs(name,level,hypo) value(?,?,?)", data.Name, data.Level, hid)
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
	data := xmux.GetData(r).Data.(*model.Update_jobs)

	var hid int64
	golog.Infof("%+v", *data)
	// 不存在这个职位的hypo，直接返回参数错误
	if data.Hyponame != "" {
		err := db.Mconn.GetOne("select id from jobs where name=?", data.Hyponame).Scan(&hid)
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

	_, err := db.Mconn.Update("update jobs set name=?,level=?,hypo=? where id=?", data.Name, data.Level, hid, data.Id)
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
	// nickname := xmux.GetData(r).Get("nickname").(string)
	uid := xmux.GetData(r).Get("uid").(int64)
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
			golog.Error(err)
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
	uid := xmux.GetData(r).Get("uid").(int64)
	if uid != cache.SUPERID {
		errorcode.Msg = "暂时只能管理员才能添加用户"
		w.Write(errorcode.Success())
		return
	}
	rows, err := db.Mconn.GetRows("select name from jobs")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Error(err)
			continue
		}
		data.Positions = append(data.Positions, name)
	}
	rows.Close()
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
