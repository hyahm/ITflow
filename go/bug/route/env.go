package route

import (
	"itflow/bug/handle"

	"github.com/hyahm/xmux"
)

var Env *xmux.GroupRoute

func init() {
	Env = xmux.NewGroupRoute("env")
	Env.Pattern("/env/list").Post(handle.EnvList)
	Env.Pattern("/env/add").Get(handle.AddEnv)
	Env.Pattern("/env/update").Post(handle.UpdateEnv)
	Env.Pattern("/env/delete").Get(handle.DeleteEnv)
}
