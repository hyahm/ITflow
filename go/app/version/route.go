package version

import (
	"itflow/model"

	"github.com/hyahm/xmux"
)

// Version  bug版本
var Version *xmux.RouteGroup

func init() {
	// 所有路由同一格式， 第一行是pattern, method, bind
	// 第二行是中间件     如果没有中间件就下移
	// 第三行是api的主处理， ApiCreateGroup, ApiDelGroup
	// 请求头
	// 后面是api接口的次要处理
	// 最后是错误码
	Version = xmux.NewRouteGroup().AddPageKeys("version")
	Version.Post("/version/add", Create).BindJson(&model.Version{})

	Version.Post("/version/list", Read)

	Version.Get("/version/remove", Delete)

	Version.Post("/get/version", GetVersion)

	Version.Post("/version/update", Update).BindJson(&model.Version{})
}
