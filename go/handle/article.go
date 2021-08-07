package handle

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/comment"
	"itflow/model"
	"itflow/response"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
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

type envList struct {
	EnvList []string `json:"envlist"`
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
}

func (el *envList) Marshal() []byte {
	send, err := json.Marshal(el)
	if err != nil {
		golog.Error(err)

	}
	return send
}

func (el *envList) Error(msg string) []byte {
	el.Code = 1
	el.Msg = msg
	return el.Marshal()
}
func (el *envList) ErrorE(err error) []byte {
	return el.Error(err.Error())
}

func GetEnv(w http.ResponseWriter, r *http.Request) {
	el := &envList{
		EnvList: make([]string, 0),
	}

	rows, err := db.Mconn.GetRows("select name from environment")
	if err != nil {
		golog.Error(err)
		w.Write(el.ErrorE(err))
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Info(err)
			continue
		}
		el.EnvList = append(el.EnvList, name)
	}
	rows.Close()
	w.Write(el.Marshal())
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
	Msg   string   `json:"msg"`
}

func (ul *userList) Marshal() []byte {
	send, err := json.Marshal(ul)
	if err != nil {
		golog.Error(err)
	}
	return send
}

func (ul *userList) ErrorE(err error) []byte {
	ul.Code = 1
	ul.Msg = err.Error()
	return ul.Marshal()
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	ul := &userList{
		Users: make([]string, 0),
	}
	rows, err := db.Mconn.GetRows("select realname from user")
	if err != nil {
		golog.Error(err)
		w.Write(ul.ErrorE(err))
		return
	}
	for rows.Next() {
		realname := new(string)
		// var realname string
		err = rows.Scan(realname)
		if err != nil {
			golog.Info(err)
			continue
		}
		ul.Users = append(ul.Users, *realname)
	}
	rows.Close()
	send, _ := json.Marshal(ul)
	w.Write(send)
	return

}

func IsAdmin(w http.ResponseWriter, r *http.Request) {
	if xmux.GetInstance(r).Get("uid").(int64) == cache.SUPERID {
		w.Write([]byte(`{"code": 0 , "admin": true}`))
	} else {
		w.Write([]byte(`{"code": 0 , "admin": false}`))
	}
}

// func GetEmail(w http.ResponseWriter, r *http.Request) {
// 	if xmux.GetInstance(r).Get("uid").(int64) == cache.SUPERID {
// 		w.Write([]byte("1"))
// 	} else {
// 		w.Write([]byte("0"))
// 	}
// 	return
// }

func GetProjectUser(w http.ResponseWriter, r *http.Request) {
	// 通过project 来获取所属用户
	pid := xmux.Var(r)["id"]
	projectId, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{}
	// 通过project_id 获取对应的version_ids
	res.VersionIds, err = model.GetVersionIdsByProjectId(projectId)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	// 通过project_id 获取对应的user_id
	// 获取usergroupid
	ug := model.UserGroup{}
	ug.Id, err = model.GetUserGroupId(projectId)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}

	err = ug.GetUserIds()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res.UserIds = ug.Uids
	w.Write(res.Marshal())
}

type versionList struct {
	VersionList []string `json:"versionlist"`
	Code        int      `json:"code"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {

	vl := &versionList{
		VersionList: make([]string, 0),
	}

	// for _, v := range cache.CacheVidVersion {
	// 	vl.VersionList = append(vl.VersionList, v)
	// }
	send, _ := json.Marshal(vl)
	w.Write(send)
	return

}

type uploadImage struct {
	HasSuccess bool   `json:"hasSuccess"`
	Height     int    `json:"height"`
	Uid        int64  `json:"uid"`
	Url        string `json:"url"`
	Width      int    `json:"width"`
	Code       int    `json:"code"`
}

func UploadImgs(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}

	file, h, err := r.FormFile("image")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	ext := filepath.Ext(h.Filename)
	golog.Info("upload image")
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	p := path.Join(cache.ImgDir, filename)
	golog.Info(p)
	cfile, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0755)
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

	url := cache.ShowBaseUrl + filename

	sendurl := &uploadImage{
		HasSuccess: true,
		Url:        url,
		Uid:        time.Now().UnixNano(),
	}
	send, _ := json.Marshal(sendurl)
	w.Write(send)
	return

}

type uploadimage struct {
	Uploaded int    `json:"uploaded"`
	Url      string `json:"url"`
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
}

func UploadHeadImg(w http.ResponseWriter, r *http.Request) {
	url := &uploadimage{}
	golog.Info("uploading header image")
	errorcode := &response.Response{}
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
	err = ioutil.WriteFile(path.Join(cache.ImgDir, filename), imgcode, 0655) //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	url.Url = cache.ShowBaseUrl + filename

	url.FileName = filename
	url.Uploaded = 1
	uploadimg := "update user set headimg = ? where id=?"
	uid := xmux.GetInstance(r).Get("uid").(int64)
	golog.Info(uid)
	_, err = db.Mconn.Update(uploadimg, url.Url, uid)
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
	// 判断是否有权限访问这个bug
	bid := r.FormValue("id")
	sl := &bug.RespShowBug{
		Comments: make([]*comment.Informations, 0),
		Url:      make([]string, 0),
	}
	golog.Info(bid)
	var u1, u2 string
	err := db.Mconn.GetOne(`select b.tid, b.id,b.content,ifnull(i.name, ''), ifnull(l.name, ''), 
	ifnull(e.name, ''), b.title, p.name, ifnull(v.name, ''), ifnull(v.urlone, ''), ifnull(v.urltwo, '') from bugs as b 
	left join importants as i on i.id=b.iid 
	left join version as v on b.vid=v.id 
	left join level as l on b.lid=l.id 
	left join environment as e on b.eid=e.id 
	left join project as p on b.pid=p.id 
	left join user as u on b.uid=u.id  
	left join status as s on   b.sid=s.id  
	where b.id=? `, bid).Scan(
		&sl.Typ, &sl.Id, &sl.Content, &sl.Important, &sl.Level, &sl.Envname, &sl.Title, &sl.Projectname, &sl.Version, &u1, &u2,
	)
	if err != nil {
		golog.Error(err)
		w.Write(sl.ErrorE(err))
		return
	}
	golog.Infof("%+v", sl)
	if u1 != "" {
		sl.Url = append(sl.Url, u1)
	}
	if u2 != "" {
		sl.Url = append(sl.Url, u2)
	}
	getinfosql := "select u.realname,info,time from informations as i join user as u on bid=? and u.id=i.uid"
	rows, err := db.Mconn.GetRows(getinfosql, bid)
	if err != nil {
		golog.Error(err)
		w.Write(sl.ErrorE(err))
		return
	}
	for rows.Next() {
		im := &comment.Informations{}
		// var uid int64
		rows.Scan(&im.User, &im.Info, &im.Date)
		// im.User = cache.CacheUidRealName[uid]
		sl.Comments = append(sl.Comments, im)
	}
	rows.Close()

	send, _ := json.Marshal(sl)
	w.Write(send)
	return
}
