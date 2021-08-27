package midware

import (
	"itflow/cache"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/xmux"
)

const (
	SELECT = 1
	REMOVE = 2
	UPDATE = 4
	CREATE = 8
)

// 状态码是2
func CheckSetDefault(w http.ResponseWriter, r *http.Request) bool {
	// 检查是否设置全了默认值， 否则无法打开bug任务管理菜单
	if model.Default.Completed == 0 ||
		model.Default.Created == 0 ||
		model.Default.Pass == 0 ||
		model.Default.Receive == 0 {
		w.Write([]byte(`{"code": 2, "msg": "必须先让管理员设置默认值"}`))
		return true
	}
	return false
}

func MustBeSuperAdmin(w http.ResponseWriter, r *http.Request) bool {
	uid := xmux.GetInstance(r).Get("uid").(int64)
	if uid != cache.SUPERID {
		w.Write(response.Error("no permission"))
		return true
	}
	return false
}
