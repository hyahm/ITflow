package midware

import (
	"itflow/cache"
	"itflow/internal/response"
	"itflow/model"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

const (
	SELECT = 1
	REMOVE = 2
	UPDATE = 4
	CREATE = 8
)

func CheckStatusPermssion(w http.ResponseWriter, r *http.Request) bool {

	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "status")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckRoleGroupPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "rolegroup")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckProjectPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "project")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckLevelPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "level")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckEnvPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "env")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckLogPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "log")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckImportantPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "important")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckUserGroupPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "usergroup")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckVersionPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "version")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func CheckPositionPermssion(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetData(r).Get("uid").(int64)
	resp, ok := checkPermssion(uid, "position")
	if ok {
		w.Write(resp)
		return ok
	}
	return false
}

func checkPermssion(uid int64, pagename string) ([]byte, bool) {
	errorcode := &response.Response{}

	// 管理员
	if uid == cache.SUPERID {
		return nil, false
	}

	rg, err := model.NewRoleGroup(uid)
	if err != nil {
		golog.Error(err)

		return errorcode.ConnectMysqlFail(), true
	}
	permssion := rg.CheckPagePerm(pagename)

	if !permssion {
		golog.Error(string(errorcode.ErrorNoPermission()))
		return errorcode.ErrorNoPermission(), true
	}
	return nil, false
}
