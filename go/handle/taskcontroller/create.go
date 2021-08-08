package taskcontroller

import (
	"itflow/model"
	"itflow/response"
	"net/http"
	"time"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

// 添加
func Create(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	bug := xmux.GetInstance(r).Data.(*model.Bug)
	// 判断是否有默认值
	dv, err := model.GetDefaultValue()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	if dv.Created == 0 {
		golog.Error("please set default status to create bug ")
		w.Write(response.Error("please set default status to create bug "))
		return
	}
	bug.Sid = dv.Created
	bug.UID = uid
	bug.CreateTime = time.Now().Unix()
	err = bug.CreateBug()
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
