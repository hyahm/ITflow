package route

import (
	"itflow/app/handle"
	"itflow/midware"
	"itflow/model"

	"github.com/hyahm/xmux"
)

var Level *xmux.GroupRoute

func init() {
	Level = xmux.NewGroupRoute().AddMidware(midware.CheckLevelPermssion)

	Level.Pattern("/level/get").Post(handle.LevelGet)

	Level.Pattern("/level/add").Post(handle.LevelAdd).Bind(&model.Data_level{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Level.Pattern("/level/del").Get(handle.LevelDel).End(midware.EndLog)

	Level.Pattern("/level/update").Post(handle.LevelUpdate).Bind(&model.Update_level{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Level.Pattern("/get/levels").Post(handle.GetLevels)
}
