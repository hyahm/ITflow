package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/network/env"

	"github.com/hyahm/xmux"
)

var Env *xmux.GroupRoute

func init() {
	Env = xmux.NewGroupRoute().AddMidware(midware.CheckEnvPermssion)

	Env.Pattern("/env/list").Post(handle.EnvList)

	Env.Pattern("/env/add").Get(handle.AddEnv).End(midware.EndLog)

	Env.Pattern("/env/update").Post(handle.UpdateEnv).Bind(&env.Env{}).
		AddMidware(midware.JsonToStruct).End(midware.EndLog)

	Env.Pattern("/env/delete").Get(handle.DeleteEnv).End(midware.EndLog)
}
