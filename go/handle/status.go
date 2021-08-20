package handle

import (
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/model"
	network "itflow/model"
	"itflow/response"
	"net/http"
	"strconv"

	"itflow/internal/bug"

	"github.com/go-sql-driver/mysql"
	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func StatusList(w http.ResponseWriter, r *http.Request) {
	res := &response.Response{}
	statuss, err := model.GetStatusList()
	if err != nil {
		golog.Error(err)
		w.Write(res.ErrorE(err))
		return
	}
	res.Data = statuss
	w.Write(res.Marshal())
}

func StatusAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	var err error
	s := xmux.GetInstance(r).Data.(*bug.ReqStatus)
	errorcode.ID, err = db.Mconn.Insert("insert into status(name) values(?)", s.Name)
	if err != nil {
		golog.Error(err)
		if err.(*mysql.MySQLError).Number == 1062 {
			w.Write(errorcode.ErrorE(db.DuplicateErr))
			return
		}
		w.Write(errorcode.ConnectMysqlFail())
		return
	}

	// 更新缓存
	w.Write(errorcode.Success())
}

func StatusRemove(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

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
}

func StatusUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	s := xmux.GetInstance(r).Data.(*bug.ReqStatus)
	_, err := db.Mconn.Update("update status set name=? where id=?", s.Name, s.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 更新缓存

	send, _ := json.Marshal(errorcode)
	w.Write(send)
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
