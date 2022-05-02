package handle

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func UserKeyName(w http.ResponseWriter, r *http.Request) {
	// 获取用户keyvalue
	uid := xmux.GetInstance(r).Get("uid").(int64)
	kns, err := model.GetUserKeyName(uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = kns
}
