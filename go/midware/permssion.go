package midware

import (
	"itflow/cache"
	"itflow/internal/response"
	"itflow/model"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func CheckStatusPermssion(w http.ResponseWriter, r *http.Request) bool {

	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "status")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckRoleGroupPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "rolegroup")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckProjectPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "project")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckLevelPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "level")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckEnvPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "env")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckLogPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "log")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckImportantPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "important")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckUserGroupPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "usergroup")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckVersionPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "version")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckPositionPermssion(w http.ResponseWriter, r *http.Request) bool {
	nickname := xmux.GetData(r).Get("nickname").(string)
	resp, ok := checkPermssion(nickname, "position")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func checkPermssion(nickname, pagename string) ([]byte, bool) {
	var permssion bool
	errorcode := &response.Response{}

	// 管理员
	if cache.CacheNickNameUid[nickname] == cache.SUPERID {
		permssion = true
	} else {
		rg, err := model.NewRoleGroup(nickname)
		if err != nil {
			golog.Error(err)

			return errorcode.ConnectMysqlFail(), true
		}
		permssion = rg.CheckPagePerm("status")

	}

	if !permssion {
		golog.Error(string(errorcode.ErrorNoPermission()))
		return errorcode.ErrorNoPermission(), true
	}
	return nil, false
}
