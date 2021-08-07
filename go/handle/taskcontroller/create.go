package taskcontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

// 添加
func Create(w http.ResponseWriter, r *http.Request) {
	bug := xmux.GetInstance(r).Data.(*model.Bug)
	err := bug.CreateBug()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	// 创建的id 返回
	res := response.Response{
		ID: bug.ID,
	}
	w.Write(res.Marshal())

}
