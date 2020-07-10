package routegroup

import (
	"itflow/handle"
	"itflow/internal/log"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Log *xmux.GroupRoute

func init() {
	Log = xmux.NewGroupRoute().ApiCreateGroup("log", "日志查询", "日志")
	Log.ApiReqHeader("X-Token", "asdfasdfasdfasdfsdf")
	Log.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Log.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Log.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Log.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Log.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Log.Pattern("/search/log").Post(handle.SearchLog).Bind(&log.Search_log{}).
		AddMidware(midware.JsonToStruct).
		ApiDescribe("条件搜索日志列表").
		ApiReqStruct(&log.Search_log{}).ApiRequestTemplate(`{"page": 1, "limit": 10, "classify":"login"}`).
		ApiResStruct(&log.Loglist{}).
		ApiResponseTemplate(`{"loglist":[{"id":26,"exectime":1588840365,"classify":"login","content":"","ip":"127.0.0.1"},{"id":25,"exectime":1588840233,"classify":"login","content":"","ip":"127.0.0.1"},{"id":24,"exectime":1588837232,"classify":"login","content":"","ip":"127.0.0.1"},{"id":23,"exectime":1588837002,"classify":"login","content":"","ip":"127.0.0.1"},{"id":22,"exectime":1588833133,"classify":"login","content":"","ip":"127.0.0.1"},{"id":21,"exectime":1588833047,"classify":"login","content":"","ip":"127.0.0.1"}],"code":0, "count":100}`)

	Log.Pattern("/log/classify").Post(handle.LogClassify).
		ApiDescribe("获取classify列表")

	Log.Pattern("/log/list").Post(handle.LogList).Bind(&log.Search_log{}).
		AddMidware(midware.JsonToStruct).
		ApiDescribe("第一次打开日志列表").
		ApiReqStruct(&log.Search_log{}).ApiRequestTemplate(`{"page": 1, "limit": 10}`).
		ApiResStruct(&log.Loglist{}).ApiResponseTemplate(`{"loglist":[{"id":26,"exectime":1588840365,"classify":"login","content":"","ip":"127.0.0.1"},{"id":25,"exectime":1588840233,"classify":"login","content":"","ip":"127.0.0.1"},{"id":24,"exectime":1588837232,"classify":"login","content":"","ip":"127.0.0.1"},{"id":23,"exectime":1588837002,"classify":"login","content":"","ip":"127.0.0.1"},{"id":22,"exectime":1588833133,"classify":"login","content":"","ip":"127.0.0.1"},{"id":21,"exectime":1588833047,"classify":"login","content":"","ip":"127.0.0.1"}],"code":0, "count":100}`)
}
