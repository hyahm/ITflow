package midware

import (
	"itflow/model/datalog"
	"strings"
)

func EndLog(data interface{}) {
	dl := data.(*datalog.AddLog)

	if dl == nil {
		return
	}
	dl.Ip = strings.Split(dl.Ip, ":")[0]
	dl.Insert("nickname: %s has login", dl.Username)
}
