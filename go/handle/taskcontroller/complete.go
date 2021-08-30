package taskcontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"
	"time"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Complete(w http.ResponseWriter, r *http.Request) {
	// 需要
	bug := xmux.GetInstance(r).Data.(*model.Bug)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 判断权限
	if uid != bug.OwnerId && uid != goconfig.ReadInt64("adminid", 1) {
		w.Write(response.Error("no permission"))
		return
	}
	bug.Sid = model.Default.Completed
	bug.UpdateTime = time.Now().Unix()
	// 判断是否有默认值
	err := bug.Update()
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
