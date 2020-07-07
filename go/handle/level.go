package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/model"
	network "itflow/model"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func LevelGet(w http.ResponseWriter, r *http.Request) {

	data := &network.List_levels{
		Levels: make([]*network.Table_level, 0),
	}
	for k, v := range cache.CacheLidLevel {
		one := &network.Table_level{}
		one.Id = k
		one.Name = v
		data.Levels = append(data.Levels, one)
	}

	send, _ := json.Marshal(data)
	w.Write(send)
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

	// 增加日志
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "level",
		Action:   "add",
	}

	//更新缓存
	cache.CacheLidLevel[cache.LevelId(errorcode.Id)] = data.Name
	cache.CacheLevelLid[data.Name] = cache.LevelId(errorcode.Id)

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
		w.Write(errorcode.Error("没有bug"))
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

	// 增加日志
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "level",
		Action:   "delete",
	}

	// 删除缓存
	delete(cache.CacheLevelLid, cache.CacheLidLevel[cache.LevelId(id64)])
	delete(cache.CacheLidLevel, cache.LevelId(id64))
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
	// 增加日志
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "level",
		Action:   "update",
	}

	// 删除strings key
	delete(cache.CacheLevelLid, data.OldName)
	cache.CacheLidLevel[data.Id] = data.Name
	cache.CacheLevelLid[data.Name] = data.Id

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

type levelslist struct {
	Levels []cache.Level `json:"levels"`
	Code   int           `json:"code"`
}

func GetLevels(w http.ResponseWriter, r *http.Request) {

	data := &levelslist{
		Levels: make([]cache.Level, 0),
	}
	for _, v := range cache.CacheLidLevel {
		data.Levels = append(data.Levels, v)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
