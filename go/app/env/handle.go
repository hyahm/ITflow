package env

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {

	envs, err := model.GetAllEnv()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: envs,
	}
	w.Write(res.Marshal())
}

func Create(w http.ResponseWriter, r *http.Request) {

	env := xmux.GetInstance(r).Data.(*model.Env)
	err := env.Create()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID: env.ID,
	}
	// 添加缓存
	w.Write(res.Marshal())

}

func Update(w http.ResponseWriter, r *http.Request) {

	env := xmux.GetInstance(r).Data.(*model.Env)

	err := env.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}

func Delete(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	err := model.DeleteEnv(id)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}
