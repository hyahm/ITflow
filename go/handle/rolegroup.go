package handle

import (
	"encoding/json"
	"itflow/db"
	"itflow/internal/response"
	"itflow/internal/rolegroup"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
	"github.com/hyahm/xmux"
)

func RoleGroupList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := &rolegroup.RespRoleGroup{
		RoleList: make([]*rolegroup.ReqRoleGroup, 0),
	}

	rows, err := db.Mconn.GetRows("select id,name,permids from rolegroup")
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	for rows.Next() {
		var permids string // 保存perm表的所有id
		one := &rolegroup.ReqRoleGroup{
			RoleList: make([]*rolegroup.PermRole, 0),
		}
		rows.Scan(&one.Id, &one.Name, &permids)
		golog.Info(one.Id, " ", one.Name)
		permrows, err := db.Mconn.GetRowsIn("select find, remove, revise, increase, r.name, r.info from perm as p join roles as r on p.id in (?) and p.rid=r.id",
			(gomysql.InArgs)(strings.Split(permids, ",")).ToInArgs())
		if err != nil {
			golog.Error(err)
			w.Write(errorcode.ErrorE(err))
			return
		}
		for permrows.Next() {
			rp := &rolegroup.PermRole{}
			err = permrows.Scan(&rp.Select, &rp.Remove, &rp.Update, &rp.Add, &rp.Name, &rp.Info)
			if err != nil {
				golog.Error(err)
				w.Write(errorcode.ErrorE(err))
				return
			}
			one.RoleList = append(one.RoleList, rp)
		}
		permrows.Close()
		data.RoleList = append(data.RoleList, one)
	}
	rows.Close()
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}

func GetRoleGroupName(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	s := "select name from rolegroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	resp := &struct {
		Code     int      `json:"code"`
		RoleList []string `json:"rolelist"`
	}{
		RoleList: make([]string, 0),
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		resp.RoleList = append(resp.RoleList, name)

	}

	send, _ := json.Marshal(resp)
	w.Write(send)
}

func RoleGroupDel(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}

	id := r.FormValue("id")
	golog.Info(id)
	ssql := "select count(id) from user where rid=?"
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
	// 需要用到事务
	var permids string
	err = db.Mconn.GetOne("select permids from rolegroup where id=?", id).Scan(&permids)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	_, err = db.Mconn.DeleteIn("delete from perm where id in(?)",
		(gomysql.InArgs)(strings.Split(permids, ",")).ToInArgs())
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	isql := "delete from rolegroup where id = ?"
	_, err = db.Mconn.Update(isql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}
	// perm 里面的也要删除

	send, _ := json.Marshal(errorcode)
	w.Write(send)
	return

}

func EditRoleGroup(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetData(r).Data.(*rolegroup.ReqRoleGroup)

	uid := xmux.GetData(r).Get("uid").(int64)
	w.Write(data.Update(uid))
	return

}

func AddRoleGroup(w http.ResponseWriter, r *http.Request) {

	data := xmux.GetData(r).Data.(*rolegroup.ReqRoleGroup)
	uid := xmux.GetData(r).Get("uid").(int64)
	w.Write(data.Add(uid))
	return

}

func RoleTemplate(w http.ResponseWriter, r *http.Request) {
	errorcode := &response.Response{}
	data := make([]*rolegroup.PermRole, 0)
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
		data = append(data, role)
	}
	rows.Close()
	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
