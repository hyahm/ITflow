package handle

import (
	"encoding/json"
	"fmt"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

type ResponseKeys struct {
	Code int           `json:"code"`
	Msg  string        `json:"msg"`
	Keys []*model.Auth `json:"keys"`
}

func KeyList(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	uid := xmux.GetInstance(r).Get("uid").(int64)
	var err error
	rk := &ResponseKeys{}
	rk.Keys, err = model.GetAllAuths(uid)
	if err != nil {
		w.Write(errorcode.ErrorE(err))
		return
	}
	send, err := json.Marshal(rk)
	if err != nil {
		golog.Error(err)
	}
	w.Write(send)
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
		w.Write([]byte(fmt.Sprintf(`{"code": 1, "msg": "%v"}`, err)))
		return
	}

	w.Write([]byte(fmt.Sprintf(`{"code": 0, "time": %d, "id": %d}`, t, auth.ID)))
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	id := r.URL.Query().Get("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 判断下这个key 是不是这个用户的
	err := model.CheckMyKey(uid, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 判断doc 里面是都有在使用
	err = model.CheckKid(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	w.Write(errorcode.Success())
}

func CheckKeyName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if err := model.ChecKeyName(name); err != nil {
		send := fmt.Sprintf(`{"code": 0, "msg": "%v"}`, err)
		w.Write([]byte(send))
		return
	}
	w.Write([]byte(`{"code": 0}`))
}

func GetMykeys(w http.ResponseWriter, r *http.Request) {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	w.Write(model.GetKeyNamesByUid(uid))
}
