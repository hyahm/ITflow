package midware

import (
	"itflow/network/datalog"
	"strings"
)

func EndLog(data interface{}) {
	if data == nil {
		return
	}
	dl := data.(*datalog.AddLog)

	if dl == nil {
		return
	}
	dl.Ip = strings.Split(dl.Ip, ":")[0]
	dl.Insert("nickname: %s has login", dl.Username)
}
