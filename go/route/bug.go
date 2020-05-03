package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model/bug"

	"github.com/hyahm/xmux"
)

var Bug *xmux.GroupRoute

func init() {
	Bug = xmux.NewGroupRoute().ApiCodeField("code").ApiCodeMsg("0", "成功")

	Bug.Pattern("/bug/pass").Post(handle.PassBug).End(midware.EndLog).End(midware.EndLog)
	Bug.Pattern("/bug/create").Post(handle.BugCreate).End(midware.EndLog).End(midware.EndLog)
	Bug.Pattern("/bug/edit").Get(handle.BugEdit).End(midware.EndLog).End(midware.EndLog)

	Bug.Pattern("/bug/mybugs").Post(handle.GetMyBugs).Bind(&bug.SearchParam{}).
		End(midware.EndLog)

	Bug.Pattern("/bug/close").Get(handle.CloseBug).End(midware.EndLog).End(midware.EndLog)
	Bug.Pattern("/bug/changestatus").Post(handle.ChangeBugStatus).End(midware.EndLog).End(midware.EndLog)
	Bug.Pattern("/status/filter").Post(handle.ChangeFilterStatus)
	Bug.Pattern("/status/show").Post(handle.ShowStatus)
	Bug.Pattern("/bug/show").Get(handle.BugShow)
	Bug.Pattern("/search/allbugs").Post(handle.SearchAllBugs)
	Bug.Pattern("/search/mybugs").Post(handle.SearchMyBugs)

	Bug.Pattern("/search/mytasks").Post(handle.SearchMyTasks).Bind(&bug.SearchParam{}).
		AddMidware(midware.JsonToStruct)

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
