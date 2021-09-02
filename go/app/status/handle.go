package status

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	res := &response.Response{}
	statuss, err := model.GetAllStatus()
	if err != nil {
		golog.Error(err)
		w.Write(res.ErrorE(err))
		return
	}
	res.Data = statuss
	w.Write(res.Marshal())
}

func Create(w http.ResponseWriter, r *http.Request) {

	status := xmux.GetInstance(r).Data.(*model.Status)
	err := status.Create()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID: status.ID,
	}
	w.Write(res.Marshal())
}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	err := model.DeleteStatus(id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func Update(w http.ResponseWriter, r *http.Request) {

	status := xmux.GetInstance(r).Data.(*model.Status)
	err := status.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	// 更新缓存
	w.Write(response.Success())
}

// func StatusGroupName(w http.ResponseWriter, r *http.Request) {

// 	sl := &network.List_StatusName{}
// 	// for _, v := range cache.CacheSgidGroup {
// 	// 	sl.StatusList = append(sl.StatusList, v)
// 	// }

// 	send, _ := json.Marshal(sl)
// 	w.Write(send)
// }
