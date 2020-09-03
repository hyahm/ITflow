package midware

import (
	"fmt"
	"itflow/internal/perm"
	"itflow/internal/response"
	"net/http"

	"github.com/hyahm/xmux"
)

const (
	SELECT = 1
	REMOVE = 2
	UPDATE = 4
	CREATE = 8
)

type UserChecker interface {
	CheckUser(uid int64) error
}

func CheckUser(w http.ResponseWriter, r *http.Request) bool {

	uid := xmux.GetData(r).Get("uid").(int64)
	resp := response.Response{}
	if xmux.GetData(r).Data == nil {
		err := fmt.Sprintf("must be bind data first %s", r.URL.RequestURI())
		w.Write(resp.Error(err))
		return true
	}
	err := xmux.GetData(r).Data.(UserChecker).CheckUser(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	return false
}

func EnvPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.EnvPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func ImportantPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.ImportantPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func LevelPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.LevelPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func PositionPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.PositionPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func ProjectPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.ProjectPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func StatusPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.StatusPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func StatusgroupPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.StatusgroupPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}

func VersionPermModule(w http.ResponseWriter, r *http.Request) bool {
	resp := response.Response{}
	uid := xmux.GetData(r).Get("uid").(int64)
	op, err := perm.VersionPerm(uid)
	if err != nil {
		w.Write(resp.ErrorE(err))
		return true
	}
	xmux.GetData(r).Set("perm", op)
	return false
}
