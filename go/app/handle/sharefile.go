package handle

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"itflow/app/bugconfig"
	"itflow/db"
	"itflow/gaencrypt"
	network "itflow/model"
	"itflow/network/response"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type getpath struct {
	Path string `json:"path"` // 需要创建的目录
	Now  string `json:"now"`  // 前端当前所在路径
}

func ShareList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	fd := &network.List_sharelist{}
	uid := bugconfig.CacheNickNameUid[nickname]
	path := r.FormValue("path")
	grows, err := db.Mconn.GetRows("select id,ids from usergroup")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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
	rows, err := db.Mconn.GetRows(sql, filepath.Clean(path))
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		od := &network.Data_sharefile{}
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
	//	golog.Error(err)
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
	//	golog.Error(err)
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

func ShareUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1 << 30)
	nickname := xmux.GetData(r).Get("nickname").(string)
	errorcode := &response.Response{}

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
	var rwid int64
	var isuser bool
	var hasperm bool
	// 判断是否存在这个目录
	if filepath.Base(dir) == "." {
		//根目录
		hasperm = true
	} else {
		row, err := db.Mconn.GetOne("select id,ownerid,writeuser,wid from sharefile where isfile=false and filepath=? and name=?",
			filepath.Dir(dir), filepath.Base(dir))
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		err = row.Scan(&testid, &ownerid, &writeuser, &wid)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		//判断是否有权限

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
				row, err := db.Mconn.GetOne("select ids from usergroup where id=?", wid)
				if err != nil {
					golog.Error(err)
					w.Write(errorcode.ErrorE(err))
					return
				}
				err = row.Scan(&ids)
				if err != nil {
					golog.Error(err)
					w.Write(errorcode.ErrorE(err))
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
	}

	if !hasperm {
		golog.Error("has no permssion")
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	// 插入数据
	updatetime := time.Now().Unix()
	ssql := "insert into sharefile(filepath,name,isfile,size,readuser,rid,ownerid,writeuser,wid,updatetime) values(?,?,?,?,?,?,?,?,?,?)"
	errorcode.Id, err = db.Mconn.Insert(ssql, filepath.Clean(dir), header.Filename, true, header.Size, isuser,
		rwid, bugconfig.CacheRealNameUid[nickname], isuser,
		rwid,
		updatetime,
	)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 生成文件
	newfile := filepath.Join(basedir, header.Filename)

	f, err := os.OpenFile(newfile, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		golog.Error(err)
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

type rename struct {
	Path    string `json:"path"`
	Oldname string `json:"oldname"`
	Newname string `json:"newname"`
}

func ShareRename(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1 << 30)

	errorcode := &response.Response{}

	ps := &network.Data_sharefile{}

	pb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(pb, ps)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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
		golog.Error("has not permission")
		w.Write(errorcode.ErrorNoPermission())
		return
	}

	updatetime := time.Now().Unix()
	_, err = db.Mconn.Update("update sharefile set readuser=?,rid=?,isfile=?,writeuser=?,wid=?,updatetime=?,name=? where id=? and filepath=?",
		ps.ReadUser, rid, ps.IsFile, ps.WriteUser, wid, updatetime, ps.Name, ps.Id, ps.FilePath)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	newfile := filepath.Join(basedir, ps.Name)
	oldfile := filepath.Join(basedir, ps.OldName)
	err = os.Rename(oldfile, newfile)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	errorcode.UpdateTime = updatetime
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ShareDownload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1 << 30)

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
	//	golog.Error(err)
	//	w.Write(errorcode.ErrorNoPermission())
	//	return
	//}
	//defer f.Close()
	//
	//errorcode.Data, err = ioutil.ReadAll(f)
	//w.Header().Set("Content-Disposition", "attachment;filename="+ds.Name)
	//w.Header().Set("Content-Type", "application/octet-stream")
	//if err != nil {
	//	golog.Error(err)
	//	w.Write(errorcode.ErrorGetData())
	//	return
	//}
	//send, _ := json.Marshal(errorcode)
	//w.Write(send)
	return

}

func ShareMkdir(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1 << 30)

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	ps := &network.Data_sharefile{}
	pb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = json.Unmarshal(pb, ps)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
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
		golog.Error(err)
		w.Write(errorcode.ErrorNoPermission())
		return
	}
	updatetime := time.Now().Unix()
	ssql := "insert into sharefile(filepath,name,isfile,size,readuser,rid,writeuser,wid,ownerid,updatetime) values(?,?,?,?,?,?,?,?,?,?)"
	errorcode.Id, err = db.Mconn.Insert(ssql, filepath.Clean(ps.FilePath), ps.Name, ps.IsFile, ps.Size, ps.ReadUser, rid, ps.WriteUser, wid, bugconfig.CacheRealNameUid[nickname], updatetime)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	errorcode.UpdateTime = updatetime
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ShareRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	id := r.FormValue("id")
	// 检查文件夹权限
	var fp string
	var name string
	row, err := db.Mconn.GetOne("select filepath,name from sharefile where id=? and ownerid=?", id, bugconfig.CacheNickNameUid[nickname])
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&fp, &name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//还要递归删除子目录和文件
	//basedir := filepath.Join(bugconfig.ShareDir, fp, name)
	err = getrow(fp, name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	_, err = db.Mconn.Update("delete from sharefile where id=? ", id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = os.RemoveAll(filepath.Join(bugconfig.ShareDir, fp, name))
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func getrow(fp string, name string) error {
	dd, err := db.Mconn.GetRows("select id,filepath,name,isfile from sharefile where filepath=? ", filepath.Join(fp, name))
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
		_, err = db.Mconn.Update("delete from sharefile where id=? ", _id)
		if err != nil {
			return err
		}
		// 删除文件或文件夹

		if !_isfile {
			err = getrow(_filepath, _name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ShareShow(w http.ResponseWriter, r *http.Request) {

	//m := mux.Vars(r)
	//name := m["name"]

	id := r.FormValue("id")
	token := r.FormValue("token")

	// 只是验证token
	destoken, err := gaencrypt.RsaDecrypt(token, bugconfig.PrivateKey, true)
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	nickname, err := db.RSconn.Get(string(destoken))
	if err != nil {
		golog.Error(err)
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
	row, err := db.Mconn.GetOne(getsql, id)
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	err = row.Scan(&ownerid, &fp, &name, &readuser, &rid)
	if err != nil {
		golog.Error(err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	// 如果是所属用户
	if ownerid == uid {
		haspermision = true
	} else {
		// 如果有权限访问的用户
		if readuser && rid == uid {
			haspermision = true
		} else {
			// 判断权限组是否有权限
			var ids string
			row, err = db.Mconn.GetOne("select ids from usergroup where id=?", rid)
			if err != nil {
				golog.Error(err)
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			err = row.Scan(&ids)
			if err != nil {
				golog.Error(err)
				w.WriteHeader(http.StatusBadGateway)
				return
			}
			for _, v := range strings.Split(ids, ",") {
				if v == strconv.FormatInt(uid, 10) {
					haspermision = true
				}
			}
		}
	}

	if !haspermision {
		golog.Error("no permssion")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	f, err := ioutil.ReadFile(filepath.Join(bugconfig.ShareDir, fp, name))
	if err != nil {
		golog.Error(err)
		return
	}
	w.Header().Set("Content-Disposition", "attachment;filename="+name)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(f)
	return

}
