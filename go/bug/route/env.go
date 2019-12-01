package route

import (
	"github.com/hyahm/xmux"
	"itflow/bug/handle"
)

var Env *xmux.GroupRoute

func init() {
	Env = xmux.NewGroupRoute()
	Env.Pattern("/env/list").Post(handle.EnvList)
	Env.Pattern("/env/add").Get(handle.AddEnv)
	Env.Pattern("/env/update").Post(handle.UpdateEnv)
	Env.Pattern("/env/delete").Get(handle.DeleteEnv)
}
