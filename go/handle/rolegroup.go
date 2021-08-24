package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/rolegroup"
	"itflow/model"
	"itflow/response"
	"net/http"

	"github.com/hyahm/goconfig"
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
	ssql := "select count(id) from jobs where rid=?"
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

	data := xmux.GetInstance(r).Data.(*rolegroup.ReqRoleGroup)
	// uid := xmux.GetInstance(r).Get("uid").(int64)
	// w.Write(data.Update(uid))
	// 先修改perm表里面的

	for _, perm := range data.Perms {
		err := perm.Update()
		if err != nil {
			golog.Error(err)
			w.Write(response.ErrorE(err))
			return
		}
	}

	// 修改rolegroup的
	rolegroup := model.RoleGroup{
		ID:      data.Id,
		Name:    data.Name,
		PermIds: nil,
	}
	err := rolegroup.Update()
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	w.Write(response.Success())
}

func AddRoleGroup(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetInstance(r).Data.(*rolegroup.ReqRoleGroup)
	uid := xmux.GetInstance(r).Get("uid").(int64)
	if uid != goconfig.ReadInt64("server.adminId", 1) {
		golog.Error("no permission")
		w.Write(response.Error("no permission"))
		return
	}
	golog.Infof("%#v", *data)
	// 先插入perm表
	permids, err := model.InsertManyPerm(data.Perms)
	if err != nil {
		golog.Error(err)
		w.Write(response.ErrorE(err))
		return
	}
	// 插入到rolegroup表
	rolegroup := model.RoleGroup{
		ID:      data.Id,
		Name:    data.Name,
		PermIds: permids,
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

type ResponseRoleTemplate struct {
	Code     int                   `json:"code"`
	Msg      string                `json:"msg"`
	Template []*rolegroup.PermRole `json:"template,omitempty"`
}

func (rrt *ResponseRoleTemplate) Marshal() []byte {
	send, err := json.Marshal(rrt)
	if err != nil {
		golog.Error(err)
		return nil
	}
	return send
}

func RoleTemplate(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	data := &ResponseRoleTemplate{
		Template: make([]*rolegroup.PermRole, 0),
	}
	rows, err := db.Mconn.GetRows("select name, info from roles")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		role := &rolegroup.PermRole{}
		err = rows.Scan(&role.Name, &role.Info)
		if err != nil {
			golog.Info(err)
			continue
		}
		data.Template = append(data.Template, role)
	}
	rows.Close()
	w.Write(data.Marshal())
}
