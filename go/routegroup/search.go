package routegroup

import (
	"itflow/handle"
	"itflow/internal/search"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Search 搜索相关
var Search *xmux.RouteGroup

func init() {
	Search = xmux.NewRouteGroup().AddModule(midware.CheckSetDefault)

	Search.Post("/search/allbugs", handle.SearchAllBugs).BindJson(&search.ReqMyBugFilter{})
	Search.Post("/search/mybugs", handle.SearchMyBugs).BindJson(&search.ReqMyBugFilter{})

	Search.Post("/search/mytasks", handle.SearchMyTasks).BindJson(&search.ReqMyBugFilter{})

	Search.Post("/search/bugmanager", handle.SearchBugManager).BindJson(&search.ReqMyBugFilter{})
}
