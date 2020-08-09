package routegroup

import (
	"itflow/handle"
	"itflow/internal/bug"
	"itflow/internal/search"
	"itflow/internal/status"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Bug bug相关操作的路由组
var Bug *xmux.GroupRoute

func init() {
	Bug = xmux.NewGroupRoute().ApiCreateGroup("bug", "bug相关接口", "bug相关")
	Bug.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Bug.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Bug.ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg")
	Bug.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Bug.Pattern("/bug/pass").Post(handle.PassBug).End(midware.EndLog).Bind(&bug.PassBug{}).
		AddMidware(midware.JsonToStruct).ApiDescribe("转交bug").ApiReqStruct(&bug.PassBug{})

	Bug.Pattern("/bug/create").Post(handle.BugCreate).Bind(&bug.RespEditBug{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog).
		ApiDescribe("创建或更新bug").
		ApiRequestTemplate(`{"title":"metu","content":"<p>反反复复</p>","id":1,"selectuser":["sdfsadf"],"projectname":"123","level":"2","envname":"axi","important":"一般ee","version":"V 1.5"}`).
		ApiResponseTemplate(`{"id": 20, "code": 0, "message": "success"}`)

	Bug.Pattern("/bug/edit").Get(handle.BugEdit).
		End(midware.EndLog).
		ApiDescribe("页面编辑获取数据").
		ApiResponseTemplate(`{"title":"metu","content":"<p>反反复复</p>","id":-1,"selectuser":["sdfsadf"],"projectname":"123","level":"2","envname":"axi","important":"一般ee","version":"V 1.5"}`)

	Bug.Pattern("/bug/mybugs").Post(handle.GetMyBugs).Bind(&search.ReqMyBugFilter{}).
		End(midware.EndLog)

	Bug.Pattern("/bug/close").Get(handle.CloseBug).End(midware.EndLog)
	Bug.Pattern("/bug/changestatus").Post(handle.ChangeBugStatus).Bind(&bug.ChangeStatus{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	Bug.Pattern("/status/filter").Post(handle.ChangeFilterStatus).Bind(&status.Status{}).
		AddMidware(midware.JsonToStruct)

	Bug.Pattern("/bug/show").Get(handle.BugShow).
		ApiDescribe("获取此bug信息").ApiReqParams("id", "6").ApiResponseTemplate(`{
			"title": "sdfsdf",
			"content": "<p>sdfajsdjfsdfsaf</p>",
			"appversion": "v1.2",
			"comment": [],
			"status": "need133",
			"id": 2,
			"code": 0
		}`).ApiResStruct(bug.RespShowBug{})

	Bug.Pattern("/bug/resume").Get(handle.ResumeBug)

	Bug.Pattern("/get/permstatus").Post(handle.GetPermStatus)

	// Bug.Pattern("/get/thisrole").Get(handle.GetThisRoles)
	Bug.Pattern("/get/group").Post(handle.GetGroup)
}
