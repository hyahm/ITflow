package position

import (
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	jobs, err := model.GetAllPositions()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		Data: jobs,
	}
	w.Write(res.Marshal())
}

func PositionGet(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}

	// data := &model.Jobs{
	// 	Positions: make([]*model.Job, 0),
	// }

	// rows, err := db.Mconn.GetRows(`select j.id,j.name,level,hypo,IFNULL(s.name,''), IFNULL(r.name,'') from jobs as j
	// left join statusgroup as s  on j.bugsid = s.id
	// left join rolegroup as r on j.rid=r.id`)
	// if err != nil {
	// 	golog.Error(err)
	// 	w.Write(errorcode.ErrorE(err))
	// 	return
	// }
	// x := make(map[int64]string)
	// for rows.Next() {
	// 	one := &model.Job{}
	// 	rows.Scan(&one.Id, &one.Name, &one.Level, &one.Hid, &one.StatusGroup, &one.RoleGroup)
	// 	x[one.Id] = one.Name
	// 	data.Positions = append(data.Positions, one)
	// }
	// rows.Close()
	// for i := range data.Positions {
	// 	if data.Positions[i].Hid > 0 {
	// 		data.Positions[i].HypoName = x[data.Positions[i].Hid]
	// 	}
	// }
	// send, _ := json.Marshal(data)
	// w.Write(send)
	// return

}

func Create(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}

	job := xmux.GetInstance(r).Data.(*model.Job)
	err := job.Insert()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID: job.Id,
	}
	w.Write(res.Marshal())

}

func Delete(w http.ResponseWriter, r *http.Request) {

	// errorcode := &response.Response{}
	id := r.FormValue("id")
	uid := xmux.GetInstance(r).Get("uid").(int64)

	err := model.DeleteJob(id, uid)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}

	w.Write(response.Success())
}

func Update(w http.ResponseWriter, r *http.Request) {

	job := xmux.GetInstance(r).Data.(*model.Job)
	err := job.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())

}
