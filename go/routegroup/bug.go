package routegroup

import (
	"itflow/handle"
	"itflow/internal/bug"
	"itflow/internal/search"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Bug bug相关操作的路由组
var Bug *xmux.RouteGroup

func init() {
	Bug = xmux.NewRouteGroup()

	///  -------
	Bug.Post("/bug/pass", handle.PassBug).Bind(&handle.RequestPass{}).AddModule(midware.JsonToStruct)

	Bug.Post("/bug/mybugs", handle.GetMyBugs).Bind(&search.ReqMyBugFilter{})

	Bug.Get("/bug/close", handle.CloseBug)
	Bug.Post("/bug/changestatus", handle.ChangeBugStatus).Bind(&bug.ChangeStatus{}).
		AddModule(midware.JsonToStruct)

	Bug.Get("/bug/show", handle.BugShow)

	Bug.Get("/bug/resume", handle.ResumeBug)

	Bug.Post("/get/permstatus", handle.GetPermStatus)

	Bug.Get("/bug/delete", handle.DeleteBug)
	Bug.Post("/get/group", handle.GetGroup)
	Bug.Post("/get/task/typ", handle.GetTaskTyp)
}
