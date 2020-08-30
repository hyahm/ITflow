package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/response"
	"itflow/model"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func LevelGet(w http.ResponseWriter, r *http.Request) {

	data := &model.List_levels{
		Levels: make([]*model.Table_level, 0),
	}
	rows, err := db.Mconn.GetRows("select id,name from level")
	if err != nil {
		golog.Error(err)
		w.Write(data.ErrorE(err))
		return
	}

	for rows.Next() {
		level := &model.Table_level{}
		err = rows.Scan(&level.Id, &level.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		data.Levels = append(data.Levels, level)
	}

	w.Write(data.Marshal())
	return

}

func LevelAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Data_level)
	var err error
	errorcode.Id, err = db.Mconn.Insert("insert into level(name) value(?)", data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func LevelDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 判断bug是否在使用
	var count int
	err = db.Mconn.GetOne("select count(id) from bugs where lid=?", id64).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.IsUse())
		return
	}
	// 是否设定为了默认值
	// if cache.CacheDefault["level"] == int64(id32) {
	// 	w.Write(errorcode.Error("没有设置 level 默认值"))
	// 	return
	// }
	gsql := "delete from level where id=?"
	_, err = db.Mconn.Update(gsql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 删除缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func LevelUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := xmux.GetData(r).Data.(*model.Update_level)
	gsql := "update level set name=? where id=?"
	_, err := db.Mconn.Update(gsql, data.Name, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type levelslist struct {
	Levels []string `json:"levels"`
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
}

func (ll *levelslist) Marshal() []byte {
	send, err := json.Marshal(ll)
	if err != nil {
		golog.Error(err)
	}
	return send
}
func (ll *levelslist) Error(msg string) []byte {
	ll.Code = 1
	ll.Msg = msg
	return ll.Marshal()

}
func (ll *levelslist) ErrorE(err error) []byte {
	return ll.Error(err.Error())
}

func GetLevels(w http.ResponseWriter, r *http.Request) {

	data := &levelslist{
		Levels: make([]string, 0),
	}
	rows, err := db.Mconn.GetRows("select name from level")
	if err != nil {
		golog.Error(err)
		w.Write(data.ErrorE(err))
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Error(err)
			continue
		}
		data.Levels = append(data.Levels, name)
	}
	w.Write(data.Marshal())
	return

}
