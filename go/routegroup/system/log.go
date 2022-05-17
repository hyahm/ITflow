package system

import (
	"itflow/handle"
	"itflow/internal/log"

	"github.com/hyahm/xmux"
)

// Log 日志路由
var Log *xmux.RouteGroup

func init() {
	Log = xmux.NewRouteGroup()

	Log.Post("/search/log", handle.SearchLog).BindJson(&log.SearchLog{})

	Log.Post("/log/classify", handle.LogClassify)

	// 	Log.Post("/log/list", handle.LogList).BindJson(&log.SearchLog{})
}
