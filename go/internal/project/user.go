package project

import (
	"itflow/cache"
	"itflow/model"
	"strconv"
	"strings"
)

// 通过project 获取用户

func GetUsersByProjectName(userid int64, name cache.Project) []byte {

	resp := &MyProject{
		Name: make([]string, 0),
	}
	id, ok := cache.CacheProjectPid[name]
	if !ok {
		resp.Code = 1
		resp.Msg = "没找到用户"
		return resp.Marshal()
	}
	p, err := model.NewProjectById(id)
	if err != nil {
		resp.Code = 1
		resp.Msg = err.Error()
		return resp.Marshal()
	}
	uids := strings.Split(p.Participant, ",")
	var perm bool
	for _, u := range uids {
		uid, err := strconv.ParseInt(u, 10, 64)
		if err != nil {
			continue
		}
		if uid == userid {
			perm = true
		}
		resp.Name = append(resp.Name, cache.CacheUidRealName[uid])
	}
	if !perm {
		resp.Name = make([]string, 0)
		resp.Code = 1
		resp.Msg = "你没有此项目权限"
		return resp.Marshal()
	}
	return resp.Marshal()
}
