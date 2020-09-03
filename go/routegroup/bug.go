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
	Bug.Post("/bug/pass", handle.PassBug).Bind(&bug.PassBug{}).
		AddModule(midware.JsonToStruct).ApiDescribe("转交bug").ApiReqStruct(&bug.PassBug{})

	Bug.Post("/bug/create", handle.BugCreate).Bind(&bug.EditBug{}).
		AddModule(midware.JsonToStruct).
		ApiDescribe("创建或更新bug").
		ApiRequestTemplate(`{"title":"metu",
		"content":"<p>反反复复</p>",
		"id":-1,
		"selectuser":["sdfsadf"],
		"projectname":"123",
		"level":"2",
		"envname":"axi","important":"一般ee","version":"V 1.5"}`).
		ApiResponseTemplate(`{"id": 20, "code": 0, "message": "success"}`)

	Bug.Get("/bug/edit", handle.BugEdit).
		ApiDescribe("页面编辑获取数据").
		ApiResponseTemplate(`{"title":"metu","content":"<p>反反复复</p>","id":-1,"selectuser":["sdfsadf"],"projectname":"123","level":"2","envname":"axi","important":"一般ee","version":"V 1.5"}`)

	Bug.Post("/bug/mybugs", handle.GetMyBugs).Bind(&search.ReqMyBugFilter{})

	Bug.Get("/bug/close", handle.CloseBug)
	Bug.Post("/bug/changestatus", handle.ChangeBugStatus).Bind(&bug.ChangeStatus{}).
		AddModule(midware.JsonToStruct)

	Bug.Post("/status/filter", handle.ChangeFilterStatus).Bind(&status.Status{}).
		AddModule(midware.JsonToStruct).ApiDescribe("修改显示bug的状态")

	Bug.Get("/bug/show", handle.BugShow).
		ApiDescribe("获取此bug信息").ApiReqParams("id", "6").ApiResponseTemplate(`{
			"title": "sdfsdf",
			"content": "<p>sdfajsdjfsdfsaf</p>",
			"appversion": "v1.2",
			"comment": [],
			"status": "need133",
			"id": 2,
			"code": 0
		}`).ApiResStruct(bug.RespShowBug{})

	Bug.Get("/bug/resume", handle.ResumeBug)

	Bug.Post("/get/permstatus", handle.GetPermStatus)

	// Bug.Pattern("/get/thisrole",handle.GetThisRoles)
	Bug.Post("/get/group", handle.GetGroup)
}
