package handle

import (
	"database/sql"
	"encoding/json"
	"itflow/cache"
	"itflow/db"
	"itflow/internal/response"
	"itflow/internal/rolegroup"
	"itflow/model"
	"net/http"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

func RoleGroupList(w http.ResponseWriter, r *http.Request) {

	errorcode := &response.Response{}
	data := &rolegroup.RespRoleGroup{
		RoleList: make([]*rolegroup.ReqRoleGroup, 0),
	}

	s := "select id,name,permids from rolegroup"
	rows, err := db.Mconn.GetRows(s)
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
		for _, v := range strings.Split(permids, ",") {
			perm, err := model.NewPerm(v)
			if err != nil {
				if err == sql.ErrNoRows {
					continue
				}
				golog.Error(err)
				w.Write(data.ErrorE(err))
				return
			}
			if name, ok := cache.CacheRidRole[perm.Rid]; ok {
				if info, infook := cache.CacheRidInfo[perm.Rid]; infook {
					one.RoleList = append(one.RoleList, &rolegroup.PermRole{
						Add:    perm.Increase,
						Select: perm.Find,
						Update: perm.Revise,
						Remove: perm.Remove,
						Name:   name,
						Info:   info,
					})
				}

			}
			// 最好清除无效的数据

		}
		data.RoleList = append(data.RoleList, one)
	}
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
	isql := "delete from  rolegroup where id = ?"
	_, err = db.Mconn.Update(isql, id)
	if err != nil {
		golog.Error(err)
		w.Write(errorcode.ErrorE(err))
		return
	}

	// 删除缓存

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

	data := make([]*rolegroup.PermRole, 0)

	for rid, info := range cache.CacheRidInfo {
		data = append(data, &rolegroup.PermRole{
			Info: info,
			Name: cache.CacheRidRole[rid],
		})
	}

	send, _ := json.Marshal(data)
	w.Write(send)
	return

}
