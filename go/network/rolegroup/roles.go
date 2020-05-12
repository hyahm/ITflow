package rolegroup

import (
	"itflow/app/bugconfig"
	"itflow/db"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

type RoleGroup struct {
	Id       int64    `json:"id,omitempty" type:"int" need:"是" default:"" information:"id"`
	Name     string   `json:"name,omitempty" type:"string" need:"是" default:"" information:"角色组名"`
	RoleList []string `json:"rolelist,omitempty" type:"array" need:"是" default:"" information:"角色组成员"`
}

type RoleGroupList struct {
	DataList []*RoleGroup `json:"datalist" type:"array" need:"否" default:"" information:"角色组列表"`
	Code     int          `json:"code" type:"int" need:"是" default:"" information:"状态码"`
}

type Roles struct {
	Roles []string `json:"roles"`
	Code  int      `json:"code"`
}

type Updata_role struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Oldname string `json:"oldname"`
}

func (rg *RoleGroup) Update() error {
	rl := make([]string, 0)
	for _, v := range rg.RoleList {
		rl = append(rl, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
	}
	gsql := "update rolegroup set name=?,rolelist=?  where id=?"
	_, err := db.Mconn.Update(gsql, rg.Name, strings.Join(rl, ","), rg.Id)
	if err != nil {
		golog.Error(err)
		return err
	}
	return nil
}

func (rg *RoleGroup) Insert() error {
	ids := make([]string, 0)
	for _, v := range rg.RoleList {
		ids = append(ids, strconv.FormatInt(bugconfig.CacheRoleRid[v], 10))
	}
	gsql := "insert rolegroup(name,rolelist) values(?,?)"
	var err error
	rg.Id, err = db.Mconn.Insert(gsql, rg.Name, strings.Join(ids, ","))
	if err != nil {
		golog.Error(err)
		return err
	}
	return nil
}
