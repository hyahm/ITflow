package version

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	version := xmux.GetInstance(r).Data.(*model.Version)
	version.CreateUid = xmux.GetInstance(r).Get("uid").(int64)
	err := version.Create()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID:         version.Id,
		CreateTime: version.CreateTime,
	}
	w.Write(res.Marshal())

}

func Read(w http.ResponseWriter, r *http.Request) {
	vs, err := model.GetAllVersion()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: vs,
	}
	w.Write(res.Marshal())

}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	err := model.DeleteVersion(id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}

func Update(w http.ResponseWriter, r *http.Request) {

	version := xmux.GetInstance(r).Data.(*model.Version)
	err := version.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func GetVersion(w http.ResponseWriter, r *http.Request) {

	// for _, v := range cache.CacheVidVersion {
	// 	vl.VersionList = append(vl.VersionList, v)
	// }
	// send, _ := json.Marshal(vl)
	// w.Write(send)
	w.Write(response.Success())
}
