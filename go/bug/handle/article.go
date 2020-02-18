package handle

import (
	"database/sql"
	"encoding/json"
	"html"
	"io"
	"io/ioutil"
	"itflow/bug/bugconfig"
	"itflow/db"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hyahm/golog"
)

type ArticleList struct {
	id int
}

type Article struct {
	Items *ArticleList
	Total int
}

type articledetail struct {
	ID          int    `json:"id"`
	Importance  string `json:"importance"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Spusers     string `json:"spusers"`
	Selectoses  string `json:"selectoses"`
	AppVersion  string `json:"appversion"`
	Content     string `json:"content"`
	Level       string `json:"level"`
	Projectname string `json:"projectname"`
}

//func Detail(w http.ResponseWriter, r *http.Request) {
//	headers(w, r)
//	if r.Method == http.MethodOptions {
//		w.WriteHeader(http.StatusOK)
//		return
//	}
//
//	if r.Method == http.MethodGet {
//		bid := r.FormValue("id")
//		id, err := strconv.Atoi(bid)
//		logger, conn, name, err := logtokenmysql(r)
//
//		uid, err := asset.NicknameGetUid(name, conn)
//		if err != nil {
//			logger.ErrorLog(err)
//			conn.Db.Close()
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		getlistsql := "select bugtitle,status,content,importent,selectclass,appversion,selectoses,spusers,level,pid from bugs where id=? and uid=?"
//		rows, err := conn.GetRows(getlistsql, bid, uid)
//		if err != nil {
//			logger.ErrorLog("file:article.go,line:70,%v", err)
//			conn.Db.Close()
//			w.Write([]byte("fail"))
//			return
//		}
//		for rows.Next() {
//			rows.Scan()
//		}
//		pname, err := asset.PidGetProject(data[0][9], conn)
//
//		ad := &articledetail{}
//
//		ad.ID = id
//		statusid, err := asset.StatusidGetName(data[0][1], conn)
//		if err != nil {
//			logger.ErrorLog("file:article.go,line:111,%v", err)
//			conn.Db.Close()
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//
//		userlist := strings.Split(data[0][7], ",")
//		ul := ""
//
//		for _, v := range userlist {
//			user, err := asset.UidGetNicknameAndRealname(v, conn)
//			if err != nil {
//				logger.ErrorLog("file:article.go,line:122,%v", err)
//				conn.Db.Close()
//				w.WriteHeader(http.StatusNotFound)
//				return
//			}
//			if ul == "" {
//				ul = user
//			} else {
//				ul = ul + "," + user
//			}
//		}
//
//		ad.Status = statusid
//		ad.Title = data[0][0]
//		ad.Content = data[0][2]
//		ad.Projectname = pname
//		ad.Importance = data[0][3]
//		ad.SelectClass = data[0][4]
//		ad.AppVersion = data[0][5]
//		ad.Selectoses = data[0][6]
//		ad.Spusers = ul
//		ad.Level = data[0][8]
//		send, err := json.Marshal(ad)
//		if err != nil {
//			logger.ErrorLog("file:article.go,line:118,%v", err)
//			conn.Db.Close()
//			w.Write([]byte("fail"))
//			return
//		}
//		conn.Db.Close()
//		w.Write(send)
//		return
//	}
//	w.WriteHeader(http.StatusNotFound)
//}

type envList struct {
	EnvList []string `json:"envlist"`
	Code    int      `json:"code"`
}

func GetEnv(w http.ResponseWriter, r *http.Request) {

	errorcode := &errorstruct{}
	_, err := logtokenmysql(r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	el := &envList{}

	for _, v := range bugconfig.CacheEidName {
		el.EnvList = append(el.EnvList, v)
	}
	send, _ := json.Marshal(el)
	w.Write(send)
	return

}

type senduserinfo struct {
	Nickname string `json:"nickname"`
	Realname string `json:"realname"`
}

// 用户名和真实名称
type nickreal struct {
	NickName string `json:"nickname"`
	RealName string `json:"realname"`
}

type userList struct {
	Users []string `json:"users"`
	Code  int      `json:"code"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	ul := &userList{}

	getusersql := "select realname from user"
	rows, err := db.Mconn.GetRows(getusersql)

	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	for rows.Next() {
		var realname string
		rows.Scan(&realname)
		ul.Users = append(ul.Users, realname)
	}
	send, _ := json.Marshal(ul)
	w.Write(send)
	return

}

type versionList struct {
	VersionList []string `json:"versionlist"`
	Code        int      `json:"code"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {

	_, err := logtokenmysql(r)
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	vl := &versionList{}

	for _, v := range bugconfig.CacheVidName {
		vl.VersionList = append(vl.VersionList, v)
	}
	send, _ := json.Marshal(vl)
	w.Write(send)
	return

}

type getArticle struct {
	Status      string   `json:"status"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Id          int      `json:"id"`
	Selectusers []string `json:"selectuser"`
	Important   string   `json:"important"`
	Level       string   `json:"level"`
	Projectname string   `json:"projectname"`
	Envname     string   `json:"envname"`
	Version     string   `json:"version"`
}

type uploadImage struct {
	HasSuccess bool   `json:"hasSuccess"`
	Height     int    `json:"height"`
	Uid        uint64 `json:"uid"`
	Url        string `json:"url"`
	Width      int    `json:"width"`
}

func UploadImgs(w http.ResponseWriter, r *http.Request) {
	errorcode := &errorstruct{}
	file, h, err := r.FormFile("file")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	ext := filepath.Ext(h.Filename)
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext

	cfile, err := os.OpenFile(path.Join(bugconfig.ImgDir, filename), os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	defer cfile.Close()

	_, err = io.Copy(cfile, file)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	l := len(bugconfig.ShowBaseUrl)
	url := ""
	if bugconfig.ShowBaseUrl[l-1:l] == "/" {
		url = bugconfig.ShowBaseUrl + filename
	} else {
		url = bugconfig.ShowBaseUrl + "/" + filename
	}

	sendurl := &uploadImage{
		HasSuccess: true,
		Url:        url,
	}
	send, _ := json.Marshal(sendurl)
	w.Write(send)
	return

}

type informations struct {
	User string `json:"user"`
	Date int64  `json:"date"`
	Info string `json:"info"`
}

type showArticle struct {
	Title      string          `json:"title"`
	Content    string          `json:"content"`
	Appversion string          `json:"appversion"`
	Comment    []*informations `json:"comment"`
	Status     string          `json:"status"`
	Id         int             `json:"id"`
	Code       int             `json:"code"`
}

type uploadimage struct {
	Uploaded int    `json:"uploaded"`
	Url      string `json:"url"`
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
}

func UploadHeadImg(w http.ResponseWriter, r *http.Request) {
	url := &uploadimage{}
	nickname, err := logtokenmysql(r)
	golog.Info("uploading header image")
	errorcode := &errorstruct{}
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	image, header, err := r.FormFile("upload")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	imgcode := make([]byte, header.Size)
	_, err = image.Read(imgcode)
	if err != nil {
		golog.Errorf("parse uploadImage struct fail,%v", err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	prefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	filename := prefix + ".png"
	err = ioutil.WriteFile(path.Join(bugconfig.ImgDir, filename), imgcode, 0655) //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	ul := len(bugconfig.ShowBaseUrl)
	if bugconfig.ShowBaseUrl[ul-1:ul] == "/" {
		url.Url = bugconfig.ShowBaseUrl + filename
	} else {
		url.Url = bugconfig.ShowBaseUrl + "/" + filename
	}

	url.FileName = filename
	url.Uploaded = 1
	uploadimg := "update user set headimg = ? where nickname=?"

	_, err = db.Mconn.Update(uploadimg, url.Url, nickname)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	s, _ := json.Marshal(url)

	w.Write(s)
	return

}

func BugShow(w http.ResponseWriter, r *http.Request) {
	bid := r.FormValue("id")
	sl := &showArticle{}
	errorcode := &errorstruct{}
	_, err := logtokenmysql(r)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	getinfosql := "select uid,info,time from informations where bid=?"
	rows, err := db.Mconn.GetRows(getinfosql, bid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		im := &informations{}
		var uid int64
		rows.Scan(&uid, &im.Info, &im.Date)
		im.User = bugconfig.CacheUidRealName[uid]
		sl.Comment = append(sl.Comment, im)
	}

	getlistsql := "select bugtitle,content,vid,sid,id from bugs where id=?"
	var statusid int64
	var vid int64
	row, err := db.Mconn.GetOne(getlistsql, bid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = row.Scan(&sl.Title, &sl.Content, &vid, &statusid, &sl.Id)
	if err != nil && err != sql.ErrNoRows {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	sl.Status = bugconfig.CacheSidStatus[statusid]
	sl.Appversion = bugconfig.CacheVidName[vid]
	sl.Content = html.UnescapeString(sl.Content)
	send, _ := json.Marshal(sl)
	w.Write(send)
	return
}
