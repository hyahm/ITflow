package taskcontroller

import (
	"fmt"
	"itflow/model"
	"itflow/response"
	"net/http"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

// 添加
func Receive(w http.ResponseWriter, r *http.Request) {
	// 需要

	bug := xmux.GetInstance(r).Data.(*model.Bug)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 判断是否有默认值
	if model.Default.Receive <= 0 {
		w.Write([]byte("not set default status"))
		return
	}
	bug.Sid = model.Default.Receive
	err := bug.UpdateStatus(model.Default.Created, model.Default.Pass)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	bug.Uids = []int64{uid}
	// 还要增加information
	information := model.Information{
		Uid:  uid,
		Bid:  bug.ID,
		Info: fmt.Sprintf("uid: %d 领取了任务， 完成时间为： %s", uid, time.Unix(bug.DeadLine, 0).Format("2006-01-02 15:04:05")),
		Time: time.Now().Unix(),
	}
	err = information.Insert()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}

	res := response.Response{
		UserIds: bug.Uids,
	}
	w.Write(res.Marshal())

}
