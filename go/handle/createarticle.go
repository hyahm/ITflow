package handle

import (
	"encoding/json"
	"fmt"
	"itflow/cache"
	"itflow/internal/bug"
	"itflow/internal/datalog"
	"itflow/internal/project"
	"itflow/internal/response"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/hyahm/xmux"
	//"strings"
)

type projectList struct {
	ProjectList []string `json:"projectlist"`
	Code        int      `json:"code"`
}

func GetProject(w http.ResponseWriter, r *http.Request) {

	pl := &projectList{}

	for _, v := range cache.CachePidProject {
		pl.ProjectList = append(pl.ProjectList, v.ToString())
	}
	send, _ := json.Marshal(pl)
	w.Write(send)
	return

}

func GetMyProject(w http.ResponseWriter, r *http.Request) {
	myproject := &project.MyProject{
		Name: make([]string, 0),
	}
	uid := xmux.GetData(r).Get("uid").(int64)

	w.Write(myproject.Get(uid))
	return

}

// 添加或编辑
func BugCreate(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	data := xmux.GetData(r).Data.(*bug.RespEditBug)
	createdId := cache.DefaultCreateSid
	if createdId == 0 {
		w.Write([]byte("必须给定一个状态默认值"))
		return
	}

	bug, err := data.ToBug()
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	bug.StatusId = createdId
	bug.Uid = xmux.GetData(r).Get("uid").(int64)
	//
	go datalog.InsertLog("bug", nickname+"create bug: "+data.Title, r.RemoteAddr, nickname, "create")

	if data.Id <= 0 {
		// 插入bug
		bug.CreateTime = time.Now().Unix()

		err = bug.CreateBug()
		if err != nil {
			w.Write(errorcode.ErrorE(err))
			return
		}
		if cache.CacheEmail.Enable {
			emails := make([]string, 0)
			for _, v := range strings.Split(string(bug.OprateUsers), ",") {
				thisUid, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					continue
				}
				emails = append(emails, cache.CacheUidEmail[thisUid])
			}
			cache.CacheEmail.SendMail("创建bug", fmt.Sprintf("由%s创建bug： title： %s", cache.CacheUidRealName[bug.Uid], bug.Title), emails...)
		}

		errorcode.Id = bug.ID

	} else {
		// update
		errorcode.Id = data.Id
		bug.UpdateTime = time.Now().Unix()
		err = bug.EditBug()
		if err != nil {
			w.Write(errorcode.ErrorE(err))
			return
		}
		go datalog.InsertLog("bug", nickname+fmt.Sprintf(" update bug id: %d", data.Id), r.RemoteAddr, nickname, "update")

	}

	w.Write(errorcode.Success())
	return

}
