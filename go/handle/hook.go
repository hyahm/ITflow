package handle

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Gitee(w http.ResponseWriter, r *http.Request) {
	hook := xmux.GetInstance(r).Get("hook").(string)
	if hook != r.Header.Get("X-Gitee-Token") {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	did := xmux.GetInstance(r).Get("did").(int64)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	doc, err := model.NewDocById(did, uid)
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	_git, err := doc.NewGit()
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	err = _git.GitPull()
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	w.WriteHeader(200)
	return
}

func Gitlab(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Gitlab-Token")
	hook := xmux.GetInstance(r).Get("hook").(string)

	if token != hook {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	did := xmux.GetInstance(r).Get("did").(int64)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	doc, err := model.NewDocById(did, uid)
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	_git, err := doc.NewGit()
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	err = _git.GitPull()
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	w.WriteHeader(200)
	return
}

func Github(w http.ResponseWriter, r *http.Request) {
	hook := xmux.GetInstance(r).Get("hook").(string)
	x, err := ioutil.ReadAll(r.Body)
	if err != nil {
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		return
	}
	s := hmac.New(sha1.New, []byte(hook))
	s.Write(x)
	token := fmt.Sprintf("%x", s.Sum(nil))
	golog.Info(token)
	// 判断token 是否相等
	if "sha1="+token != r.Header.Get("X-Hub-Signature") {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	did := xmux.GetInstance(r).Get("did").(int64)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	doc, err := model.NewDocById(did, uid)
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	_git, err := doc.NewGit()
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	err = _git.GitPull()
	if err != nil {
		w.WriteHeader(http.StatusNetworkAuthenticationRequired)
		return
	}
	w.WriteHeader(200)
	return
}

func RandomHook(w http.ResponseWriter, r *http.Request) {
	// 	uid := xmux.GetInstance(r).Get("uid").(int64)
	// 	errorcode := &response.Response{}
	// 	rand := strconv.FormatInt(time.Now().UnixNano(), 10)
	// 	s1 := sha1.New()
	// 	io.WriteString(s1, rand)
	// 	hook := fmt.Sprintf("%x", s1.Sum(nil))
	// 	fmt.Println(hook)
	// 	err := model.SetHookById(hook, uid)
	// 	if err != nil {
	// 		golog.Error(err)
	// 		w.Write(errorcode.Error(err))
	// 		return
	// 	}
	// 	w.Write([]byte(fmt.Sprintf(`{"code" : 0, "hook": "%s"}`, hook)))
	// 	return
}
