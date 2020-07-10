package model

import (
	"itflow/cache"
	"itflow/db"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

// RoleGroup: 角色组表， 由管理员分配， 管理可以操作的页面, 与用户rid关联
type RoleGroup struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Permids string `json:"permids"`
}

func NewRoleGroup(uid int64) (*RoleGroup, error) {
	// 通过nickname 来获取rid
	rid := cache.CacheUidRid[uid]
	rg := &RoleGroup{}
	err := db.Mconn.GetOne("select name, permids from rolegroup where id=?", rid).Scan(&rg.Name, &rg.Permids)
	if err != nil {
		golog.Error(err)
		return nil, err
	}

	return rg, nil
}

func (rg *RoleGroup) Insert() error {
	var err error
	rg.ID, err = db.Mconn.Insert("insert into rolegroup(name,permids) values(?,?)", rg.Name, rg.Permids)
	return err
}

func (rg *RoleGroup) Update() error {
	_, err := db.Mconn.Update("update rolegroup set name=? where id=?", rg.Name, rg.ID)
	// 查询到permids
	err = db.Mconn.GetOne("select permids from rolegroup where id=?", rg.ID).Scan(&rg.Permids)
	return err
}

func (rg *RoleGroup) CheckPagePerm(name string) bool {
	pl := strings.Split(rg.Permids, ",")
	for _, v := range pl {
		id, _ := strconv.ParseInt(v, 10, 64)
		if cache.CacheRidRole[id] == name {
			return true
		}
	}
	return false
}
