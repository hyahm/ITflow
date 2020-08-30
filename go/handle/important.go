package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/response"
	"itflow/model"
	"net/http"
	"strconv"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func ImportantGet(w http.ResponseWriter, r *http.Request) {

	data := &model.List_importants{
		ImportantList: make([]*model.Importants, 0),
	}

	rows, err := db.Mconn.GetRows("select id,name from importants")
	if err != nil {
		golog.Error(err)
		w.Write(data.ErrorE(err))
		return
	}
	for rows.Next() {
		im := &model.Importants{}
		err = rows.Scan(&im.Id, &im.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		data.ImportantList = append(data.ImportantList, im)
	}
	w.Write(data.Marshal())
	return

}

func ImportantAdd(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Data_importants)

	var err error
	errorcode.Id, err = db.Mconn.Insert("insert into importants(name) value(?)", data.Name)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志

	//更新缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ImportantDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	id32, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 判断是否有bug在使用
	var count int
	err = db.Mconn.GetOne("select count(id) from bugs where iid=?", id32).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.IsUse())
		return
	}
	// 是否设定为了默认值
	// if cache.CacheDefault["important"] == int64(id32) {
	// 	w.Write(errorcode.Error("没有设定默认值"))
	// 	return
	// }
	gsql := "delete from importants where id=?"

	_, err = db.Mconn.Update(gsql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志

	// 删除缓存
	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func ImportantUpdate(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	data := xmux.GetData(r).Data.(*model.Importants)
	gsql := "update importants set name=? where id=?"

	_, err := db.Mconn.Update(gsql, data.Name, data.Id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 增加日志

	// 删除strings key
	w.Write(errorcode.Success())
	return

}

type importantslist struct {
	Importants []string `json:"importants"`
	Code       int      `json:"code"`
	Msg        string   `json:"msg"`
}

func (im *importantslist) Marshal() []byte {
	send, err := json.Marshal(im)
	if err != nil {
		golog.Error(err)
	}
	return send
}
func (im *importantslist) Error(msg string) []byte {
	im.Code = 1
	im.Msg = msg
	return im.Marshal()
}
func (im *importantslist) ErrorE(err error) []byte {
	return im.Error(err.Error())
}

func GetImportants(w http.ResponseWriter, r *http.Request) {
	data := &importantslist{
		Importants: make([]string, 0),
	}
	rows, err := db.Mconn.GetRows("select name from importants")
	if err != nil {
		golog.Error(err)
		w.Write(data.ErrorE(err))
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			golog.Error(err)
			continue
		}
		data.Importants = append(data.Importants, name)
	}
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
