package project

import (
	"itflow/db"
	"strings"

	"github.com/hyahm/golog"
	"github.com/hyahm/gomysql"
)

// 通过project 获取用户

func GetUsersByProjectName(userid int64, name string) []byte {
	// 根据项目组获取所属用户和所属版本
	resp := &MyProject{
		Name:     make([]string, 0),
		Versions: make([]string, 0),
	}
	var ids string
	err := db.Mconn.GetOne("select ids from usergroup where id = (select ugid from project where name=?)",
		name).Scan(&ids)
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err)
	}
	namessql := "select realname from user where id in (?)"
	rows, err := db.Mconn.GetRowsIn(namessql, (gomysql.InArgs)(strings.Split(ids, ",")).ToInArgs())
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err)
	}
	n := new(string)
	for rows.Next() {
		err = rows.Scan(n)
		if err != nil {
			golog.Error(err)
			continue
		}
		resp.Name = append(resp.Name, *n)
	}

	vrows, err := db.Mconn.GetRows("select name from version where pid=(select id from project where name=?)", name)
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err)
	}
	for vrows.Next() {
		err = vrows.Scan(n)
		if err != nil {
			golog.Error(err)
			continue
		}
		resp.Versions = append(resp.Versions, *n)
	}

	return resp.Marshal()
}
