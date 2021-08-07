package setting

import (
	"itflow/handle"
	"itflow/internal/env"
	"itflow/midware"

	"github.com/hyahm/xmux"
)

// Env 环境路由
var Env *xmux.GroupRoute

func init() {
	Env = xmux.NewGroupRoute().ApiCreateGroup("env", "环境相关", "env").AddModule(midware.EnvPermModule)
	Env.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Env.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Env.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Env.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")

	Env.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Env.Post("/env/list", handle.EnvList)

	Env.Get("/env/add", handle.AddEnv)

	Env.Post("/env/update", handle.UpdateEnv).Bind(&env.Env{}).
		AddModule(midware.JsonToStruct)
	Env.Post("/get/env", handle.GetEnv)
	Env.Get("/env/delete", handle.DeleteEnv)
}
