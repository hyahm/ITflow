package usergroup

import (
	"errors"
	"itflow/cache"
	"itflow/db"

	"github.com/hyahm/golog"
)

type RespUpdateUserGroup struct {
	Id    int64    `json:"id"`
	Name  string   `json:"name"`
	Users []string `json:"users"`
}

func (ug *RespUpdateUserGroup) CheckUser(uid int64) error {
	var dbUid int64
	err := db.Mconn.GetOne("select uid from usergroup where id=?", ug.Id).Scan(&dbUid)
	if err != nil {
		golog.Error(err)
		return err
	}

	if dbUid != cache.SUPERID && dbUid != uid {
		golog.Debug("没有权限")

		return errors.New("没有权限")
	}
	return nil
}
