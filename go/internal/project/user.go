package project

import (
	"itflow/cache"
	"itflow/db"
	"itflow/model"

	"github.com/hyahm/golog"
)

// 通过project 获取用户

func GetUsersByProjectName(userid int64, name cache.Project) []byte {
	// 根据项目组获取所属用户和所属版本
	resp := &MyProject{
		Name: make([]string, 0),
	}
	namessql := "select realname from user where in in (select ids from usergroup where id = (select ugid from project where name=?))"
	rows, err := db.Mconn.GetRows(namessql, name)
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err)
	}
	for rows.Next() {
		name := new(string)
		err = rows.Scan(name)
		if err != nil {
			golog.Error(err)
			continue
		}
		resp.Name = append(resp.Name, *name)
	}

	version := &model.Version{}
	ps, err := version.GetProjectNameByPid(name.Id())
	if err != nil {
		resp.Name = make([]string, 0)
		resp.Code = 1
		resp.Msg = err.Error()
		return resp.Marshal()
	}
	resp.Versions = ps
	return resp.Marshal()
}
