package midware

import (
	"encoding/json"
	"io/ioutil"
	"itflow/cache"
	"itflow/jwt"
	"itflow/model"
	"itflow/response"
	"net/http"
	"strings"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func JsonToStruct(w http.ResponseWriter, r *http.Request) bool {
	resp := &response.Response{}
	if goconfig.ReadBool("debug", false) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			golog.Error(err)
			w.Write(resp.ErrorE(err))
			return true
		}
		golog.Info(string(b))
		err = json.Unmarshal(b, xmux.GetInstance(r).Data)
		if err != nil {
			golog.Error(err)
			w.Write(resp.ErrorE(err))
			return true
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(xmux.GetInstance(r).Data)
		if err != nil {
			golog.Error(err)
			w.Write(resp.ErrorE(err))
			return true
		}

	}
	return false
}

func CheckToken(w http.ResponseWriter, r *http.Request) bool {
	errorcode := &response.Response{}
	a := r.Header.Get("Authorization")
	if a == "" {
		golog.Error("not found token")
		w.Write(errorcode.TokenNotFound())
		return true
	}
	token := &jwt.Token{}
	if !token.CheckJwt(strings.Split(a, " ")[1]) {
		w.Write(errorcode.TokenNotFound())
		return true
	}

	// 检查权限
	xmux.GetInstance(r).Set("nickname", token.Nickname)
	xmux.GetInstance(r).Set("uid", token.Id)

	return false
}

func CheckRole(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	if uid == cache.SUPERID {
		// 超级管理员有任何权限
		return false
	}
	//
	pages := xmux.GetInstance(r).Get(xmux.PAGES).(map[string]struct{})

	// 如果长度为空，就是有页面权限
	if len(pages) == 0 {
		return false
	}
	//  请求/project/read     map[admin:{} project:{}]
	// 判断 pages 是否存在 perm
	// 注意点： 这里的页面权限本应该只会匹配到一个， 这个是对于的页面权限的值
	var pl = []string{"Read", "Create", "Update", "Delete"}
	permissionMap := make(map[string]int, len(pl))
	for i, v := range pl {
		permissionMap[v] = i
	}
	// 根据uid 获取 permids
	permids, err := model.GetPermIdsByUid(uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return true
	}

	perm, err := model.GetPermsionPageAndPVById(permids)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return true
	}
	page := ""
	// 判断页面权限的
	hasPerm := false
	for role := range perm {
		if _, ok := pages[role]; ok {
			hasPerm = true
			page = role
			break
		}
	}
	if !hasPerm {
		w.Write(response.Error("没有页面权限"))
		return true
	}

	// permMap := make(map[string]bool)
	result := xmux.GetPerm(pl, perm[page])
	handleName := xmux.GetInstance(r).Get(xmux.CURRFUNCNAME).(string)
	// 这个值就是判断有没有这个操作权限
	if !result[permissionMap[handleName]] {
		w.Write(response.Error("没有权限"))
		return true
	}
	return false
}
