package routegroup

import (
	"itflow/handle"
	"itflow/internal/bug"
	"itflow/internal/search"
	"itflow/midware"
	"itflow/model"

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

	Bug.Post("/bug/mybugs", handle.GetMyBugs).Bind(&search.ReqMyBugFilter{})

	Bug.Get("/bug/close", handle.CloseBug)
	Bug.Post("/bug/changestatus", handle.ChangeBugStatus).Bind(&bug.ChangeStatus{}).
		AddModule(midware.JsonToStruct)

	Bug.Post("/status/filter", handle.ChangeFilterStatus).Bind(&model.User{}).
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

	Bug.Get("/bug/delete", handle.DeleteBug)
	Bug.Post("/get/group", handle.GetGroup)
	Bug.Post("/get/task/typ", handle.GetTaskTyp)
}
