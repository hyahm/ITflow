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
	ID          int     `json:"id"`
	Importance  string  `json:"importance"`
	Status      string  `json:"status"`
	Title       string  `json:"title"`
	Spusers     []int64 `json:"spusers"`
	Selectoses  string  `json:"selectoses"`
	AppVersion  string  `json:"appversion"`
	Content     string  `json:"content"`
	Level       string  `json:"level"`
	Projectname string  `json:"projectname"`
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
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
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

	xmux.GetInstance(r).Response.(*response.Response).Data = ul
}

func IsAdmin(w http.ResponseWriter, r *http.Request) {
	xmux.GetInstance(r).Response.(*response.Response).IsAdmin =
		xmux.GetInstance(r).Get("uid").(int64) == cache.SUPERID
}

func GetProjectUser(w http.ResponseWriter, r *http.Request) {
	// 通过project 来获取所属用户
	pid := xmux.Var(r)["id"]
	projectId, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	res := response.Response{}
	// 通过project_id 获取对应的version_ids
	res.VersionIds, err = model.GetVersionIdsByProjectId(projectId)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	// 通过project_id 获取对应的user_id
	// 获取usergroupid
	ug := model.UserGroup{}
	ug.Id, err = model.GetUserGroupId(projectId)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	err = ug.GetUserIds()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).UserIds = ug.Uids
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

	file, h, err := r.FormFile("image")
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
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
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	defer cfile.Close()

	_, err = io.Copy(cfile, file)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	url := cache.ShowBaseUrl + filename

	sendurl := &uploadImage{
		HasSuccess: true,
		Url:        url,
		Uid:        time.Now().UnixNano(),
	}

	xmux.GetInstance(r).Response.(*response.Response).Data = sendurl

}

type uploadimage struct {
	Uploaded int    `json:"uploaded"`
	Url      string `json:"url"`
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
}

func UploadHeadImg(w http.ResponseWriter, r *http.Request) {
	url := &uploadimage{}
	image, header, err := r.FormFile("upload")
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	imgcode := make([]byte, header.Size)
	_, err = image.Read(imgcode)
	if err != nil {
		golog.Errorf("parse uploadImage struct fail,%v", err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	prefix := strconv.FormatInt(time.Now().UnixNano(), 10)
	filename := prefix + ".png"
	err = ioutil.WriteFile(path.Join(cache.ImgDir, filename), imgcode, 0655) //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	url.Url = cache.ShowBaseUrl + filename

	url.FileName = filename
	url.Uploaded = 1
	uploadimg := "update user set headimg = ? where id=?"
	uid := xmux.GetInstance(r).Get("uid").(int64)
	result := db.Mconn.Update(uploadimg, url.Url, uid)
	if result.Err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = url
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
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
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
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
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
	xmux.GetInstance(r).Response.(*response.Response).Data = sl
}
