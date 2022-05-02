package taskcontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

// 通过id 过去 bug
func Get(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)
	bug, err := model.GetBugById(id, uid)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = bug

}
