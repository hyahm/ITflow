package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/xmux"
)

// RoleGroup: 角色组表， 由管理员分配， 管理可以操作的页面, 与用户rid关联
type RoleGroup struct {
	ID      int64   `json:"id" db:"id,default"`
	Name    string  `json:"name" db:"name,default"`
	PermIds []int64 `json:"permids" db:"permids"`
}

func GetRoleKeyName() ([]KeyName, error) {
	s := "select id, name from rolegroup"
	rows, err := db.Mconn.GetRows(s)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	kns := make([]KeyName, 0)
	defer rows.Close()
	for rows.Next() {
		kn := KeyName{}
		err = rows.Scan(&kn.ID, &kn.Name)
		if err != nil {
			golog.Error(err)
			continue
		}
		kns = append(kns, kn)

	}
	return kns, nil
}

func (rg *RoleGroup) Delete() error {
	result := db.Mconn.Delete("delete from rolegroup where id=?", rg.ID)
	return result.Err
}

func (rg *RoleGroup) GetRoleGroupById(id interface{}) error {
	result := db.Mconn.Select(&rg, "select * from rolegroup where id=?", id)
	return result.Err
}

func (rg *RoleGroup) GetEditDataById(id interface{}) (interface{}, error) {
	// 通过uid 来获取rid
	err := rg.GetRoleGroupById(id)
	if err != nil {
		return nil, err
	}
	// 需要特殊返回值
	//  id: 0,
	// rid: v.id,
	// label: this.defaultPerm,
	// value: [],
	// info: v.info,
	type perm struct {
		ID    int64    `json:"id"`
		Label []string `json:"label"`
		Value []string `json:"value"`
		Rid   int64    `json:"rid"`
		Info  string   `json:"info"`
	}

	rows, err := db.Mconn.GetRowsIn("select id, pv, rid from perm where id in (?)", rg.PermIds)
	if err != nil {
		return nil, err
	}
	editPerm := make([]perm, 0, len(cache.CacheRoleID))
	defer rows.Close()
	for rows.Next() {
		p := perm{
			Label: []string{"read", "create", "update", "delete"},
			Value: make([]string, 0, 4),
		}
		var pv uint8
		err = rows.Scan(&p.ID, &pv, &p.Rid)
		if err != nil {
			golog.Error(err)
			continue
		}
		p.Info = cache.CacheRoleID[p.Rid].Info
		result := xmux.GetPerm(p.Label, pv)
		for i, v := range result {
			if v {
				p.Value = append(p.Value, p.Label[i])
			}
		}
		editPerm = append(editPerm, p)
	}
	return editPerm, nil
}

func RoleGroupList() ([]RoleGroup, error) {
	// 通过uid 来获取rid
	rg := make([]RoleGroup, 0)
	result := db.Mconn.Select(&rg, "select * from rolegroup")
	return rg, result.Err
}

func (rg *RoleGroup) Insert() error {
	result := db.Mconn.InsertInterfaceWithID(rg, "insert into rolegroup($key) values($value)")
	if result.Err != nil {
		return result.Err
	}
	rg.ID = result.LastInsertId
	return nil
}

func (rg *RoleGroup) Update() error {
	result := db.Mconn.UpdateInterface(rg, "update rolegroup set $set where id=?", rg.ID)
	return result.Err
}

func CheckRoleNameInGroup(name string, rid *int64) error {
	err := db.Mconn.GetOne("select id from rolegroup where name=?",
		strings.Trim(name, " ")).Scan(rid)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("不存在此角色组")
		}
		return err
	}

	return nil
}

// 获取用户组里面的 permids
func GetPermIdsByUid(uid interface{}) ([]int64, error) {
	var perms []byte
	err := db.Mconn.GetOne("select permids from rolegroup where id=(select rgid from jobs where id=(select jid from user where id=?))", uid).Scan(
		&perms,
	)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	permids := make([]int64, 0)
	err = json.Unmarshal(perms, &permids)
	if err != nil {
		golog.Error(err)
		return nil, err
	}
	return permids, nil
}
