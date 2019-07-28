package handle

import (
	"itflow/bug/bugconfig"
	"itflow/bug/buglog"
	"itflow/bug/model"
	"encoding/json"
	"fmt"
	"github.com/hyahm/golog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func TypeList(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		tl := &model.List_types{}
		conn, _, err := logtokenmysql(r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		rows, err := conn.GetRows("select id,name,type,opts,tid from types")
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		fmt.Printf("%+v \n", bugconfig.CacheTidName)
		for rows.Next() {
			var opts string
			tr := &model.Data_types{}
			var tid int64
			rows.Scan(&tr.Id, &tr.Name, &tr.Types, &opts, &tid)
			tr.Listtype = bugconfig.CacheTidName[tid]

			if tr.Types == 2 {

				optrows, err := conn.GetRows(fmt.Sprintf("select id,name,info,tid,df,need  from options where id in (%s)", opts))
				if err != nil {
					golog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}

				for optrows.Next() {
					one := &model.Options{}
					var ts int64
					optrows.Scan(&one.Id, &one.Name, &one.Info, &ts, &one.Default, &one.Need)

					if tstring, ok := bugconfig.CacheTidName[ts]; ok {
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
}

func TypeUpdate(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		data := &model.Data_types{}
		send := &model.Send_types{}
		respbyte, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		fmt.Println(string(respbyte))
		err = json.Unmarshal(respbyte, data)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}

		var t int
		err = conn.GetOne("select type from types where id=?", data.Id).Scan(&t)
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

			if tid, ok := bugconfig.CacheNameTid[data.Listtype]; ok {
				if tid == data.Id {
					golog.Error("this type is updating")
					w.Write(errorcode.ErrorType())
					return
				}
				_, err = conn.Update("update types set name=?,type=?,tid=? where id=?", data.Name, data.Types, tid, data.Id)
				if err != nil {
					golog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
			} else {
				golog.Error("params error, type:%d", data.Types)
				errorcode.ErrorParams()
				return
			}
		case 2:
			// object
			optids := make([]string, 0)
			l := 0
			for _, v := range data.Opts {
				if v.Id >= 0 {
					if tid, ok := bugconfig.CacheNameTid[v.Type]; ok {
						if tid == data.Id {
							golog.Error("this type is updating")
							w.Write(errorcode.ErrorType())
							return
						}
						_, err = conn.Update("update options set name=?,info=?,tid=?,df=?,need=? where id=?",
							v.Name, v.Info, tid, v.Default, v.Need, v.Id)
						if err != nil {
							golog.Error(err.Error())
							w.Write(errorcode.ErrorConnentMysql())
							return
						}
						l++
						optids = append(optids, strconv.FormatInt(v.Id, 10))
						send.Opts = append(send.Opts, v)
					}
				} else if v.Id < 0 {
					if tid, ok := bugconfig.CacheNameTid[v.Type]; ok {
						v.Id, err = conn.InsertWithID("insert into options(name,info,tid,df,need) values(?,?,?,?,?)",
							v.Name, v.Info, tid, v.Default, v.Need)
						if err != nil {
							golog.Error(err.Error())
							w.Write(errorcode.ErrorConnentMysql())
							return
						}

						l++
						optids = append(optids, strconv.FormatInt(v.Id, 10))
						send.Opts = append(send.Opts, v)
					}

				} else {
					golog.Error("params error, type:%d", data.Types)
					w.Write(errorcode.ErrorParams())
					return
				}

			}
			if l == 0 {
				golog.Error("need opts")
				w.Write(errorcode.ErrorParams())
				return
			}
			opts := strings.Join(optids, ",")
			_, err = conn.Update("update types set name=?,type=?,opts=? where id=?", data.Name, data.Types, opts, data.Id)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}

		default:
			golog.Error("params error, type:%d", data.Types)
			w.Write(errorcode.ErrorParams())
			return

		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "type",
		}
		err = il.Update(
			nickname, data.Id, data.Name)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		delete(bugconfig.CacheNameTid, bugconfig.CacheTidName[data.Id])
		bugconfig.CacheTidName[data.Id] = data.Name
		bugconfig.CacheNameTid[data.Name] = data.Id
		s, _ := json.Marshal(send)
		fmt.Println(string(s))
		w.Write(s)
		return
	}
}

func TypeAdd(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		errorcode := &errorstruct{}
		conn, _, err := logtokenmysql(r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		data := &model.Data_types{}
		send := &model.Send_types{}
		bytedata, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		fmt.Println(string(bytedata))
		err = json.Unmarshal(bytedata, data)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		switch data.Types {
		case 1:
			// list 类型
			if tid, ok := bugconfig.CacheNameTid[data.Listtype]; ok {
				send.Id, err = conn.Insert("insert into types(name,type,tid) value(?,?,?)", data.Name, data.Types, tid)
				if err != nil {
					golog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
			} else {
				golog.Error("params error, type:%d", data.Types)
				errorcode.ErrorParams()
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

				if tid, ok := bugconfig.CacheNameTid[v.Type]; ok {
					v.Id, err = conn.InsertWithID("insert into options(name,info,tid,df,need) values(?,?,?,?,?)",
						v.Name, v.Info, tid, v.Default, v.Need)
					if err != nil {
						golog.Error(err.Error())
						w.Write(errorcode.ErrorConnentMysql())
						return
					}
					l++
					optids = append(optids, strconv.FormatInt(v.Id, 10))
					send.Opts = append(send.Opts, v)
				} else {
					golog.Error("params error, type:%d", data.Types)
					w.Write(errorcode.ErrorParams())
					return
				}

			}
			if l == 0 {
				golog.Error("need opts")
				w.Write(errorcode.ErrorParams())
				return
			}
			opts := strings.Join(optids, ",")
			send.Id, err = conn.InsertWithID("insert into types(name,type,opts) values(?,?,?)", data.Name, data.Types, opts)
			if err != nil {
				golog.Error(err.Error())
				w.Write(errorcode.ErrorConnentMysql())
				return
			}

		default:
			golog.Error("params error, type:%d", data.Types)
			w.Write(errorcode.ErrorParams())
			return

		}
		fmt.Println("-------------------")
		//fmt.Printf("%+v \n", data)
		bugconfig.CacheTidName[send.Id] = data.Name
		bugconfig.CacheNameTid[data.Name] = send.Id
		//return
		//errorcode.Id, err = conn.Insert("insert into types(name) values(?)", name)
		//if err != nil {
		//	golog.Error(err.Error())
		//	w.Write(errorcode.ErrorConnentMysql())
		//	return
		//}
		// 增加日志
		//il := buglog.AddLog{
		//	Conn:     conn,
		//	Ip:       strings.Split(r.RemoteAddr, ":")[0],
		//	Classify: "type",
		//}
		//err = il.Add(
		//	nickname, errorcode.Id, name)
		//if err != nil {
		//	golog.Error(err.Error())
		//	w.Write(errorcode.ErrorConnentMysql())
		//	return
		//}
		s, _ := json.Marshal(send)
		w.Write(s)
		return
	}
}

func TypeDel(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		id := r.FormValue("id")
		id32, err := strconv.Atoi(id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		conn, nickname, err := logtokenmysql(r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		var t int
		err = conn.GetOne("select type from types where id=?", id).Scan(&t)
		if t == 0 {
			golog.Error("can not delete base type")
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		_, err = conn.Update("delete from types where id=?", id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 增加日志
		il := buglog.AddLog{
			Conn:     conn,
			Ip:       strings.Split(r.RemoteAddr, ":")[0],
			Classify: "type",
		}
		err = il.Del(
			nickname, id)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		delete(bugconfig.CacheNameTid, bugconfig.CacheTidName[int64(id32)])
		delete(bugconfig.CacheTidName, int64(id32))
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
}

func GetType(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		errorcode := &errorstruct{}
		conn, _, err := logtokenmysql(r)
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		defer conn.Db.Close()
		types := &model.Send_Types{}
		rows, err := conn.GetRows("select name from types")
		if err != nil {
			golog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
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
}
