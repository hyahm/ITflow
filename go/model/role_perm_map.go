package model

import (
	"database/sql"
	"encoding/json"
	"errors"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
)

// RoleGroup: 角色组表， 由管理员分配， 管理可以操作的页面, 与用户rid关联
type RolePermMap struct {
	Rid    int64 `json:"rid" gorm:"column:rid"`
	Permid int64 `json:"perm_id" gorm:"column:perm_id"`
	Enable bool  `json:"enable" gorm:"column:enable"`
}

func (RolePermMap) TableName() string {
	return "role_perm_map"
}

func (rg *RolePermMap) Create() error {
	return db.Gorm.Table(rg.TableName()).Create(rg).Error
}

func (rg *RolePermMap) CreateMany(rolePerms *[]RolePermMap) error {
	return db.Gorm.Table(rg.TableName()).Create(rolePerms).Error
}

func (rg *RolePermMap) Delete() error {
	return db.Gorm.Table(rg.TableName()).Select("rid").Delete(rg).Error
}

// 启用所有权限
func (rg *RolePermMap) EnablePerm(rolePerms []int64) error {
	if rg.Rid <= 0 {
		return errors.New("rid没有设置")
	}
	rg.Enable = true
	return db.Gorm.Table(rg.TableName()).Where("rid=?", rg.Rid).Select("enable").Where("perm_id in ?", rolePerms).Updates(rg).Error
}

// 禁用所有权限
func (rg *RolePermMap) DisablePerm(rolePerms []int64) error {
	if rg.Rid <= 0 {
		return errors.New("rid没有设置")
	}
	return db.Gorm.Table(rg.TableName()).Select("enable").Where("rid=?", rg.Rid).Updates(rg).Error
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

func (rg *RolePermMap) GetRoleGroupById(id interface{}) error {
	result := db.Mconn.Select(&rg, "select * from rolegroup where id=?", id)
	return result.Err
}

func (rg *RolePermMap) GetEditDataById(id interface{}) (interface{}, error) {
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

	return nil, nil
}

func RoleGroupList() ([]RolePermMap, error) {
	// 通过uid 来获取rid
	rg := make([]RolePermMap, 0)
	result := db.Mconn.Select(&rg, "select * from rolegroup")
	return rg, result.Err
}

func (rg *RolePermMap) Insert() error {
	result := db.Mconn.InsertInterfaceWithID(rg, "insert into rolegroup($key) values($value)")
	if result.Err != nil {
		return result.Err
	}
	rg.Rid = result.LastInsertId
	return nil
}

func (rg *RolePermMap) Update() error {
	result := db.Mconn.UpdateInterface(rg, "update rolegroup set $set where id=?", rg.Rid)
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
