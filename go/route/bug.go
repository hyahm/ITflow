package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/bug"
	"itflow/network/status"

	"github.com/hyahm/xmux"
)

var Bug *xmux.GroupRoute

func init() {
	Bug = xmux.NewGroupRoute().ApiCodeField("code").ApiCodeMsg("0", "成功").ApiCreateGroup("bug", "bug相关接口", "bug相关")

	Bug.Pattern("/bug/pass").Post(handle.PassBug).End(midware.EndLog).Bind(&bug.PassBug{}).
		AddMidware(midware.JsonToStruct).ApiDescribe("转交bug").ApiReqStruct(&bug.PassBug{})

	Bug.Pattern("/bug/create").Post(handle.BugCreate).Bind(&bug.GetArticle{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("创建或更新bug")

	Bug.Pattern("/bug/edit").Get(handle.BugEdit).End(midware.EndLog)

	Bug.Pattern("/bug/mybugs").Post(handle.GetMyBugs).Bind(&bug.SearchParam{}).
		End(midware.EndLog)

	Bug.Pattern("/bug/close").Get(handle.CloseBug).End(midware.EndLog)
	Bug.Pattern("/bug/changestatus").Post(handle.ChangeBugStatus).Bind(&bug.ChangeStatus{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	Bug.Pattern("/status/filter").Post(handle.ChangeFilterStatus).Bind(&status.Status{}).
		AddMidware(midware.JsonToStruct)

	Bug.Pattern("/status/show").Post(handle.ShowStatus).Bind(&status.Status{}).
		AddMidware(midware.JsonToStruct)

	Bug.Pattern("/bug/show").Get(handle.BugShow).
		ApiDescribe("获取此bug信息").ApiReqParams(map[string]string{"id": "6"}).ApiResponseTemplate(`{
			"title": "sdfsdf",
			"content": "<p>sdfajsdjfsdfsaf</p>",
			"appversion": "v1.2",
			"comment": null,
			"status": "need133",
			"id": 2,
			"code": 0
		}`).ApiResStruct(bug.ShowBug{})

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
