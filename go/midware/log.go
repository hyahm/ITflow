package midware

import (
	"itflow/internal/datalog"
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
	dl.Insert()
}
