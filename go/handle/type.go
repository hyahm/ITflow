package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"itflow/model"
	network "itflow/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func TypeList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	tl := &network.List_types{}

	rows, err := db.Mconn.GetRows("select id,name,type,opts,tid from types")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		var opts string
		tr := &network.Data_types{}
		var tid int64
		rows.Scan(&tr.Id, &tr.Name, &tr.Types, &opts, &tid)
		tr.Listtype = cache.CacheTidName[tid]

		if tr.Types == 2 {

			optrows, err := db.Mconn.GetRows(fmt.Sprintf("select id,name,info,tid,df,need  from options where id in (%s)", opts))
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}

			for optrows.Next() {
				one := &network.Options{}
				var ts int64
				optrows.Scan(&one.Id, &one.Name, &one.Info, &ts, &one.Default, &one.Need)

				if tstring, ok := cache.CacheTidName[ts]; ok {
					one.Type = tstring

					tr.Opts = append(tr.Opts, one)
				}
			}
		}
		tl.List = append(tl.List, tr)
	}
	send, _ := json.Marshal(tl)
	w.Write(send)
	return

}

func TypeUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Data_types)

	send := &network.Send_types{}

	var t int
	err := db.Mconn.GetOne("select type from types where id=?", data.Id).Scan(&t)
	if t == 0 {
		golog.Error("can not delete base type")
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	// 垃圾没清理
	send.Id = data.Id
	switch data.Types {
	case 1:
		// list 类型

		if tid, ok := cache.CacheNameTid[data.Listtype]; ok {
			if tid == data.Id {
				golog.Error("this type is updating")
				w.Write(errorcode.Error("this type is updating"))
				return
			}
			_, err = db.Mconn.Update("update types set name=?,type=?,tid=? where id=?", data.Name, data.Types, tid, data.Id)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}
		} else {
			golog.Errorf("params error, type:%d", data.Types)
			errorcode.Errorf("params error, type:%d", data.Types)
			return
		}
	case 2:
		// object
		optids := make([]string, 0)
		l := 0
		for _, v := range data.Opts {
			if v.Id >= 0 {
				if tid, ok := cache.CacheNameTid[v.Type]; ok {
					if tid == data.Id {
						golog.Error("this type is updating")
						w.Write(errorcode.Error("this type is updating"))
						return
					}
					_, err = db.Mconn.Update("update options set name=?,info=?,tid=?,df=?,need=? where id=?",
						v.Name, v.Info, tid, v.Default, v.Need, v.Id)
					if err != nil {
						golog.Error(err)
						w.Write(errorcode.ErrorE(err))
						return
					}
					l++
					optids = append(optids, strconv.FormatInt(v.Id, 10))
					send.Opts = append(send.Opts, v)
				}
			} else if v.Id < 0 {
				if tid, ok := cache.CacheNameTid[v.Type]; ok {
					v.Id, err = db.Mconn.Insert("insert into options(name,info,tid,df,need) values(?,?,?,?,?)",
						v.Name, v.Info, tid, v.Default, v.Need)
					if err != nil {
						golog.Error(err)
						w.Write(errorcode.ErrorE(err))
						return
					}

					l++
					optids = append(optids, strconv.FormatInt(v.Id, 10))
					send.Opts = append(send.Opts, v)
				}

			} else {
				golog.Errorf("params error, type:%d", data.Types)
				w.Write(errorcode.Errorf("params error, type:%d", data.Types))
				return
			}

		}
		if l == 0 {
			golog.Error("need opts")
			w.Write(errorcode.Error("need opts"))
			return
		}
		opts := strings.Join(optids, ",")
		_, err = db.Mconn.Update("update types set name=?,type=?,opts=? where id=?", data.Name, data.Types, opts, data.Id)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

	default:
		golog.Errorf("params error, type:%d", data.Types)
		w.Write(errorcode.Errorf("params error, type:%d", data.Types))
		return

	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "type",
		Action:   "update",
	}

	delete(cache.CacheNameTid, cache.CacheTidName[data.Id])
	cache.CacheTidName[data.Id] = data.Name
	cache.CacheNameTid[data.Name] = data.Id
	s, _ := json.Marshal(send)
	w.Write(s)
	return

}

func TypeAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Data_types)

	send := &model.Send_types{}
	var err error
	switch data.Types {
	case 1:
		// list 类型
		if tid, ok := cache.CacheNameTid[data.Listtype]; ok {
			send.Id, err = db.Mconn.Insert("insert into types(name,type,tid) value(?,?,?)", data.Name, data.Types, tid)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}
		} else {
			golog.Errorf("params error, type:%d", data.Types)
			errorcode.Errorf("params error, type:%d", data.Types)
			return
		}
	case 2:
		// object
		optids := make([]string, 0)
		l := 0
		for _, v := range data.Opts {
			if v.Id >= 0 {
				continue
			}

			if tid, ok := cache.CacheNameTid[v.Type]; ok {
				v.Id, err = db.Mconn.Insert("insert into options(name,info,tid,df,need) values(?,?,?,?,?)",
					v.Name, v.Info, tid, v.Default, v.Need)
				if err != nil {
					golog.Error(err)
					w.Write(errorcode.ErrorE(err))
					return
				}
				l++
				optids = append(optids, strconv.FormatInt(v.Id, 10))
				send.Opts = append(send.Opts, v)
			} else {
				golog.Errorf("params error, type:%d", data.Types)
				w.Write(errorcode.Errorf("params error, type:%d", data.Types))
				return
			}

		}
		if l == 0 {
			golog.Error("need opts")
			w.Write(errorcode.Error("need opts"))
			return
		}
		opts := strings.Join(optids, ",")
		send.Id, err = db.Mconn.Insert("insert into types(name,type,opts) values(?,?,?)", data.Name, data.Types, opts)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}

	default:
		golog.Errorf("params error, type:%d", data.Types)
		w.Write(errorcode.Errorf("params error, type:%d", data.Types))
		return

	}

	cache.CacheTidName[send.Id] = data.Name
	cache.CacheNameTid[data.Name] = send.Id
	nickname := xmux.GetData(r).Get("nickname").(string)
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "type",
		Action:   "delete",
		Msg:      fmt.Sprintf("delete type: %s", data.Types),
	}
	s, _ := json.Marshal(send)
	w.Write(s)
	return

}

func TypeDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	id := r.FormValue("id")
	id32, err := strconv.Atoi(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	nickname := xmux.GetData(r).Get("nickname").(string)
	var t int
	err = db.Mconn.GetOne("select type from types where id=?", id).Scan(&t)
	if t == 0 {
		golog.Error("can not delete base type")
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	_, err = db.Mconn.Update("delete from types where id=?", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 增加日志
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "type",
		Action:   "delete",
		Msg:      fmt.Sprintf("delete type id: %s", id),
	}

	delete(cache.CacheNameTid, cache.CacheTidName[int64(id32)])
	delete(cache.CacheTidName, int64(id32))
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func GetType(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	types := &network.Send_Types{}
	rows, err := db.Mconn.GetRows("select name from types")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var _type string
		rows.Scan(&_type)
		types.Types = append(types.Types, _type)
	}
	send, _ := json.Marshal(types)
	w.Write(send)
	return

}
