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
	// w.Write(bug.BugById(id, uid))
	bug, err := model.GetBugById(id, uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: bug,
	}
	w.Write(res.Marshal())

}
