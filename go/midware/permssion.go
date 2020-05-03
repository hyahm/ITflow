package midware

import (
	"itflow/app/asset"
	"itflow/app/bugconfig"
	"itflow/model/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func CheckPermssion(w http.ResponseWriter, r *http.Request) bool {
	var permssion bool
	var err error
	errorcode := &response.Response{}
	nickname := xmux.GetData(r).Get("nickname").(string)
	// 管理员
	if bugconfig.CacheNickNameUid[nickname] == bugconfig.SUPERID {
		permssion = true
	} else {
		permssion, err = asset.CheckPerm("status", nickname)
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ConnectMysqlFail())
			return true
		}
	}

	if !permssion {
		golog.Error(string(errorcode.ErrorNoPermission()))
		w.Write(errorcode.ErrorNoPermission())
		return true
	}
	return false
}
