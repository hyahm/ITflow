package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/env"

	"github.com/hyahm/xmux"
)

var Env *xmux.GroupRoute

func init() {
	Env = xmux.NewGroupRoute().AddMidware(midware.CheckEnvPermssion).ApiCreateGroup("env", "环境相关", "env")
	Env.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Env.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Env.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Env.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")

	Env.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Env.Pattern("/env/list").Post(handle.EnvList)

	Env.Pattern("/env/add").Get(handle.AddEnv).End(midware.EndLog)

	Env.Pattern("/env/update").Post(handle.UpdateEnv).Bind(&env.Env{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	Env.Pattern("/env/delete").Get(handle.DeleteEnv).End(midware.EndLog)
}
