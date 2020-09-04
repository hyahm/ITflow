package routegroup

import (
	"itflow/handle"
	"itflow/internal/version"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Version  bug版本
var Version *xmux.GroupRoute

func init() {
	// 所有路由同一格式， 第一行是pattern, method, bind
	// 第二行是中间件     如果没有中间件就下移
	// 第三行是api的主处理， ApiCreateGroup, ApiDelGroup
	// 请求头
	// 后面是api接口的次要处理
	// 最后是错误码
	Version = xmux.NewGroupRoute().ApiCreateGroup("version", "版本相关", "版本列表").AddModule(midware.VersionPermModule)
	Version.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Version.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Version.ApiCodeField("code").ApiCodeMsg("1", "其他错误,请查看返回的msg")
	Version.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")
	Version.Post("/version/add", handle.AddVersion).Bind(&version.RespVersion{}).AddModule(midware.JsonToStruct).
		ApiDescribe("添加版本")

	Version.Post("/version/list", handle.VersionList).ApiDescribe("显示版本")

	Version.Get("/version/remove", handle.VersionRemove).ApiDescribe("删除版本")

	Version.Post("/get/version", handle.GetVersion)

	Version.Post("/version/update", handle.VersionUpdate).Bind(&version.RespVersion{}).AddModule(midware.JsonToStruct).
		ApiDescribe("修改版本")
}
