package handle

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type ResponseKeys struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Keys []model.Auth `json:"keys"`
}

func KeyList(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	keys, err := model.GetAllAuths(uid)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = keys
}

func AddKey(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	auth := xmux.GetInstance(r).Data.(*model.Auth)
	var err error
	var t int64
	if auth.ID > 0 {
		_, err = auth.Update(uid)
		t = auth.UpTime
	} else {
		auth.ID, err = auth.Insert(uid)
		t = auth.Created
	}
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).CreateTime = t
	xmux.GetInstance(r).Response.(*response.Response).ID = auth.ID
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 判断下这个key 是不是这个用户的
	err := model.CheckMyKey(uid, id)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	// 判断doc 里面是都有在使用
	err = model.CheckKid(id)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func CheckKeyName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if err := model.ChecKeyName(name); err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func GetMykeys(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	auths, err := model.GetKeyNamesByUid(uid)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = auths
}
