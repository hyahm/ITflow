package model

import (
	"itflow/app/bugconfig"
	"itflow/db"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

// RoleGroup: 角色组表， 由管理员分配， 管理可以操作的页面, 与用户rid关联
type RoleGroup struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	RoleList  string   `json:"rolelist"`
	RoleArray []string `json:"rolearray"`
}

//
func (rg *RoleGroup) split() {
	rg.RoleArray = strings.Split(rg.RoleList, ",")

}

func (rg *RoleGroup) Join() {
	rg.RoleList = strings.Join(rg.RoleArray, ",")

}

func NewRoleGroup(nickname string) (*RoleGroup, error) {
	// 通过nickname 来获取rid
	uid := bugconfig.CacheNickNameUid[nickname]
	rid := bugconfig.CacheUidRid[uid]
	rg := &RoleGroup{}
	row, err := db.Mconn.GetOne("select name, rolelist from rolegroup where id=?", rid)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	err = row.Scan(&rg.Name, &rg.RoleList)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	rg.split()
	return rg, nil
}

func (rg *RoleGroup) CheckPagePerm(name string) bool {
	rg.RoleArray = strings.Split(rg.RoleList, ",")
	for _, v := range rg.RoleArray {
		id, _ := strconv.ParseInt(v, 10, 64)
		if bugconfig.CacheRidRole[id] == name {
			return true
		}
	}
	return false
}
