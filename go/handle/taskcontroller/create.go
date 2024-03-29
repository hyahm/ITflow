package taskcontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

// 添加
func Create(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	bug := xmux.GetInstance(r).Data.(*model.Bug)

	if strings.Trim(bug.Title, "") == "" || strings.Trim(bug.Content, " ") == "" || len(bug.Uids) == 0 || bug.Pid == 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "缺少必要字段"
		return
	}
	if bug.Tid == 1 {
		if bug.Eid == 0 || bug.Vid == 0 || bug.Lid == 0 {
			xmux.GetInstance(r).Response.(*response.Response).Code = 1
			xmux.GetInstance(r).Response.(*response.Response).Msg = "缺少必要字段"
			return
		}
	}
	// 判断是否有默认值
	dv, err := model.GetDefaultValue()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	if dv.Created == 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "please set default status to create bug"
		return
	}
	bug.Sid = dv.Created
	bug.Uid = uid
	bug.CreateTime = time.Now().Unix()
	err = bug.CreateBug()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).ID = bug.ID

}

// 添加
func Update(w http.ResponseWriter, r *http.Request) {
	bug := xmux.GetInstance(r).Data.(*model.Bug)
	// 判断是否有默认值
	err := bug.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}
