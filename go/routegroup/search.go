package routegroup

import (
	"itflow/handle"
	"itflow/internal/search"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Search *xmux.GroupRoute

func init() {
	Search = xmux.NewGroupRoute().ApiCreateGroup("search", "搜索相关接口", "搜索")
	Bug.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Bug.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Bug.ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg")
	Bug.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Search.Pattern("/search/allbugs").Post(handle.SearchAllBugs).Bind(&search.ReqMyBugFilter{}).AddMidware(midware.JsonToStruct)
	Search.Pattern("/search/mybugs").Post(handle.SearchMyBugs).Bind(&search.ReqMyBugFilter{}).AddMidware(midware.JsonToStruct)

	Search.Pattern("/search/mytasks").Post(handle.SearchMyTasks).Bind(&search.ReqMyBugFilter{}).
		AddMidware(midware.JsonToStruct)

	Search.Pattern("/search/bugmanager").Post(handle.SearchBugManager).Bind(&search.ReqMyBugFilter{}).AddMidware(midware.JsonToStruct)
}
