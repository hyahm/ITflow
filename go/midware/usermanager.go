package midware

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func JobAuth(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 根据uid 获取 job_id
	jid, err := model.GetJobIdByUid(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return true
	}
	// jobs: 能管理的这些职位
	jobs, err := model.GetJobIdsByJobId(jid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return true
	}
	//
	xmux.GetInstance(r).Set("jobs", jobs)
	return false
}
