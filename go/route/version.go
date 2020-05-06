package route

import (
	"itflow/app/handle"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

var Version *xmux.GroupRoute

func init() {
	// 所有路由同一格式， 第一行是pattern, method, bind
	// 第二行是中间件     如果没有中间件就下移
	// 第三行是api的主处理， ApiCreateGroup, ApiDelGroup
	// 请求头
	// 后面是api接口的次要处理
	// 最后是错误码
	Version = xmux.NewGroupRoute().AddMidware(midware.CheckVersionPermssion)

	Version.Pattern("/version/add").Post(handle.AddVersion).End(midware.EndLog)

	Version.Pattern("/version/list").Post(handle.VersionList)

	Version.Pattern("/version/remove").Get(handle.VersionRemove).End(midware.EndLog)

	Version.Pattern("/version/update").Post(handle.VersionUpdate).End(midware.EndLog)
}
