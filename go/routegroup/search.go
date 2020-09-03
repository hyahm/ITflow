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
	Search = xmux.NewGroupRoute().ApiCreateGroup("search", "搜索相关接口", "搜索")
	Search.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Search.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Search.ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg")
	Search.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Search.Post("/search/allbugs", handle.SearchAllBugs).Bind(&search.ReqMyBugFilter{}).AddModule(midware.JsonToStruct).
		ApiDescribe("所有所有bug")
	Search.Post("/search/mybugs", handle.SearchMyBugs).Bind(&search.ReqMyBugFilter{}).AddModule(midware.JsonToStruct).
		ApiDescribe("所有我创建的bug")

	Search.Post("/search/mytasks", handle.SearchMyTasks).Bind(&search.ReqMyBugFilter{}).
		AddModule(midware.JsonToStruct).ApiDescribe("所有我的任务")

	Search.Post("/search/bugmanager", handle.SearchBugManager).Bind(&search.ReqMyBugFilter{}).
		AddModule(midware.JsonToStruct).ApiDescribe("搜索垃圾箱的bug")
}
