package service

import (
	"itflow/model"

	"github.com/hyahm/golog"
)

type RequestRole struct {
	ID      int64   `json:"id" form:"id"`
	Name    string  `json:"name" `
	PermIds []int64 `json:"perm_ids"`
}

func CreateRole(rr *RequestRole) (int64, error) {
	role := model.Role{
		Name: rr.Name,
	}

	err := role.Create()
	if err != nil {
		golog.Error(err)
		return 0, err
	}
	rpms := make([]model.RolePermMap, 0, len(model.PagePerms))
	for key := range model.PagePerms {
		rpms = append(rpms, model.RolePermMap{
			Rid:    role.Id,
			Permid: key,
		})
	}

	rpm := model.RolePermMap{
		Rid: role.Id,
	}
	// 插入的时候,将所有的权限都插入,默认没有权限

	err = rpm.CreateMany(&rpms)
	if err != nil {
		golog.Error(err)
		return 0, err
	}
	if len(rr.PermIds) > 0 {
		err = rpm.EnablePerm(rr.PermIds)
		if err != nil {
			golog.Error(err)
			return 0, err
		}
	}
	return role.Id, nil
}

func DeleteRole(rr *RequestRole) error {
	role := model.Role{
		Id: rr.ID,
	}

	err := role.Delete()
	if err != nil {
		golog.Error(err)
		return err
	}
	// rpms := make([]model.RolePermMap, 0, len(model.PagePerms))
	// for key := range model.PagePerms {
	// 	rpms = append(rpms, model.RolePermMap{
	// 		Rid:    role.Id,
	// 		Permid: key,
	// 	})
	// }

	rpm := model.RolePermMap{
		Rid: role.Id,
	}
	// 插入的时候,将所有的权限都插入,默认没有权限

	return rpm.Delete()
}
