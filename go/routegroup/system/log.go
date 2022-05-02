package system

import (
	"itflow/handle"
	"itflow/internal/log"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Log 日志路由
var Log *xmux.RouteGroup

func init() {
	Log = xmux.NewRouteGroup()

	Log.Post("/search/log", handle.SearchLog).Bind(&log.SearchLog{}).
		AddModule(midware.JsonToStruct)

	Log.Post("/log/classify", handle.LogClassify)

	// 	Log.Post("/log/list", handle.LogList).Bind(&log.SearchLog{}).
	// 		AddModule(midware.JsonToStruct).
	// 		ApiDescribe("第一次打开日志列表").
	// 		ApiReqStruct(&log.SearchLog{}).ApiRequestTemplate(`{"page": 1, "limit": 10}`).
	// 		ApiResStruct(&log.Loglist{}).ApiResponseTemplate(`{"loglist":[{"id":26,"exectime":1588840365,"classify":"login","content":"","ip":"127.0.0.1"},{"id":25,"exectime":1588840233,"classify":"login","content":"","ip":"127.0.0.1"},{"id":24,"exectime":1588837232,"classify":"login","content":"","ip":"127.0.0.1"},{"id":23,"exectime":1588837002,"classify":"login","content":"","ip":"127.0.0.1"},{"id":22,"exectime":1588833133,"classify":"login","content":"","ip":"127.0.0.1"},{"id":21,"exectime":1588833047,"classify":"login","content":"","ip":"127.0.0.1"}],"code":0, "count":100}`)
}
