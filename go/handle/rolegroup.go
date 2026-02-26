package handle

import (
	"itflow/app/service"
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
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = rgs
}

func GetRoleGroupName(w http.ResponseWriter, r *http.Request) {

	// kns, err := model.GetRoleKeyName()
	// if err != nil {
	// 	golog.Error(err)
	// 	xmux.GetInstance(r).Response.(*response.Response).Code = 1
	// 	xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
	// 	return
	// }
	// xmux.GetInstance(r).Response.(*response.Response).Data = kns
}

func RoleGroupDel(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	golog.Info(id)
	ssql := "select count(id) from jobs where rgid=?"
	var count int
	err := db.Mconn.GetOne(ssql, id).Scan(&count)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	if count > 0 {
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = "有用户在使用， 无法删除"
		return
	}
	// 先删除perm
	// 获取 permids
	rolegroup := model.RolePermMap{}
	err = rolegroup.GetRoleGroupById(id)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

	// 然后删除rolegroup
	err = rolegroup.Delete()
	// perm 里面的也要删除
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}

}

func EditRoleGroup(w http.ResponseWriter, r *http.Request) {

	// rr := xmux.GetInstance(r).Data.(*RequestRole)
	// ids := make([]int64, 0, len(rr.PermIds))
	// for _, v := range rr.PermIds {
	// 	err := v.Update()
	// 	if err != nil {
	// 		golog.Error(err)
	// 	}
	// 	ids = append(ids, v.Id)
	// }
	rolegroup := model.RolePermMap{
		// ID:   rr.ID,
		// Name: rr.Name,
	}
	err := rolegroup.Update()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
}

func AddRoleGroup(w http.ResponseWriter, r *http.Request) {

	// rr := xmux.GetInstance(r).Data.(*RequestRole)
	// rolegroup := model.RolePermMap{
	// 	ID:   rr.ID,
	// 	Name: rr.Name,
	// }
	// err := rolegroup.Insert()
	// if err != nil {
	// 	golog.Error(err)
	// 	xmux.GetInstance(r).Response.(*response.Response).Code = 1
	// 	xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
	// 	return
	// }
	// xmux.GetInstance(r).Response.(*response.Response).ID = rolegroup.ID

}

// 添加角色
func AddRole(w http.ResponseWriter, r *http.Request) {
	rr := xmux.GetInstance(r).Data.(*service.RequestRole)
	id, err := service.CreateRole(rr)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	// err := rolegroup.Insert()
	// if err != nil {
	// 	golog.Error(err)
	// 	xmux.GetInstance(r).Response.(*response.Response).Code = 1
	// 	xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
	// 	return
	// }
	xmux.GetInstance(r).Response.(*response.Response).ID = id

}

func DelRole(w http.ResponseWriter, r *http.Request) {
	rr := xmux.GetInstance(r).Data.(*service.RequestRole)
	golog.Info(rr.ID)
	err := service.DeleteRole(rr)
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	// err := rolegroup.Insert()
	// if err != nil {
	// 	golog.Error(err)
	// 	xmux.GetInstance(r).Response.(*response.Response).Code = 1
	// 	xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
	// 	return
	// }
	// xmux.GetInstance(r).Response.(*response.Response).ID = id

}

// 角色列表
func RoleList(w http.ResponseWriter, r *http.Request) {
	role := model.Role{}
	roles, err := role.List()
	if err != nil {
		golog.Error(err)
		xmux.GetInstance(r).Response.(*response.Response).Code = 1
		xmux.GetInstance(r).Response.(*response.Response).Msg = err.Error()
		return
	}
	xmux.GetInstance(r).Response.(*response.Response).Data = roles

}
