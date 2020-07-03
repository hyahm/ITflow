package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/internal/bug"
	"itflow/internal/datalog"
	"itflow/internal/response"
	"net/http"
	"time"

	"github.com/hyahm/golog"

	"github.com/hyahm/xmux"
	//"strings"
)

type projectList struct {
	ProjectList []string `json:"projectlist"`
	Code        int      `json:"code"`
}

func GetProject(w http.ResponseWriter, r *http.Request) {

	pl := &projectList{}

	for _, v := range cache.CachePidName {
		pl.ProjectList = append(pl.ProjectList, v)
	}
	send, _ := json.Marshal(pl)
	w.Write(send)
	return

}

// 添加或编辑
func BugCreate(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	data := xmux.GetData(r).Data.(*bug.RespEditBug)
	statusId := cache.DefaultSid
	if statusId == 0 {
		w.Write([]byte("必须给定一个状态默认值"))
		return
	}

	bug, err := data.ToBug()
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	bug.StatusId = statusId
	bug.Uid = xmux.GetData(r).Get("uid").(int64)
	//
	golog.Info("777777")
	xmux.GetData(r).End = &datalog.AddLog{
		Ip:       r.RemoteAddr,
		Username: nickname,
		Classify: "bug",
		Action:   "create",
	}
	if data.Id <= 0 {
		// 插入bug
		bug.CreateTime = time.Now().Unix()
		err = bug.CreateBug()
		if err != nil {

			w.Write(errorcode.ErrorE(err))
			return
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
		xmux.GetData(r).End.(*datalog.AddLog).Action = "update"

	}

	w.Write(errorcode.Success())
	return

}
