package handle

import (
	"encoding/json"
	"fmt"
	"itflow/model"
	"itflow/response"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Name(w http.ResponseWriter, r *http.Request) {
	// 检查name 是否重复
	errorcode := &response.Response{}
	name := r.URL.Query().Get("name")
	golog.Info(name)
	err := model.CheckName(name)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	errorcode.Success()
}

type ResponseDocs struct {
	Code int          `json:"code"`
	Doc  []*model.Doc `json:"doc"`
	Msg  string       `json:"msg"`
}

func DocList(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	var err error
	resp := &response.Response{}
	rd := &ResponseDocs{}
	rd.Doc, err = model.GetAllDocs(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return
	}
	send, err := json.Marshal(rd)
	if err != nil {
		golog.Error(err)
	}
	w.Write(send)
}

func Docs(w http.ResponseWriter, r *http.Request) {
	name := xmux.Var(r)["name"]
	golog.Info(r.URL.RequestURI())
	golog.Info(name)
	port := model.GetPort(name)
	if port == 0 {
		w.WriteHeader(404)
		return
	}
	host := fmt.Sprintf("localhost:%d", port)
	targetUrl, err := url.Parse("http://" + host)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	// Scheme      string
	// Opaque      string    // encoded opaque data
	// User        *Userinfo // username and password information
	// Host        string    // host or host:port
	// Path        string    // path (relative paths may omit leading slash)
	// RawPath     string    // encoded path hint (see EscapedPath method)
	// ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	// RawQuery    string    // encoded query values, without '?'
	// Fragment    string    // fragment for references, without '#'
	// RawFragment string
	r.URL.Scheme = "http"
	r.URL.Host = host
	r.URL.Path = r.URL.Path[len("/docs/"+name):]
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(w, r)

}

type Resp struct {
	Code int        `json:"code"`
	Doc  *model.Doc `json:"doc"`
	Msg  string     `json:"msg"`
}

func DocCreate(w http.ResponseWriter, r *http.Request) {
	// 创建文档， 先判断name是否合法
	// 判断 name 是都存在
	// 直接拉取如果成功 就 获取可用的端口， 使用scs 启动服务
	// 插入数据到数据库
	uid := xmux.GetInstance(r).Get("uid").(int64)
	doc := xmux.GetInstance(r).Data.(*model.Doc)
	errorcode := &response.Response{}
	_, err := regexp.MatchString("^[a-z]{1}[a-z0-9]{1,20}$", doc.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info("check document name")
	if err := model.ChecDocName(doc.Name); err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info("init a git object")
	_git, err := doc.NewGit()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info("git clone project")
	if err := _git.GitClone(); err != nil && err != git.ErrRepositoryAlreadyExists {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	//如果存在配置文件就 newpath 是 git的根目录, 创建目录
	golog.Info("start docsify server")
	err = doc.Startup()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	golog.Info("insert into mysql")
	err = doc.Insert(uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	rd := &Resp{
		Doc: doc,
	}
	send, _ := json.Marshal(rd)
	w.Write(send)
	return
}

// func DocDrop(w http.ResponseWriter, r *http.Request) {
// 	// 通过id 获取 name
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)
// 	id := r.URL.Query().Get("id")
// 	doc, err := model.NewDocById(id, uid)
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.Error(err))
// 		return
// 	}
// 	// 关闭服务
// 	doc.StopServer()

// 	// 删除文件夹
// 	err = doc.RemoveFolder()
// 	if err != nil {
// 		golog.Error(err)
// 	}
// 	// 删除数据
// 	err = doc.Delete()
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.Error(err))
// 		return
// 	}
// 	w.Write(errorcode.Success())
// 	return
// }

// func DocUpdate(w http.ResponseWriter, r *http.Request) {
// 	// 通过id 获取 name
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)
// 	id := r.URL.Query().Get("id")
// 	doc, err := model.NewDocById(id, uid)
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.Error(err))
// 		return
// 	}
// 	// 关闭服务
// 	_git, err := doc.NewGit()
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.Error(err))
// 		return
// 	}
// 	err = _git.GitPull()
// 	if err != nil {
// 		golog.Error(err)
// 		w.Write(errorcode.Error(err))
// 		return
// 	}

// 	w.Write(errorcode.Success())
// 	return
// }

// func DocShowFiles(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")

// 	// 先获取did

// 	//验证权限
// 	var select_uid int64
// 	var wl string
// 	err := varconfig.Mysql_Admin.GetOne("select uid, writelist from doc where id=?", did).Scan(&select_uid, &wl)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	golog.Debug("uid: ", uid)
// 	golog.Debug("wl: ", wl)
// 	var hasperm bool
// 	if select_uid == uid {
// 		hasperm = true
// 	} else {
// 		users := strings.Split(wl, ",")
// 		if users[0] != "" {
// 			for _, v := range users {
// 				wluid, err := strconv.ParseInt(v, 10, 64)
// 				if err != nil {
// 					golog.Error("字段writelst值有问题，请修正， ID: %d", select_uid)
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				if wluid == uid {
// 					hasperm = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	if hasperm {
// 		rows, err := varconfig.Mysql_Admin.GetRows("select file from docfile where did=? ", did)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorConnentMysql())
// 			return
// 		}
// 		ss := &struct {
// 			Files []string `json:"files"`
// 			Code  int      `json:"code"`
// 		}{}
// 		ss.Files = make([]string, 0)
// 		for rows.Next() {
// 			var file string
// 			rows.Scan(&file)
// 			ss.Files = append(ss.Files, file)
// 		}
// 		send, _ := json.Marshal(ss)
// 		w.Write(send)
// 		return
// 	}
// 	golog.Error("hasperm false")
// 	w.Write(errorcode.ErrorNoPermission())
// 	return
// }

// func DocSaveContent(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")
// 	name := r.FormValue("name")
// 	// 权限uid验证
// 	var select_id int64
// 	var domain string
// 	err := varconfig.Mysql_Admin.GetOne("select uid,name from doc where id=?", did).Scan(&select_id, &domain)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	if select_id != uid {
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	// 验证文件权限
// 	var id int64
// 	err = varconfig.Mysql_Admin.GetOne("select id from docfile where did=? and file=?", did, name).Scan(&id)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}

// 	data := &struct {
// 		Code int    `json:"code"`
// 		Data string `json:"data"`
// 	}{}
// 	path := filepath.Join(goconfig.ReadString("scs.mount"), domain, name)
// 	fmt.Println(path)
// 	bb, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorFileNotFound())
// 		return
// 	}
// 	data.Data = string(bb)
// 	fmt.Println(data.Data)
// 	send, err := json.Marshal(data)
// 	w.Write(send)
// 	return
// }

// func DocGetContent(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")
// 	name := r.FormValue("name")
// 	// 权限uid验证
// 	var select_id int64
// 	var domain string
// 	var wl string
// 	err := varconfig.Mysql_Admin.GetOne("select uid,name,writelist from doc where id=?", did).Scan(&select_id, &domain, &wl)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.Error(err))
// 		return
// 	}
// 	var hasperm bool
// 	if select_id == uid {
// 		hasperm = true
// 	} else {
// 		users := strings.Split(wl, ",")
// 		if users[0] != "" {
// 			for _, v := range users {
// 				wluid, err := strconv.ParseInt(v, 10, 64)
// 				if err != nil {
// 					golog.Error("字段writelst值有问题，请修正， ID: %d", did)
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				if wluid == uid {
// 					hasperm = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	// 验证文件权限
// 	if hasperm {
// 		var id int64
// 		err = varconfig.Mysql_Admin.GetOne("select id from docfile where did=? and file=?", did, name).Scan(&id)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorConnentMysql())
// 			return
// 		}

// 		path := filepath.Join(goconfig.ReadString("scs.mount"), domain, name)
// 		data := &struct {
// 			Data string `json:"data"`
// 			Code int    `json:"code"`
// 		}{}
// 		bb, err := ioutil.ReadFile(path)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorFileNotFound())
// 			return
// 		}
// 		data.Data = string(bb)
// 		send, _ := json.Marshal(data)
// 		w.Write(send)
// 		return
// 	}
// 	w.Write(errorcode.ErrorNoPermission())
// 	return
// }

// func DocSaveFile(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")
// 	name := r.FormValue("name")
// 	content := r.FormValue("content")
// 	// 权限uid验证
// 	var select_id int64
// 	var domain string
// 	var wl string
// 	err := varconfig.Mysql_Admin.GetOne("select uid,name,writelist from doc where id=?", did).Scan(&select_id, &domain, &wl)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	if select_id != uid {
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	var hasperm bool
// 	if select_id == uid {
// 		hasperm = true
// 	} else {
// 		users := strings.Split(wl, ",")
// 		if users[0] != "" {
// 			for _, v := range users {
// 				wluid, err := strconv.ParseInt(v, 10, 64)
// 				if err != nil {
// 					golog.Error("字段writelst值有问题，请修正， ID: %d", select_id)
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				if wluid == uid {
// 					hasperm = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	// 验证文件权限
// 	if hasperm {
// 		var id int64
// 		err = varconfig.Mysql_Admin.GetOne("select id from docfile where did=? and file=?", did, name).Scan(&id)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorConnentMysql())
// 			return
// 		}
// 		// 写入文件
// 		path := filepath.Join(goconfig.ReadString("scs.mount"), domain, name)
// 		err = ioutil.WriteFile(path, []byte(content), 0644)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorFileNotFound())
// 			return
// 		}

// 		w.Write(errorcode.Success())
// 		return
// 	}
// 	w.Write(errorcode.ErrorNoPermission())
// 	return
// }

// func DocCreateFile(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")
// 	name := r.FormValue("name")
// 	// 权限uid验证
// 	var select_id int64
// 	var domain, wl string
// 	err := varconfig.Mysql_Admin.GetOne("select uid,name,writelist from doc where id=?", did).Scan(&select_id, &domain, &wl)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	var hasperm bool
// 	if select_id == uid {
// 		hasperm = true
// 	} else {
// 		users := strings.Split(wl, ",")
// 		if users[0] != "" {
// 			for _, v := range users {
// 				wluid, err := strconv.ParseInt(v, 10, 64)
// 				if err != nil {
// 					golog.Error("字段writelst值有问题，请修正， ID: %d", select_id)
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				if wluid == uid {
// 					hasperm = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	if hasperm {
// 		//创建文件, 如果这里错误，那么就是有路径
// 		path := filepath.Join(goconfig.ReadString("scs.mount"), domain, name)
// 		//先判断是否存在文件
// 		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
// 			w.Write(errorcode.ErrorExsit())
// 			return
// 		}
// 		f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorNoPermission())
// 			return
// 		}
// 		defer f.Close()
// 		// 文件加入到数据库
// 		errorcode.ID, err = varconfig.Mysql_Admin.Insert("insert into docfile(did, file) values(?,?) ", did, name)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorConnentMysql())
// 			return
// 		}

// 		w.Write(errorcode.Success())
// 		return
// 	}
// 	w.Write(errorcode.ErrorNoPermission())
// 	return
// }

// func DocDeleteFile(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")
// 	name := r.FormValue("name")
// 	// 权限uid验证
// 	var select_id int64
// 	var domain, wl string
// 	if name == "SUMMARY.md" || name == "README.md" || name == "book.json" {
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	err := varconfig.Mysql_Admin.GetOne("select uid,name,writelist from doc where id=?", did).Scan(&select_id, &domain, &wl)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	var hasperm bool
// 	if select_id == uid {
// 		hasperm = true
// 	} else {
// 		users := strings.Split(wl, ",")
// 		if users[0] != "" {
// 			for _, v := range users {
// 				wluid, err := strconv.ParseInt(v, 10, 64)
// 				if err != nil {
// 					golog.Error("字段writelst值有问题，请修正， ID: %d", select_id)
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				if wluid == uid {
// 					hasperm = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	if hasperm {
// 		//删除文件
// 		path := filepath.Join(goconfig.ReadString("scs.mount"), domain, name)
// 		err = os.Remove(path)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorNoPermission())
// 			return
// 		}
// 		// 文件加入到数据库
// 		_, err = varconfig.Mysql_Admin.Delete("delete from  docfile where did=? and file=? ", did, name)
// 		if err != nil {
// 			golog.Error(err.Error())
// 			w.Write(errorcode.ErrorConnentMysql())
// 			return
// 		}

// 		w.Write(errorcode.Success())
// 		return
// 	}
// 	w.Write(errorcode.ErrorNoPermission())
// 	return
// }

// func DocDownload(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	did := r.FormValue("id")

// 	var domain string
// 	err := varconfig.Mysql_Admin.GetOne("select name from doc where id=? and uid=?", did, uid).Scan(&domain)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	d := command.NewDocker(domain)
// 	b, err := d.DownloadTar()
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	bb := &struct {
// 		Data []byte `json:"data"`
// 		Code int    `json:"code"`
// 	}{}
// 	bb.Data = b
// 	send, _ := json.Marshal(bb)
// 	w.Write(send)
// 	return
// }

// func DocUpload(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	if err := r.ParseMultipartForm(1 << 20 * 100); err != nil {
// 		w.Write(errorcode.ErrorDateTooLarge())
// 		return
// 	}
// 	//验证
// 	id := r.FormValue("id")
// 	domain := r.FormValue("name")
// 	var select_id int64
// 	if err := varconfig.Mysql_Admin.GetOne("select uid from doc where id=? and name=?", id, domain).Scan(&select_id); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	if select_id != uid {
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	file, _, err := r.FormFile("file")
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorGetInstance())
// 		return
// 	}
// 	data, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorGetInstance())
// 		return
// 	}
// 	//保存文件
// 	savefile := filepath.Join("/", domain+".tar")
// 	if err := ioutil.WriteFile(savefile, data, 0644); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorGetInstance())
// 		return
// 	}
// 	d := command.NewDocker(domain)
// 	nfm := fm.NewFileManager()
// 	nfm.Name = savefile
// 	nfm.Pack = "book"
// 	fn := filepath.Join(d.Mount, domain)
// 	nfm.SrcDir = fn
// 	if err := nfm.UnTar(); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorGetInstance())
// 		return
// 	}
// 	w.Write(errorcode.Success())
// 	return
// }

// func DocUser(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	var select_id int64
// 	// 权限验证， 只有所属者才可以编辑用户
// 	id := r.FormValue("id")
// 	var wl string
// 	if err := varconfig.Mysql_Admin.GetOne("select uid,writelist from doc where id=?",
// 		id,
// 	).Scan(&select_id, &wl); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	if uid != select_id {
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	//
// 	data := &struct {
// 		Wl   []string `json:"wl"`
// 		Code int      `json:"code"`
// 	}{}
// 	for _, v := range strings.Split(wl, ",") {

// 		wluid, err := strconv.ParseInt(v, 10, 64)
// 		if err != nil {
// 			golog.Error("数据库字段writelist的值不对，id: %s", id)
// 			w.Write(errorcode.ErrorConnentMysql())
// 			return
// 		}

// 		data.Wl = append(data.Wl, varconfig.CachedUidNickName[wluid])
// 	}

// 	send, _ := json.Marshal(data)
// 	w.Write(send)
// 	return
// }

// func DocUserAdd(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	var select_id int64
// 	// 权限验证， 只有所属者才可以添加用户
// 	id := r.FormValue("id")
// 	name := r.FormValue("name")
// 	name = strings.Trim(name, " ")
// 	var writelist string
// 	if err := varconfig.Mysql_Admin.GetOne("select uid, writelist from doc where id=?",
// 		id,
// 	).Scan(&select_id, &writelist); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	if uid != select_id {
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	// 确认是否存在这个用户
// 	var hasperm bool

// 	for _, v := range varconfig.CachedUidNickName {
// 		if v == name {
// 			hasperm = true
// 			break
// 		}
// 	}
// 	if hasperm {
// 		//如果用户存在
// 		var hasuser bool
// 		// 记录uid
// 		var add_uid string
// 		users := strings.Split(writelist, ",")
// 		if users[0] != "" {
// 			for _, v := range users {
// 				uid64, err := strconv.ParseInt(v, 10, 64)
// 				if err != nil {
// 					golog.Error("数据库字段writelist的值不对，id: %s", id)
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				if varconfig.CachedUidNickName[uid64] == name {
// 					hasuser = true
// 					break
// 				}
// 			}

// 			if hasuser {
// 				w.Write(errorcode.ErrorExsit())
// 				return
// 			}
// 			for id, v := range varconfig.CachedUidNickName {
// 				if v == name {
// 					add_uid = strconv.FormatInt(id, 10)
// 					break
// 				}
// 			}
// 			users = append(users, add_uid)
// 			insertwl := strings.Join(users, ",")
// 			_, err := varconfig.Mysql_Admin.Update("update doc set writelist=? where id=?", insertwl, id)
// 			if err != nil {
// 				golog.Error(err.Error())
// 				w.Write(errorcode.ErrorConnentMysql())
// 				return
// 			}

// 			//添加用户
// 			w.Write(errorcode.Success())
// 			return
// 		}

// 	} else {
// 		w.Write(errorcode.ErrorUserNotFound())
// 		return
// 	}

// }

// func DocUserDel(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &errorStruct{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	var select_id int64
// 	// 权限验证， 只有所属者才可以添加用户
// 	id := r.FormValue("id")
// 	name := r.FormValue("name")
// 	name = strings.Trim(name, " ")
// 	var writelist string
// 	if err := varconfig.Mysql_Admin.GetOne("select uid, writelist from doc where id=?",
// 		id,
// 	).Scan(&select_id, &writelist); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	if uid != select_id {

// 		w.Write(errorcode.ErrorNoPermission())
// 		return
// 	}
// 	// 确认是否存在这个用户
// 	var hasperm bool

// 	for _, v := range varconfig.CachedUidNickName {
// 		if v == name {
// 			hasperm = true
// 			break
// 		}
// 	}
// 	if hasperm {
// 		//如果用户存在
// 		//var hasuser bool
// 		// 记录uid

// 		users := strings.Split(writelist, ",")
// 		if users[0] == "" {
// 			golog.Error("没有用户")
// 			w.Write(errorcode.ErrorUserNotFound())
// 			return
// 		}
// 		for k, v := range users {
// 			uid64, err := strconv.ParseInt(v, 10, 64)
// 			if err != nil {
// 				golog.Error("数据库字段writelist的值不对，id: %s", id)
// 				w.Write(errorcode.ErrorConnentMysql())
// 				return
// 			}
// 			if varconfig.CachedUidNickName[uid64] == name {

// 				//删除
// 				tmp := make([]string, 0)
// 				tmp = append(tmp, users[:k]...)
// 				tmp = append(tmp, users[k+1:]...)
// 				_, err = varconfig.Mysql_Admin.Update("update doc set writelist=? where id=?",
// 					strings.Join(tmp, ","), id)
// 				if err != nil {
// 					golog.Error(err.Error())
// 					w.Write(errorcode.ErrorConnentMysql())
// 					return
// 				}
// 				w.Write(errorcode.Success())
// 				return
// 			}
// 		}
// 	}
// 	w.Write(errorcode.ErrorUserNotFound())
// 	return
// }

// func DocAddGit(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &response.Response{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)

// 	data := &struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	}{}

// 	bd, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorGetInstance())
// 		return
// 	}
// 	golog.Debug(string(bd))

// 	if err := json.Unmarshal(bd, data); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorParams())
// 		return
// 	}
// 	name := strings.Trim(data.Name, " ")
// 	ok, err := regexp.MatchString("^[a-z]{1}[a-z0-9]{1,20}$", name)
// 	if err != nil || !ok {
// 		golog.Error("string %s not format", name)
// 		w.Write(errorcode.ErrorGetInstance())
// 		return
// 	}

// 	//如果存在配置文件就return
// 	newpath := filepath.Join(goconfig.ReadPath("scs.path"), name+".conf")
// 	if _, err := os.Stat(newpath); err == nil {
// 		//如果存在就返回
// 		golog.Error("this conf is exsit , remove first")
// 		w.Write(errorcode.ErrorExsit())
// 		return
// 	}

// 	// 更新docker-compose, 会启动新的
// 	dc := dockercompose.NewYaml()
// 	dcfile := filepath.Join(goconfig.ReadString("scs.compose"), "docker-compose.yaml")
// 	if err := dc.Load(dcfile); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorDocker())
// 		return
// 	}
// 	if _, ok := dc.Services[name]; ok {
// 		golog.Error("存在此项目，请更换名字")
// 		errorcode.Msg = "存在此项目，请更换名字"
// 		w.Write(errorcode.ErrorExsit())
// 		return
// 	}

// 	// 启动docker
// 	d := command.NewDocker(name)
// 	d.UserName = varconfig.CachedUidNickName[uid]
// 	d.Email = varconfig.CacheUidEmail[uid]

// 	//获取用户名和email
// 	if err := d.Clone(data.Url); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorDocker())
// 		return
// 	}
// 	// 创建doc
// 	date := time.Now().Unix()
// 	getcategorysql := "insert into doc(uid,created,name,git,giturl) values(?,?,?,true,?)"
// 	did, err := varconfig.Mysql_Admin.Insert(getcategorysql, uid, date, name, data.Url)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	dc.AddService(name, dockercompose.DocPod(name))
// 	// 增加links
// 	dc.Services["nginx"].Links = append(dc.Services["nginx"].Links, name)
// 	if err := dc.WriteFile(dcfile); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorDocker())
// 		return
// 	}

// 	// 更新重启nginx
// 	if err := d.Add(); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorDocker())
// 		return
// 	}

// 	errorcode.ID = did
// 	errorcode.Date = date
// 	w.Write(errorcode.Success())
// 	return

// }

// func DocGitPull(w http.ResponseWriter, r *http.Request) {
// 	errorcode := &response.Response{}
// 	uid := xmux.GetInstance(r).Get("uid").(int64)
// 	// 权限验证， 只有所属者才可以添加用户
// 	id := r.FormValue("id")

// 	//name := strings.Trim(data.Name, " ")
// 	//ok, err := regexp.MatchString("^[a-z]{1}[a-z0-9]{1,20}$", name)
// 	//if err != nil || !ok {
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorGetInstance())
// 	//	return
// 	//}
// 	//如果存在配置文件就return
// 	//newpath := filepath.Join(goconfig.ReadString("scs.path"), name+".conf")
// 	//if _, err := os.Stat(newpath); err == nil {
// 	//	//如果存在就返回
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorExsit())
// 	//	return
// 	//}

// 	// 更新docker-compose, 会启动新的
// 	//dc := dockercompose.NewYaml()
// 	//dcfile := filepath.Join(goconfig.ReadString("scs.compose"), "docker-compose.yaml")
// 	//if err := dc.Load(dcfile); err != nil {
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorDocker())
// 	//	return
// 	//}
// 	//if _, ok := dc.Services[name]; ok {
// 	//	golog.Error(err.Error())
// 	//	errorcode.Msg = "存在此项目，请更换名字"
// 	//	w.Write(errorcode.ErrorExsit())
// 	//	return
// 	//}

// 	// 创建doc的基础名称  没有编辑文件功能

// 	var name string
// 	err := varconfig.Mysql_Admin.GetOne("select name from doc where id=?", id).Scan(&name)
// 	if err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorConnentMysql())
// 		return
// 	}
// 	// 启动docker
// 	d := command.NewDocker(name)
// 	d.UserName = varconfig.CachedUidNickName[uid]
// 	d.Email = varconfig.CacheUidEmail[uid]
// 	//if err = d.DefaultFile(); err != nil {
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorConnentMysql())
// 	//	return
// 	//}
// 	//获取用户名和email
// 	if err := d.Pull(); err != nil {
// 		golog.Error(err.Error())
// 		w.Write(errorcode.ErrorDocker())
// 		return
// 	}
// 	// 创建doc
// 	//date := time.Now().Unix()
// 	//getcategorysql := "insert into doc(uid,created,name,git,giturl) values(?,?,?,true,?)"
// 	//did, err := varconfig.Mysql_Admin.Insert(getcategorysql, uid, date, name, data.Url)
// 	//if err != nil {
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorConnentMysql())
// 	//	return
// 	//}
// 	//dc.AddService(name, dockercompose.DocPod(name))
// 	//// 增加links
// 	//dc.Services["nginx"].Links = append(dc.Services["nginx"].Links, name)
// 	//if err := dc.WriteFile(dcfile); err != nil {
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorDocker())
// 	//	return
// 	//}

// 	// 更新重启nginx
// 	//if err := d.Add(); err != nil {
// 	//	golog.Error(err.Error())
// 	//	w.Write(errorcode.ErrorDocker())
// 	//	return
// 	//}
// 	//
// 	//errorcode.ID = did
// 	//errorcode.Date = date
// 	w.Write(errorcode.Success())
// 	return

// }

func ProxyDoc(w http.ResponseWriter, r *http.Request) {
	name := xmux.Var(r)["name"]
	golog.Info(r.URL.RequestURI())
	golog.Info(r.URL.RawQuery)
	golog.Info(name)
	port := model.GetPort(name)
	if port == 0 {
		w.WriteHeader(404)
		return
	}
	host := fmt.Sprintf("localhost:%d", port)
	targetUrl, err := url.Parse("http://" + host)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	// Scheme      string
	// Opaque      string    // encoded opaque data
	// User        *Userinfo // username and password information
	// Host        string    // host or host:port
	// Path        string    // path (relative paths may omit leading slash)
	// RawPath     string    // encoded path hint (see EscapedPath method)
	// ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	// RawQuery    string    // encoded query values, without '?'
	// Fragment    string    // fragment for references, without '#'
	// RawFragment string
	r.URL.Scheme = "http"
	r.URL.Host = host
	r.URL.Path = r.URL.Path[len("/docs/"+name):]
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(w, r)

}

func DocUpdate(w http.ResponseWriter, r *http.Request) {
	// 通过id 获取 name
	errorcode := &response.Response{}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	id := r.URL.Query().Get("id")
	doc, err := model.NewDocById(id, uid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 关闭服务
	_git, err := doc.NewGit()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = _git.GitPull()
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write(errorcode.Success())
}
