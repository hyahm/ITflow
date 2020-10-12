package usergroup

import (
	"errors"
	"itflow/cache"
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
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

func (ug *RespUpdateUserGroup) GetIds() (string, error) {
	rows, err := db.Mconn.GetRowsIn("select id from user where realname in (?)",
		(gomysql.InArgs)(ug.Users).ToInArgs())
	if err != nil {
		golog.Error(err)
		return "", err
	}
	ids := make([]string, 0)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			golog.Info(err)
			continue
		}
		ids = append(ids, id)
	}
	rows.Close()
	return strings.Join(ids, ","), nil
}
