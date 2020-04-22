package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Bug *xmux.GroupRoute

func init() {
	Bug = xmux.NewGroupRoute("bug")

	Bug.Pattern("/bug/pass").Post(handle.PassBug)
	Bug.Pattern("/bug/create").Post(handle.BugCreate)
	Bug.Pattern("/bug/edit").Get(handle.BugEdit)
	Bug.Pattern("/bug/mybugs").Post(handle.GetMyBugs)
	Bug.Pattern("/bug/close").Get(handle.CloseBug)
	Bug.Pattern("/bug/changestatus").Post(handle.ChangeBugStatus)
	Bug.Pattern("/status/filter").Post(handle.ChangeFilterStatus)
	Bug.Pattern("/status/show").Post(handle.ShowStatus)
	Bug.Pattern("/bug/show").Get(handle.BugShow)
	Bug.Pattern("/search/allbugs").Post(handle.SearchAllBugs)
	Bug.Pattern("/search/mybugs").Post(handle.SearchMyBugs)
	Bug.Pattern("/search/mytasks").Post(handle.SearchMyTasks)
	Bug.Pattern("/search/bugmanager").Post(handle.SearchBugManager)

	Bug.Pattern("/get/user").Post(handle.GetUser)
	Bug.Pattern("/get/project").Post(handle.GetProject)
	Bug.Pattern("/get/version").Post(handle.GetVersion)
	Bug.Pattern("/get/env").Post(handle.GetEnv)
	Bug.Pattern("/get/status").Post(handle.GetStatus)

	Bug.Pattern("/get/permstatus").Post(handle.GetPermStatus)
	Bug.Pattern("/get/info").Post(handle.GetInfo)

	Bug.Pattern("/get/thisrole").Get(handle.GetThisRoles)
	Bug.Pattern("/get/group").Post(handle.GetGroup)
}
