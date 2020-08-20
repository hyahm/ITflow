package handle

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/bug"
	"itflow/internal/comment"
	"itflow/internal/project"
	"itflow/internal/response"
	"itflow/model"
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
	EnvList []cache.Env `json:"envlist"`
	Code    int         `json:"code"`
}

func GetEnv(w http.ResponseWriter, r *http.Request) {
	el := &envList{
		EnvList: make([]cache.Env, 0),
	}

	for _, v := range cache.CacheEidEnv {
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
			golog.Error(err)
			continue
		}
		ul.Users = append(ul.Users, *realname)
	}
	send, _ := json.Marshal(ul)
	w.Write(send)
	return

}

func GetProjectUser(w http.ResponseWriter, r *http.Request) {
	// 通过project 来获取所属用户
	projectname := r.FormValue("name")
	uid := xmux.GetData(r).Get("uid").(int64)
	// 先要判断下， 这个用户是否有这个项目的权限
	w.Write(project.GetUsersByProjectName(uid, cache.Project(projectname)))
	return

}

type versionList struct {
	VersionList []cache.Version `json:"versionlist"`
	Code        int             `json:"code"`
}

func GetVersion(w http.ResponseWriter, r *http.Request) {

	vl := &versionList{
		VersionList: make([]cache.Version, 0),
	}

	for _, v := range cache.CacheVidVersion {
		vl.VersionList = append(vl.VersionList, v)
	}
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

	l := len(cache.ShowBaseUrl)
	url := ""
	if cache.ShowBaseUrl[l-1:l] == "/" {
		url = cache.ShowBaseUrl + filename
	} else {
		url = cache.ShowBaseUrl + "/" + filename
	}

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
	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// golog.Info(string(body))
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
	ul := len(cache.ShowBaseUrl)
	if cache.ShowBaseUrl[ul-1:ul] == "/" {
		url.Url = cache.ShowBaseUrl + filename
	} else {
		url.Url = cache.ShowBaseUrl + "/" + filename
	}

	url.FileName = filename
	url.Uploaded = 1
	// uploadimg := "update user set headimg = ? where nickname=?"
	// nickname := xmux.GetData(r).Get("nickname").(string)
	// _, err = db.Mconn.Update(uploadimg, url.Url, nickname)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	s, _ := json.Marshal(url)

	w.Write(s)
	return

}

func BugShow(w http.ResponseWriter, r *http.Request) {
	// 判断是否有权限访问这个bug
	bid := r.FormValue("id")
	sl := &bug.RespShowBug{
		Comments: make([]*comment.Informations, 0),
	}
	err := db.Mconn.GetOne(`select b.id,b.content,i.name, l.name, e.name, b.title,p.name from bugs as b 
	join importants as i 
	join version as v 
	join level as l 
	join environment as e 
	join project as p 
	join user as u 
	join statusgroup as s 
	on b.id=1 
	and b.uid=u.id and i.id=b.iid and b.sid=s.id and b.vid=v.id and b.lid=l.id and b.eid=e.id and b.pid=p.id`).Scan(
		&sl.Id, &sl.Content, &sl.Important, &sl.Level,
	)

	err = model.NewInformationsByBid(bid, sl.Comments)
	if err != nil {
		sl.Code = 1
		sl.Msg = err.Error()
		w.Write(sl.Marshal())
		return
	}

	send, _ := json.Marshal(sl)
	w.Write(send)
	return
}
