package project

import (
	"itflow/cache"
	"itflow/internal/response"
	"itflow/model"
	"strconv"
	"strings"

	"github.com/hyahm/golog"
)

func (rp *ReqProject) Update(userid int64) ([]byte, error) {
	// 处理数据， 返回所需
	resp := &response.Response{}
	err := rp.checkValid()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	uids := make([]string, 0)
	needAdd := true
	for _, v := range rp.SelectUser {
		// 获取的真实用户名
		if uid, ok := cache.CacheRealNameUid[v]; ok {
			if uid == userid {
				needAdd = false
			}
			uids = append(uids, strconv.FormatInt(uid, 10))
		}
	}
	if needAdd {
		uids = append(uids, strconv.FormatInt(userid, 10))
	}
	project := &model.Project{
		Id:          rp.Id,
		Name:        rp.ProjectName,
		Participant: strings.Join(uids, ","),
	}
	golog.Info("update")
	err = project.Update()
	if err != nil {
		golog.Error(err)
		return resp.ErrorE(err), err
	}
	resp.Id = project.Id
	cache.CachePidName[resp.Id] = rp.ProjectName
	return resp.Success(), nil
}
