package publiccontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ShowStatus(w http.ResponseWriter, r *http.Request) {
	// 获取显示的状态id
	// sl := xmux.GetInstance(r).Data.(*status.Status)

	uid := xmux.GetInstance(r).Get("uid").(int64)
	us := model.UserStatus{}
	ss, err := us.GetShowStatus(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = ss
}
