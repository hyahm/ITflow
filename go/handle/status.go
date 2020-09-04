package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/perm"
	"itflow/internal/response"
	"itflow/internal/status"
	network "itflow/model"
	"net/http"
	"strconv"

	"itflow/internal/bug"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func StatusList(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Select {
		w.Write(errorcode.Error("no perm"))
		return
	}
	w.Write(status.StatusList())
	return

}

func StatusAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Insert {
		w.Write(errorcode.Error("no perm"))
		return
	}
	var err error
	s := xmux.GetData(r).Data.(*bug.ReqStatus)
	errorcode.Id, err = db.Mconn.Insert("insert into status(name) values(?)", s.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ConnectMysqlFail())
		return
	}

	// 更新缓存
	w.Write(errorcode.Success())
	return

}

func StatusRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Delete {
		w.Write(errorcode.Error("no perm"))
		return
	}
	id := r.FormValue("id")
	sid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	if cache.DefaultCreateSid == sid {
		w.Write(errorcode.Error("this status is set to default, can not remove. you can change to other status and delete"))
		return

	}
	// 如果bug有这个状态，就不能修改
	var bcount int
	err = db.Mconn.GetOne("select count(id) from bugs where sid=?", sid).Scan(&bcount)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if bcount > 0 {
		golog.Errorf("sid:%d 删除失败", sid)
		w.Write(errorcode.IsUse())
		return
	}

	//如果状态组存在也无法删除

	// var hasgroup bool
	// for _, ids := range cache.CacheSgidGroup {
	// 	for _, v := range strings.Split(ids, ",") {
	// 		if v == id {
	// 			hasgroup = true
	// 			break
	// 		}
	// 	}
	// 	if hasgroup {
	// 		w.Write(errorcode.Error("还有group"))
	// 	}
	// }

	_, err = db.Mconn.Update("delete from  status where id=?", sid)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// 默认值

	// 更新缓存
	// 获取status的索引

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	perm := xmux.GetData(r).Get("perm").(perm.OptionPerm)
	if !perm.Update {
		w.Write(errorcode.Error("no perm"))
		return
	}
	s := xmux.GetData(r).Data.(*bug.ReqStatus)
	_, err := db.Mconn.Update("update status set name=? where id=?", s.Name, s.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 更新缓存

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return
}

func StatusGroupName(w http.ResponseWriter, r *http.Request) {

	sl := &network.List_StatusName{}
	// for _, v := range cache.CacheSgidGroup {
	// 	sl.StatusList = append(sl.StatusList, v)
	// }

	send, _ := json.Marshal(sl)
	w.Write(send)
	return

}
