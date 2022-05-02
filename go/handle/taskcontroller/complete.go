package taskcontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Complete(w http.ResponseWriter, r *http.Request) {

	if model.Default.Completed <= 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "no permission"
		return
	}
	// 需要
	bug := xmux.GetInstance(r).Data.(*model.Bug)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	bug.Sid = model.Default.Completed
	bug.UpdateTime = time.Now().Unix()
	bug.Uid = uid

	// 判断是否有默认值
	err := bug.UpdateStatus()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	// res := response.Response{
	// 	UserIds: bug.Uids,
	// }

}
