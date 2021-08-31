package handle

import (
	"itflow/db"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func RoleGroupList(w http.ResponseWriter, r *http.Request) {

	rgs, err := model.RoleGroupList()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := &response.Response{
		Data: rgs,
	}
	w.Write(res.Marshal())
}

func GetRoleGroupName(w http.ResponseWriter, r *http.Request) {

	kns, err := model.GetRoleKeyName()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	errorcode := &response.Response{
		Data: kns,
	}
	w.Write(errorcode.Marshal())
}

func RoleGroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	golog.Info(id)
	ssql := "select count(id) from jobs where rgid=?"
	var count int
	err := db.Mconn.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	if count > 0 {
		w.Write(errorcode.Error("有用户在使用， 无法删除"))
		return
	}
	// 先删除perm
	// 获取 permids
	rolegroup := model.RoleGroup{}
	err = rolegroup.GetRoleGroupById(id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	err = model.DeletePerms(rolegroup.PermIds...)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 然后删除rolegroup
	err = rolegroup.Delete()
	// perm 里面的也要删除
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	w.Write(response.Success())

}

func EditRoleGroup(w http.ResponseWriter, r *http.Request) {

	rr := xmux.GetInstance(r).Data.(*RequestRoleGroup)
	ids := make([]int64, 0, len(rr.PermIds))
	for _, v := range rr.PermIds {
		err := v.Update()
		if err != nil {
			golog.Error(err)
		}
		ids = append(ids, v.Id)
	}
	rolegroup := model.RoleGroup{
		ID:      rr.ID,
		Name:    rr.Name,
		PermIds: ids,
	}
	err := rolegroup.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

type RequestRoleGroup struct {
	ID      int64        `json:"id" db:"id,default"`
	Name    string       `json:"name" db:"name,default"`
	PermIds []model.Perm `json:"rolelist" db:"rolelist"`
}

func AddRoleGroup(w http.ResponseWriter, r *http.Request) {

	rr := xmux.GetInstance(r).Data.(*RequestRoleGroup)
	ids, err := model.InsertManyPerm(rr.PermIds)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	rolegroup := model.RoleGroup{
		ID:      rr.ID,
		Name:    rr.Name,
		PermIds: ids,
	}
	err = rolegroup.Insert()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	res := response.Response{
		ID: rolegroup.ID,
	}
	w.Write(res.Marshal())

}
