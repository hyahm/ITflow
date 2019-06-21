package handle

import (
	"bug/asset"
	"bug/bugconfig"
	"bug/model"
	"encoding/json"
	"gadb"
	"gaencrypt"
	"galog"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type onefd struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	ModDate  int64  `json:"date"`
	IsFile   bool   `json:"isfile"`
	IsOwner  bool   `json:"isowner"`
	HasWrite bool   `json:"haswrite"`
	Ru       bool   `json:"readuser"`
	Rname    string `json:"readname"`
	Wu       bool   `json:"writeuser"`
	Wname    string `json:"writename"`
}

type getpath struct {
	Path string `json:"path"` // 需要创建的目录
	Now  string `json:"now"`  // 前端当前所在路径
}

func ShareList(w http.ResponseWriter, r *http.Request) {
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
		fd := &model.List_sharelist{}
		uid := bugconfig.CacheNickNameUid[nickname]
		path := r.FormValue("path")
		grows, err := conn.GetRows("select id,ids from usergroup")
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 哪些gid 包含了此用户权限
		fd.RealName = bugconfig.CacheUidRealName[uid]
		rg := make([]int64, 0)
		for grows.Next() {
			var id int64
			var ids string
			grows.Scan(&id, &ids)
			users := strings.Split(ids, ",")
			for _, v := range users {
				if v == strconv.FormatInt(uid, 10) {
					rg = append(rg, id)
					break
				}
			}
		}

		// 如果是自己
		sql := "select id,isfile,size,updatetime,name,ownerid,readuser,rid,wid,writeuser from sharefile where filepath=? "
		rows, err := conn.GetRows(sql, filepath.Clean(path))
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		for rows.Next() {
			od := &model.Data_sharefile{}
			var oid int64
			var rid int64
			var wid int64
			var hasperm bool
			rows.Scan(&od.Id, &od.IsFile, &od.Size, &od.UpdateTime, &od.Name, &oid, &od.ReadUser, &rid, &wid, &od.WriteUser)
			//1, 自己的，直接可以添加正确的
			od.IsOwner = uid == oid
			if od.IsOwner {
				hasperm = true
			}
			// 2,授权用户，添加
			if od.ReadUser {
				if hasperm || rid == uid {
					hasperm = true
					od.ReadName = bugconfig.CacheUidRealName[rid]
				}

			} else {
				// 3,授权的组，添加
				// 找得到就退出循环
				for _, v := range rg {
					if hasperm || v == rid {
						od.ReadName = bugconfig.CacheGidGroup[rid]
						hasperm = true
						break
					}
				}
			}
			if od.WriteUser {
				if hasperm || wid == uid {
					hasperm = true
					od.WriteName = bugconfig.CacheUidRealName[wid]
				} else {
					continue
				}
			} else {
				for _, v := range rg {
					if hasperm || v == wid {
						od.WriteName = bugconfig.CacheGidGroup[wid]
						hasperm = true
						break
					}
				}
			}

			if !hasperm {
				continue
			}

			fd.FDList = append(fd.FDList, od)
		}

		////用户
		//usersql := "select id,isfile,size,updatetime,name,ownerid,readuser,rid,wid,writeuser from sharefile where filepath=? and ownerid<>? and readuser=true"
		//userrows, err := conn.GetRows(usersql, path, uid)
		//if err != nil {
		//	galog.Error(err.Error())
		//	w.Write(errorcode.ErrorConnentMysql())
		//	return
		//}
		//for userrows.Next() {
		//	od := &onefd{}
		//	var oid int
		//	var rid int64
		//	var ru bool
		//	var wid int64
		//	var wu bool
		//	userrows.Scan(&od.Id, &od.IsFile, &od.Size, &od.ModDate, &od.Name, &oid, &ru, &rid, &wid, &wu)
		//	od.Ru = ru
		//	if od.Ru {
		//		od.Wname = bugconfig.CacheUidRealName[rid]
		//	} else {
		//		od.Wname = bugconfig.CacheGidGroup[rid]
		//	}
		//
		//	od.Wu = wu
		//	if od.Wu {
		//		od.Wname = bugconfig.CacheUidRealName[wid]
		//	} else {
		//		od.Wname = bugconfig.CacheGidGroup[wid]
		//	}
		//	od.IsOwner = false
		//	fd.FDList = append(fd.FDList, od)
		//}
		//
		//groupsql := "select id,isfile,size,updatetime,name,ownerid,readuser,rid,wid,writeuser from sharefile where filepath=? and ownerid<>?  and readuser=false "
		////insql := ""
		//wsql := ""
		//fmt.Println("rg:", rg)
		//if len(rg) > 2 {
		//	mark := ""
		//	for i, _ := range rg {
		//		if i == 0 {
		//			mark = "?"
		//		} else {
		//			mark = mark + ",?"
		//		}
		//	}
		//	wsql = fmt.Sprintf(" and rid in (%s) ", mark)
		//}
		//fmt.Println("gw:", groupsql+wsql)
		//grouprows, err := conn.GetRows(groupsql+wsql,
		//	rg...)
		//if err != nil {
		//	galog.Error(err.Error())
		//	w.Write(errorcode.ErrorConnentMysql())
		//	return
		//}
		//
		//for grouprows.Next() {
		//	od := &onefd{}
		//	var oid int
		//	var rid int64
		//	var ru bool
		//	var wid int64
		//	var wu bool
		//	grouprows.Scan(&od.Id, &od.IsFile, &od.Size, &od.ModDate, &od.Name, &oid, &ru, &rid, &wid, &wu)
		//	od.Ru = ru
		//	od.Rname = bugconfig.CacheGidGroup[rid]
		//	od.Wu = wu
		//	if od.Wu {
		//		od.Wname = bugconfig.CacheUidRealName[wid]
		//	} else {
		//		od.Wname = bugconfig.CacheGidGroup[wid]
		//	}
		//	od.IsOwner = false
		//	fd.FDList = append(fd.FDList, od)
		//}
		//fd.FDList = append(fd.FDList, &onefd{
		//	Name:   "..",
		//	IsFile: false,
		//})
		send, _ := json.Marshal(fd)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func ShareUpload(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(1 << 30)
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

		dir := r.FormValue("dir")

		file, header, err := r.FormFile("share")
		defer file.Close()
		// 是否为设定的根目录
		l := len(bugconfig.ShareDir)
		basedir := filepath.Join(bugconfig.ShareDir, dir)
		if len(basedir) < l && basedir[:l] != bugconfig.ShareDir {
			w.Write(errorcode.ErrorNoPermission())
			return
		}

		//验证是否有权限上传
		var (
			testid    int64
			ownerid   int64
			writeuser bool
			wid       int64
		)
		err = conn.GetOne("select id,ownerid,writeuser,wid from sharefile where isfile=false and filepath=? and name=?",
			filepath.Dir(dir), filepath.Base(dir)).Scan(&testid, &ownerid, &writeuser, &wid)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		var rwid int64
		var isuser bool
		var hasperm bool
		if ownerid == bugconfig.CacheNickNameUid[nickname] {
			rwid = ownerid
			isuser = true
			hasperm = true
		} else {
			if writeuser {
				if wid == bugconfig.CacheNickNameUid[nickname] {
					rwid = ownerid
					isuser = true
					hasperm = true
				}
			} else {
				var ids string
				err = conn.GetOne("select ids from groups where id=?", wid).Scan(&ids)
				if err != nil {
					galog.Error(err.Error())
					w.Write(errorcode.ErrorConnentMysql())
					return
				}
				for _, v := range strings.Split(ids, ",") {
					if v == strconv.FormatInt(bugconfig.CacheNickNameUid[nickname], 10) {
						rwid = wid
						isuser = false
						hasperm = true
						break
					}
				}
			}
		}
		if !hasperm {
			galog.Error("has no permssion")
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		updatetime := time.Now().Unix()
		ssql := "insert into sharefile(filepath,name,isfile,size,readuser,rid,ownerid,writeuser,wid,updatetime) values(?,?,?,?,?,?,?,?,?,?)"
		errorcode.Id, err = conn.Insert(ssql, filepath.Clean(dir), header.Filename, true, header.Size, isuser,
			rwid, bugconfig.CacheRealNameUid[nickname], isuser,
			rwid,
			updatetime,
		)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		// 生成文件
		newfile := filepath.Join(basedir, header.Filename)

		f, err := os.OpenFile(newfile, os.O_CREATE|os.O_RDWR, 0744)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		defer f.Close()
		// 写入内容到文件
		if _, err := io.Copy(f, file); err != nil {
			w.Write(errorcode.ErrorNoPermission())
		}

		errorcode.UpdateTime = updatetime
		errorcode.Size = header.Size
		errorcode.Filename = header.Filename
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

type rename struct {
	Path    string `json:"path"`
	Oldname string `json:"oldname"`
	Newname string `json:"newname"`
}

func ShareRename(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(1 << 30)
		conn, _, err := logtokenmysql(r)
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
		ps := &model.Data_sharefile{}

		pb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(pb, ps)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var (
			rid int64
			wid int64
		)
		if ps.ReadUser {
			rid = bugconfig.CacheRealNameUid[ps.ReadName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if v == ps.ReadName {
					rid = k
					break
				}
			}
		}
		if ps.WriteUser {
			wid = bugconfig.CacheRealNameUid[ps.WriteName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if ps.WriteName == v {
					wid = k
					break
				}
			}
		}
		// 检查权限
		l := len(bugconfig.ShareDir)
		basedir := filepath.Join(bugconfig.ShareDir, ps.FilePath)
		if len(basedir) < l && basedir[:l] != bugconfig.ShareDir {
			galog.Error("has not permission")
			w.Write(errorcode.ErrorNoPermission())
			return
		}

		updatetime := time.Now().Unix()
		_, err = conn.Update("update sharefile set readuser=?,rid=?,isfile=?,writeuser=?,wid=?,updatetime=?,name=? where id=? and filepath=?",
			ps.ReadUser, rid, ps.IsFile, ps.WriteUser, wid, updatetime, ps.Name, ps.Id, ps.FilePath)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		newfile := filepath.Join(basedir, ps.Name)
		oldfile := filepath.Join(basedir, ps.OldName)
		err = os.Rename(oldfile, newfile)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		errorcode.UpdateTime = updatetime
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func ShareDownload(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		r.ParseMultipartForm(1 << 30)
		conn, _, err := logtokenmysql(r)
		if err != nil {
			galog.Error(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		defer conn.Db.Close()

		// 检查权限

		//l := len(bugconfig.ShareDir)
		//basedir := filepath.Join(bugconfig.ShareDir, ds.Path, ds.Name)
		//if len(basedir) < l && basedir[:l] != bugconfig.ShareDir {
		//
		//	w.Write(errorcode.ErrorNoPermission())
		//	return
		//}
		//
		//f, err := os.Open(basedir)
		//if err != nil {
		//	galog.Error(err.Error())
		//	w.Write(errorcode.ErrorNoPermission())
		//	return
		//}
		//defer f.Close()
		//
		//errorcode.Data, err = ioutil.ReadAll(f)
		//w.Header().Set("Content-Disposition", "attachment;filename="+ds.Name)
		//w.Header().Set("Content-Type", "application/octet-stream")
		//if err != nil {
		//	galog.Error(err.Error())
		//	w.Write(errorcode.ErrorGetData())
		//	return
		//}
		//send, _ := json.Marshal(errorcode)
		//w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func ShareMkdir(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(1 << 30)
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

		ps := &model.Data_sharefile{}
		pb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorGetData())
			return
		}
		err = json.Unmarshal(pb, ps)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorParams())
			return
		}
		var (
			rid int64
			wid int64
		)
		if ps.ReadUser {
			rid = bugconfig.CacheRealNameUid[ps.ReadName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if v == ps.ReadName {
					rid = k
					break
				}
			}
		}
		if ps.WriteUser {
			wid = bugconfig.CacheRealNameUid[ps.WriteName]
		} else {
			for k, v := range bugconfig.CacheGidGroup {
				if ps.WriteName == v {
					wid = k
					break
				}
			}
		}
		// 验证目录
		l := len(bugconfig.ShareDir)
		basedir := filepath.Join(bugconfig.ShareDir, ps.FilePath)
		// 是否为设定的根目录
		if len(basedir) < l && basedir[:l] != bugconfig.ShareDir {
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		err = os.MkdirAll(filepath.Join(basedir, ps.Name), 0755)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorNoPermission())
			return
		}
		updatetime := time.Now().Unix()
		ssql := "insert into sharefile(filepath,name,isfile,size,readuser,rid,writeuser,wid,ownerid,updatetime) values(?,?,?,?,?,?,?,?,?,?)"
		errorcode.Id, err = conn.Insert(ssql, filepath.Clean(ps.FilePath), ps.Name, ps.IsFile, ps.Size, ps.ReadUser, rid, ps.WriteUser, wid, bugconfig.CacheRealNameUid[nickname], updatetime)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}

		errorcode.UpdateTime = updatetime
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func ShareRemove(w http.ResponseWriter, r *http.Request) {
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

		id := r.FormValue("id")
		// 检查文件夹权限
		var fp string
		var name string
		err = conn.GetOne("select filepath,name from sharefile where id=? and ownerid=?", id, bugconfig.CacheNickNameUid[nickname]).Scan(&fp, &name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		//还要递归删除子目录和文件
		//basedir := filepath.Join(bugconfig.ShareDir, fp, name)
		err = getrow(conn, fp, name)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		_, err = conn.Update("delete from sharefile where id=? ", id)
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		err = os.RemoveAll(filepath.Join(bugconfig.ShareDir, fp, name))
		if err != nil {
			galog.Error(err.Error())
			w.Write(errorcode.ErrorConnentMysql())
			return
		}
		send, _ := json.Marshal(errorcode)
		w.Write(send)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func getrow(conn *gadb.Db, fp string, name string) error {
	dd, err := conn.GetRows("select id,filepath,name,isfile from sharefile where filepath=? ", filepath.Join(fp, name))
	if err != nil {
		return err
	}
	for dd.Next() {
		var _id int64
		var _filepath string
		var _name string
		var _isfile bool
		err = dd.Scan(&_id, &_filepath, &_name, &_isfile)
		if err != nil {
			return err
		}
		// 删除数据库
		_, err = conn.Update("delete from sharefile where id=? ", _id)
		if err != nil {
			return err
		}
		// 删除文件或文件夹

		if !_isfile {
			err = getrow(conn, _filepath, _name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ShareShow(w http.ResponseWriter, r *http.Request) {
	headers(w, r)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		//m := mux.Vars(r)
		//name := m["name"]
		id := r.FormValue("id")
		token := r.FormValue("token")
		conn, err := gadb.NewSqlConfig().ConnDB()
		if err != nil {
			galog.Error(err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		// 只是验证token
		destoken, err := gaencrypt.RsaDecrypt(token, bugconfig.PrivateKey, true)
		if err != nil {
			galog.Error(err.Error())
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		nickname, err := asset.Getvalue(string(destoken))
		if err != nil {
			galog.Error(err.Error())
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		uid := bugconfig.CacheNickNameUid[nickname]
		var haspermision bool
		getsql := "select ownerid,filepath,name,readuser,rid from sharefile where id=? "
		var fp string
		var name string
		var readuser bool
		var rid int64
		var ownerid int64
		err = conn.GetOne(getsql, id).Scan(&ownerid, &fp, &name, &readuser, &rid)
		if err != nil {
			galog.Error(err.Error())
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if ownerid == uid {
			haspermision = true
		}
		if readuser && rid == uid {
			haspermision = true
		} else {
			var ids string
			err = conn.GetOne("select ids from groups where id=?", rid).Scan(&ids)
			if err != nil {
				galog.Error(err.Error())
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			for _, v := range strings.Split(ids, ",") {
				if v == strconv.FormatInt(uid, 10) {
					haspermision = true
				}
			}
		}
		if !haspermision {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		f, err := ioutil.ReadFile(filepath.Join(bugconfig.ShareDir, fp, name))
		if err != nil {
			galog.Error(err.Error())
			return
		}
		w.Header().Set("Content-Disposition", "attachment;filename="+name)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(f)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
