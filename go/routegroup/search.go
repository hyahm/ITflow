package routegroup

import (
	"itflow/handle"
	"itflow/internal/search"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Search 搜索相关
var Search *xmux.GroupRoute

func init() {
	Search = xmux.NewGroupRoute().AddModule(midware.CheckSetDefault)

	Search.Post("/search/allbugs", handle.SearchAllBugs).Bind(&search.ReqMyBugFilter{}).AddModule(midware.JsonToStruct)
	Search.Post("/search/mybugs", handle.SearchMyBugs).Bind(&search.ReqMyBugFilter{}).AddModule(midware.JsonToStruct)

	Search.Post("/search/mytasks", handle.SearchMyTasks).Bind(&search.ReqMyBugFilter{}).
		AddModule(midware.JsonToStruct)

	Search.Post("/search/bugmanager", handle.SearchBugManager).Bind(&search.ReqMyBugFilter{}).
		AddModule(midware.JsonToStruct)
}
