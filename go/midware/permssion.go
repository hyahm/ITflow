package midware

import (
	"itflow/app/bugconfig"
	"itflow/model"
	"itflow/network/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func CheckStatusPermssion(w http.ResponseWriter, r *http.Request) bool {
	var permssion bool
	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)

	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		rg, err := model.NewRoleGroup(nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ConnectMysqlFail())
			return true
		}
		permssion = rg.CheckPagePerm("status")

	}

	if !permssion {
		golog.Error(string(errorcode.ErrorNoPermission()))
		w.Write(errorcode.ErrorNoPermission())
		return true
	}
	return false
}
