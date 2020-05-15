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
	Level.ApiCodeField("code").ApiCodeMsg("0", "成功")
	Level.ApiCodeField("code").ApiCodeMsg("20", "token过期")
	Level.ApiCodeField("code").ApiCodeMsg("2", "系统错误")
	Level.ApiCodeField("code").ApiCodeMsg("", "其他错误,请查看返回的msg")
	Level.ApiReqHeader("X-Token", "xxxxxxxxxxxxxxxxxxxxxxxxxx")

	Level.Pattern("/level/get").Post(handle.LevelGet)

	Level.Pattern("/level/add").Post(handle.LevelAdd).Bind(&model.Data_level{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Level.Pattern("/level/del").Get(handle.LevelDel).End(midware.EndLog)

	Level.Pattern("/level/update").Post(handle.LevelUpdate).Bind(&model.Update_level{}).AddMidware(midware.JsonToStruct).
		End(midware.EndLog)

	Level.Pattern("/get/levels").Post(handle.GetLevels)
}
