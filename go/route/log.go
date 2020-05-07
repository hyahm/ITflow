package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/log"

	"github.com/hyahm/xmux"
)

var Log *xmux.GroupRoute

func init() {
	Log = xmux.NewGroupRoute().AddMidware(midware.CheckLogPermssion).ApiCreateGroup("log", "日志相关", "log")

	Log.Pattern("/search/log").Post(handle.SearchLog).Bind(&log.Search_log{}).
		AddMidware(midware.JsonToStruct).
		ApiReqStruct(&log.Search_log{})

	Log.Pattern("/log/classify").Post(handle.LogClassify)

	Log.Pattern("/log/list").Post(handle.LogList).Bind(&log.Search_log{}).
		AddMidware(midware.JsonToStruct).
		ApiDescribe("第一次打开日志列表").
		ApiReqStruct(&log.Search_log{}).ApiRequestTemplate(`{"page": 1, "limit": 10}`).
		ApiResStruct(&log.Loglist{}).ApiResponseTemplate(`{"loglist":[{"id":26,"exectime":1588840365,"classify":"login","content":"","ip":"127.0.0.1"},{"id":25,"exectime":1588840233,"classify":"login","content":"","ip":"127.0.0.1"},{"id":24,"exectime":1588837232,"classify":"login","content":"","ip":"127.0.0.1"},{"id":23,"exectime":1588837002,"classify":"login","content":"","ip":"127.0.0.1"},{"id":22,"exectime":1588833133,"classify":"login","content":"","ip":"127.0.0.1"},{"id":21,"exectime":1588833047,"classify":"login","content":"","ip":"127.0.0.1"}],"code":0, "count":100}`)
}
