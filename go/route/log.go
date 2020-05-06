package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Log *xmux.GroupRoute

func init() {
	Log = xmux.NewGroupRoute().AddMidware(midware.CheckLogPermssion)

	Log.Pattern("/search/log").Post(handle.SearchLog)

	Log.Pattern("/log/classify").Post(handle.LogClassify)

	Log.Pattern("/log/list").Post(handle.LogList)
}
